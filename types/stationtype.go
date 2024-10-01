package types

type Station struct {
	StationType string `json:"Type"`
	SeriesId    string `json:"seriesId"`
	SiteCode    string `json:"siteCode"`
	CalId       string `json:"calId"`
	CorId       string `json:"corId"`
	VarId       string `json:"varId"`
	Format      string `json:"format"`
	All         string `json:"all"`
}

type Stations struct {
	Stations []Station `json:"locations"`
}
