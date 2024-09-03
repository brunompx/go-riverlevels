package model

type Location struct {
	SeriesId string `json:"seriesId"`
	SiteCode string `json:"siteCode"`
	CalId    string `json:"calId"`
}

type Locations struct {
	Locations []Location `json:"locations"`
}
