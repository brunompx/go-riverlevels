package repository

import (
	"fmt"

	"github.com/brunompx/go-riverlevels/types"
	"gorm.io/gorm"
)

type MeasureRepo struct {
	db *gorm.DB
}

func NewMeasureRepo(db *gorm.DB) *MeasureRepo {
	return &MeasureRepo{
		db: db,
	}
}

func (r *MeasureRepo) Save(forecast *types.Forecast) error {
	result := r.db.Create(&forecast)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *MeasureRepo) FindForecast(forecast *types.Forecast) (types.Forecast, error) {

	fmt.Println("Busco?: ", forecast.EstacionNombre)
	fmt.Println("Busco?: ", forecast.CorId)

	var fore types.Forecast
	//r.db.Where("name = ? AND age >= ?", "jinzhu", "22").First(&fore)
	r.db.Where(forecast).First(&fore)
	return fore, nil
}

func (r *MeasureRepo) FindAll() ([]*types.Forecast, error) {
	return nil, nil
}

func (r *MeasureRepo) FindByID(id int) (*types.Forecast, error) {
	return nil, nil
}

func (r *MeasureRepo) Update(forecast *types.Forecast) error {
	return nil
}
