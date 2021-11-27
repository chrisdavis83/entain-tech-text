package db

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	"time"

	"git.neds.sh/matty/entain/racing/proto/racing"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequest) ([]*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

func (r *racesRepo) List(request *racing.ListRacesRequest) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	filter := request.Filter
	query = getRaceQueries()[racesList]

	query, args = r.applyFilter(query, filter)
	query, args = r.applySortOrdering(query, request)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	// Handle the filtering of only showing the visible races
	if filter.OnlyShowVisible {
		clauses = append(clauses, "visible = (?)")
		args = append(args, filter.OnlyShowVisible)
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

func (m *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart, &race.Status); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts

		races = append(races, &race)
	}

	return races, nil
}

func (r *racesRepo) applySortOrdering(query string, request *racing.ListRacesRequest) (string, []interface{}) {
	var (
		args []interface{}
	)

	if !orderByHasValidValue(request) {
		return query, args
	}

	query += "\nORDER BY " + request.OrderBy + " " + request.SortBy.String()

	return query, args

}

var validColumnNames = []string{"id", "meeting_id", "name", "number", "visible", "advertised_start_time"}

// This is to protect against SQL injection, i could also use a map from rest field to database column so the
// order by attribute could use the rest attribute not the database column name
func orderByHasValidValue(request *racing.ListRacesRequest) bool {
	var result = false
	for _, x := range validColumnNames {
		if x == request.OrderBy {
			result = true
			break
		}
	}
	return result
}
