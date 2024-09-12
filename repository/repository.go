package repository

import (
	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository/forecastrepo"
	"gorm.io/gorm"
)

type ForecastRepository interface {
	FindAll() ([]*model.Forecast, error)
	FindByID(id int) (*model.Forecast, error)
	Save(forecast *model.Forecast) error
	Update(forecast *model.Forecast) error
}

type Repositories struct {
	ForecastRepo *forecastrepo.ForecastRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	forecastRepo := forecastrepo.NewForecastRepo(db)
	return &Repositories{ForecastRepo: forecastRepo}
}
