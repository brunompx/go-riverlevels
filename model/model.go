package model

import "time"

type Location struct {
	SeriesId string `json:"seriesId"`
	SiteCode string `json:"siteCode"`
	CalId    string `json:"calId"`
	CorId    string `json:"corId"`
	VarId    string `json:"varId"`
}

type Locations struct {
	Locations []Location `json:"locations"`
}

type Forecast struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PronoID   int64     `json:"prono_id"` // Use int64 for big integers
	TimeProno time.Time `json:"time_start"`
	Valor     float64   `json:"valor"`
}
