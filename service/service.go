package service

import (
	"github.com/brunompx/go-riverlevels/repository"
	"github.com/brunompx/go-riverlevels/types"
)

type ForecastServiceInt interface {
	FindAll() ([]*types.Forecast, error)
	FindByID(id int) (*types.Forecast, error)
	Save(forecast *types.Forecast) error
	Update(forecast *types.Forecast) error
}

// Service storage of all services.
type Service struct {
	ForecastService ForecastServiceInt
}

// implementation for storage of all services.
func InitServices(repositories *repository.Repositories) *Service {
	return &Service{
		ForecastService: NewForecastService(repositories.ForecastRepo),
	}
}
