package model

func (response ForecastResponse) normalizeToForecast() Forecast {
	forecast := Forecast{
		VarId:          response.ResponseHeader.VarId,
		CorId:          response.ResponseHeader.CorId,
		SeriesId:       response.ResponseHeader.SeriesId,
		UnitId:         response.ResponseHeader.UnitId,
		SiteCode:       response.ResponseHeader.SiteCode,
		EstacionNombre: response.ResponseHeader.EstacionNombre,
		CalId:          response.ResponseHeader.CalId,
		VarNombre:      response.ResponseHeader.VarNombre,
		ModelId:        response.ResponseHeader.ModelId,
		ForecastDate:   response.ResponseHeader.ForecastDate,
	}
	forecastSets := []ForecastSet{}
	for _, entry := range response.Data {
		forecastLevel := ForecastLevel{
			PronoID:   entry.PronoId,
			Valor:     entry.Valor,
			TimeProno: entry.TimeStart,
		}
		added := false
	setrange:
		for _, set := range forecastSets {
		levelrange:
			for _, level := range set.ForecastLevels {
				if forecastLevel.PronoID == level.PronoID+1 ||
					forecastLevel.PronoID == level.PronoID-1 {
					set.ForecastLevels = append(set.ForecastLevels, forecastLevel)
					added = true
					break levelrange
				}
			}
			if added {
				break setrange
			}
		}
		if !added {
			forecastLevels := []ForecastLevel{forecastLevel}
			forecastSet := ForecastSet{
				ForecastLevels: forecastLevels,
			}
			forecastSets = append(forecastSets, forecastSet)
		}
	}
	forecast.ForecastSets = forecastSets
	return forecast
}
