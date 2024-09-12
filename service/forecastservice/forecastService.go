package taskservice

import (
	"fmt"

	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository"
)

type ForecastService struct {
	forecastRepository repository.ForecastRepository // interface of the repository, not implementation
}

func (s *ForecastService) Save(forecast *model.Forecast) error {
	//if err := s.taskValidatorService.Validate(ctx, task); err != nil {
	//	return err
	//}

	if err := s.forecastRepository.Save(forecast); err != nil {
		return fmt.Errorf("failed to save forecast, error: %w", err)
	}

	return nil
}

func (s *ForecastService) FindAll() ([]*model.Forecast, error) {
	forecasts, err := s.forecastRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to save forecast, error: %w", err)
	}
	return forecasts, err
}

func (s *ForecastService) FindByID(id int) (*model.Forecast, error) {

	return nil, nil
}

func (s *ForecastService) Update(forecast *model.Forecast) error {
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
