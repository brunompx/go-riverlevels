package forecastrepo

import "gorm.io/gorm"

type ForecastRepo struct {
	db *gorm.DB
}

func NewForecastRepo(db *gorm.DB) *ForecastRepo {
	return &ForecastRepo{
		db: db,
	}
}

func (r *ForecastRepo) Save() {

}
