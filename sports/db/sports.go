package db

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"time"

	"git.neds.sh/matty/entain/sports/proto/sports"
)

// SportsRepo provides repository access to sports.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of sports.
	List(filter *sports.ListSportsRequest) ([]*sports.Sport, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sports.
		err = r.seed()
	})

	return err
}

func (r *sportsRepo) List(request *sports.ListSportsRequest) ([]*sports.Sport, error) {
	var (
		err   error
		query string
	)

	query = getSportQueries()[sportsList]

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	return r.scanSports(rows)
}

func (m *sportsRepo) scanSports(
	rows *sql.Rows,
) ([]*sports.Sport, error) {
	var allSports []*sports.Sport

	for rows.Next() {
		var sport sports.Sport
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id, &sport.Name, &sport.SportType, &sport.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		allSports = append(allSports, &sport)
	}

	return allSports, nil
}
