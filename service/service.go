package service

import (
	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository"
	taskservice "github.com/brunompx/go-riverlevels/service/forecastservice"
)

type ForecastService interface {
	FindAll() ([]*model.Forecast, error)
	FindByID(id int) (*model.Forecast, error)
	Save(forecast *model.Forecast) error
	Update(forecast *model.Forecast) error
}

// Service storage of all services.
type Service struct {
	ForecastService ForecastService
}

// implementation for storage of all services.
func InitServices(repositories *repository.Repositories) *Service {
	return &Service{
		ForecastService: taskservice.NewForecastService(repositories.ForecastRepo),
	}
}
