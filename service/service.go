package service

import "github.com/brunompx/go-riverlevels/model"

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
