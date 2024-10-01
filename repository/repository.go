package repository

import (
	"github.com/brunompx/go-riverlevels/types"
	"gorm.io/gorm"
)

type ForecastRepository interface {
	FindAll() ([]*types.Forecast, error)
	FindByID(id int) (*types.Forecast, error)
	Save(forecast *types.Forecast) error
	Update(forecast *types.Forecast) error
	FindForecast(forecast *types.Forecast) (types.Forecast, error)
}

type MeasureRepository interface {
	FindAll() ([]*types.Forecast, error)
	FindByID(id int) (*types.Forecast, error)
	Save(forecast *types.Forecast) error
	Update(forecast *types.Forecast) error
	FindForecast(forecast *types.Forecast) (types.Forecast, error)
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
