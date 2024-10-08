package service

import (
	"fmt"

	"github.com/brunompx/go-riverlevels/repository"
	"github.com/brunompx/go-riverlevels/types"
)

type ForecastService struct {
	forecastRepository repository.ForecastRepository // interface of the repository, not implementation
}

func (s *ForecastService) Save(forecast *types.Forecast) error {

	//TODO: validate
	fmt.Println("-------------------------------------------------------")

	fmt.Println("Antes de busc: ", forecast.EstacionNombre)
	fmt.Println("Antes de busc: ", forecast.CorId)

	var existingForecast types.Forecast

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
		if err := s.forecastRepository.Save(forecast); err != nil {
			return fmt.Errorf("failed to save forecast, error: %w", err)
		}
	}

	return nil
}

func (s *ForecastService) FindForecast(forecast *types.Forecast) *types.Forecast {
	fore, err := s.forecastRepository.FindForecast(forecast)
	if err != nil {
		fmt.Errorf("failed to save forecast, error: %w", err)
	}
	return &fore
}

func (s *ForecastService) FindAll() ([]*types.Forecast, error) {
	forecasts, err := s.forecastRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to save forecast, error: %w", err)
	}
	return forecasts, err
}

func (s *ForecastService) FindByID(id int) (*types.Forecast, error) {

	return nil, nil
}

func (s *ForecastService) Update(forecast *types.Forecast) error {
	return nil
}

// Returns a new instance of the service.
func NewForecastService(
	forecastRepository repository.ForecastRepository,
) *ForecastService {
	return &ForecastService{
		forecastRepository: forecastRepository,
	}
}
