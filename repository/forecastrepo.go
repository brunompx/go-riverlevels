package repository

import (
	"fmt"

	"github.com/brunompx/go-riverlevels/types"
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

func (r *ForecastRepo) Save(forecast *types.Forecast) error {
	result := r.db.Create(&forecast)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ForecastRepo) FindForecast(forecast *types.Forecast) (types.Forecast, error) {

	fmt.Println("Busco?: ", forecast.EstacionNombre)
	fmt.Println("Busco?: ", forecast.CorId)

	var fore types.Forecast
	//r.db.Where("name = ? AND age >= ?", "jinzhu", "22").First(&fore)
	r.db.Where(forecast).First(&fore)
	return fore, nil
}

func (r *ForecastRepo) FindAll() ([]*types.Forecast, error) {
	return nil, nil
}

func (r *ForecastRepo) FindByID(id int) (*types.Forecast, error) {
	return nil, nil
}

func (r *ForecastRepo) Update(forecast *types.Forecast) error {
	return nil
}
