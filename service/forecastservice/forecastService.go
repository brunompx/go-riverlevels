package taskservice

import (
	"fmt"

	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository"
)

type ForecastService struct {
	forecastRepository repository.ForecastRepository // interface of the repository, not implementation
}

func (s *ForecastService) Create(forecast *model.Forecast) error {
	//if err := s.taskValidatorService.Validate(ctx, task); err != nil {
	//	return err
	//}

	if err := s.forecastRepository.Save(forecast); err != nil {
		return fmt.Errorf("failed to save forecast, error: %w", err)
	}

	return nil
}
