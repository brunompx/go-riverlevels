package types

import "time"

// Table: forecasts
type Measure struct {
	ID                int            `gorm:"primaryKey;autoIncrement;column:id"`
	SeriesId          int            `gorm:"column:series_id;not null"`
	SiteCode          int            `gorm:"column:site_code;not null"`
	EstacionAbrev     string         `gorm:"column:estacion_abrev;not null"`
	ResponseTimestamp time.Time      `gorm:"column:response_timestamp;not null"`
	RedId             int            `gorm:"column:red_id;not null"`
	MeasureLevels     []MeasureLevel `gorm:"foreignKey:MeasureID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-many relationship with MeasureLevel
}

// Table: forecast_levels
type MeasureLevel struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id"`
	MeasureID   int       `gorm:"column:measure_id;not null"`                   // Foreign key to Measure
	Measure     Measure   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Belongs to Measure
	ObsId       int64     `gorm:"column:obs_id;not null"`                       // Use int64 for large integers
	TimeMeasure time.Time `gorm:"column:time_measure;not null"`
	Valor       float64   `gorm:"column:valor;not null"`
}

// Struct for responseHeader in the forecast json
type MeasureResponseHeader struct {
	SiteCode          int            `json:"sitecode"`
	CreationTime      string         `json:"creationTime"`
	QueryUrl          string         `json:"queryUrl"`
	ResponseTimestamp string         `json:"responseTimestamp"`
	TimeEnd           string         `json:"timeEnd"`
	Request           string         `json:"request"`
	SeriesMetadata    SeriesMetadata `json:"seriesmetadata"`
	SeriesID          int            `json:"seriesid"`
	TimeStart         string         `json:"timeStart"`
	SiteMetadata      SiteMetadata   `json:"sitemetadata"`
}

// SeriesMetadata struct
type SeriesMetadata struct {
	VarAbrev  string `json:"var_abrev"`
	ProcID    int    `json:"procId"`
	UnitAbrev string `json:"unit_abrev"`
	UnitID    int    `json:"unitId"`
	SeriesID  int    `json:"seriesId"`
	ProcAbrev string `json:"proc_abrev"`
	VarID     int    `json:"varId"`
}

// SiteMetadata struct
type SiteMetadata struct {
	RedName       string `json:"red_name"`
	EstacionAbrev string `json:"estacion_abrev"`
	RedID         int    `json:"redId"`
	SiteCode      int    `json:"siteCode"`
}

// Struct for data entries in the forecast json
type MeasureDataEntry struct {
	ID        int     `json:"id"` // Primary key
	ObsId     int64   `json:"obs_id"`
	TimeStart string  `json:"timestart"`
	TimeEnd   string  `json:"timeend"`
	Valor     float64 `json:"valor"`
}

// Struct that encompasses the whole forecast JSON structure for unmarshalling
type MeasureResponse struct {
	ResponseHeader MeasureResponseHeader `json:"responseHeader"`
	Data           []MeasureDataEntry    `json:"data"`
}
