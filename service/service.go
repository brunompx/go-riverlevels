package service

import (
	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository"
)

type ForecastServiceInt interface {
	FindAll() ([]*model.Forecast, error)
	FindByID(id int) (*model.Forecast, error)
	Save(forecast *model.Forecast) error
	Update(forecast *model.Forecast) error
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
