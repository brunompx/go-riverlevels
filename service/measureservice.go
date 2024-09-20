package service

import (
	"fmt"

	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository"
)

type MeasureService struct {
	measureRepository repository.MeasureRepository // interface of the repository, not implementation
}

func (s *MeasureService) Save(forecast *model.Forecast) error {

	//TODO: validate
	fmt.Println("-------------------------------------------------------")

	fmt.Println("Antes de busc: ", forecast.EstacionNombre)
	fmt.Println("Antes de busc: ", forecast.CorId)

	var existingForecast model.Forecast

	if forecast.CorId != 0 && forecast.EstacionNombre != "" {
		existingForecast = *s.FindForecast(forecast)
	}

	fmt.Println("Existe?: ", existingForecast.EstacionNombre)
	fmt.Println("Existe?: ", existingForecast.CorId)

	if existingForecast.CorId != 0 && existingForecast.EstacionNombre != "" {
		fmt.Println("Ya existe: ", forecast.EstacionNombre, forecast.SeriesId, forecast.CalId)
		return nil
	}

	if forecast.ForecastSets != nil && len(forecast.ForecastSets) > 0 {
		if err := s.measureRepository.Save(forecast); err != nil {
			return fmt.Errorf("failed to save forecast, error: %w", err)
		}
	}

	return nil
}

func (s *MeasureService) FindForecast(forecast *model.Forecast) *model.Forecast {
	fore, err := s.measureRepository.FindForecast(forecast)
	if err != nil {
		fmt.Errorf("failed to save forecast, error: %w", err)
	}
	return &fore
}

func (s *MeasureService) FindAll() ([]*model.Forecast, error) {
	forecasts, err := s.measureRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to save forecast, error: %w", err)
	}
	return forecasts, err
}

func (s *MeasureService) FindByID(id int) (*model.Forecast, error) {

	return nil, nil
}

func (s *MeasureService) Update(forecast *model.Forecast) error {
	return nil
}

// Returns a new instance of the service.
func NewMeasureService(
	measureRepository repository.ForecastRepository,
) *MeasureService {
	return &MeasureService{
		measureRepository: measureRepository,
	}
}
