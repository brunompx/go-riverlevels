package types

import (
	"fmt"
	"time"
)

func (response *ForecastResponse) NormalizeToForecast() Forecast {
	if len(response.Data) == 0 {
		fmt.Println("SIN DATOS!!")
		return Forecast{}
	}
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
		ForecastDate:   stringToTime(response.ResponseHeader.ForecastDate),
	}
	forecastSets := []ForecastSet{}
	for _, entry := range response.Data {
		forecastLevel := ForecastLevel{
			PronoID:   entry.PronoId,
			Valor:     entry.Valor,
			TimeProno: stringToTime(entry.TimeStart),
		}
		added := false
	setrange:
		for i, set := range forecastSets {
		levelrange:
			for _, level := range set.ForecastLevels {
				if forecastLevel.PronoID == level.PronoID+1 ||
					forecastLevel.PronoID == level.PronoID-1 {
					forecastSets[i].ForecastLevels = append(forecastSets[i].ForecastLevels, forecastLevel)
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

func stringToTime(stringValue string) time.Time {
	parsed, err := time.Parse("2006-01-02T15:04:05", stringValue)
	if err != nil {
		fmt.Println("Error parsing string date: ", err)
	}
	return parsed
}

func (response *MeasureResponse) NormalizeToMeasure() Measure {
	if len(response.Data) == 0 {
		fmt.Println("SIN DATOS!!")
		return Measure{}
	}
	measure := Measure{
		SeriesId:          response.ResponseHeader.SeriesID,
		SiteCode:          response.ResponseHeader.SiteCode,
		EstacionAbrev:     response.ResponseHeader.SiteMetadata.EstacionAbrev,
		ResponseTimestamp: stringToTime(response.ResponseHeader.ResponseTimestamp),
		RedId:             response.ResponseHeader.SiteMetadata.RedID,
	}
	measureLevels := []MeasureLevel{}
	for _, entry := range response.Data {
		measureLevel := MeasureLevel{
			ObsId:       entry.ObsId,
			Valor:       entry.Valor,
			TimeMeasure: stringToTime(entry.TimeStart),
		}
		measureLevels = append(measureLevels, measureLevel)
	}
	measure.MeasureLevels = measureLevels

	return measure
}
