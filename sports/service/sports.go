package service

import (
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sports interface {
	// ListSports will return a collection of sports.
	ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error)
}

// sportsService implements the Sports interface.
type sportsService struct {
	sportsRepo db.SportsRepo
}

// NewSportsService instantiates and returns a new sportsService.
func NewSportsService(sportsRepo db.SportsRepo) Sports {
	return &sportsService{sportsRepo}
}

func (s *sportsService) ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error) {
	allSports, err := s.sportsRepo.List(in)
	if err != nil {
		return nil, err
	}

	return &sports.ListSportsResponse{Sports: allSports}, nil
}
