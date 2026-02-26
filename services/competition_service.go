package services

import (
	"apis/models/competition"
	"context"
)

type CompetitionService struct {
	ListCompetition   func(ctx context.Context) ([]competition.Competition, error)
	CreateCompetition func(ctx context.Context, req *competition.CompetitionRequest) (competition.Competition, error)
	GetCompetition    func(ctx context.Context, id string) (competition.Competition, error)
	UpdateCompetition func(ctx context.Context, id string, req *competition.CompetitionRequest) (competition.Competition, error)
}

var competitionService *CompetitionService

func GetCompetitionService() *CompetitionService {
	if competitionService == nil {
		return newCompetitionService()
	}
	return competitionService
}

func newCompetitionService() *CompetitionService {
	service := &CompetitionService{}

	competitionService = service
	return service
}

func init() {
	GetCompetitionService()
}
