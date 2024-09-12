package forecastrepo

import (
	"github.com/brunompx/go-riverlevels/model"
	"gorm.io/gorm"
)

type ForecastRepo struct {
	db *gorm.DB
}

func NewForecastRepo(db *gorm.DB) *ForecastRepo {
	return &ForecastRepo{
		db: db,
	}
}

func (r *ForecastRepo) Save(forecast *model.Forecast) error {
	return nil
}

func (r *ForecastRepo) FindAll() ([]*model.Forecast, error) {
	return nil, nil
}

func (r *ForecastRepo) FindByID(id int) (*model.Forecast, error) {
	return nil, nil
}

func (r *ForecastRepo) Update(forecast *model.Forecast) error {
	return nil
}
