package repository

import (
	"github.com/brunompx/go-riverlevels/model"
	"gorm.io/gorm"
)

type ForecastRepository interface {
	FindAll() ([]*model.Forecast, error)
	FindByID(id int) (*model.Forecast, error)
	Save(forecast *model.Forecast) error
	Update(forecast *model.Forecast) error
	FindForecast(forecast *model.Forecast) (model.Forecast, error)
}

type MeasureRepository interface {
	FindAll() ([]*model.Forecast, error)
	FindByID(id int) (*model.Forecast, error)
	Save(forecast *model.Forecast) error
	Update(forecast *model.Forecast) error
	FindForecast(forecast *model.Forecast) (model.Forecast, error)
}

type Repositories struct {
	ForecastRepo *ForecastRepo
	MeasureRepo  *MeasureRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	forecastRepo := NewForecastRepo(db)
	measureRepo := NewMeasureRepo(db)
	return &Repositories{
		ForecastRepo: forecastRepo,
		MeasureRepo:  measureRepo,
	}
}
