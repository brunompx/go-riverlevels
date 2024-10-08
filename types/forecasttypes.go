package types

import (
	"time"
)

// Table: forecasts
type Forecast struct {
	ID             int           `gorm:"primaryKey;autoIncrement;column:id"`
	VarId          int           `gorm:"column:var_id;not null"`
	CorId          int           `gorm:"column:cor_id;not null"`
	SeriesId       int           `gorm:"column:series_id;not null"`
	UnitId         int           `gorm:"column:unit_id;not null"`
	SiteCode       int           `gorm:"column:site_code;not null"`
	EstacionNombre string        `gorm:"column:estacion_nombre;not null"`
	CalId          int           `gorm:"column:cal_id;not null"`
	VarNombre      string        `gorm:"column:var_nombre;not null"`
	ModelId        int           `gorm:"column:model_id;not null"`
	ForecastDate   time.Time     `gorm:"column:forecast_date;not null"`
	ForecastSets   []ForecastSet `gorm:"foreignKey:ForecastID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-many relationship with ForecastSet
}

// Table: forecast_sets
type ForecastSet struct {
	ID             int             `gorm:"primaryKey;autoIncrement;column:id"`
	ForecastID     int             `gorm:"column:forecast_id;not null"`                                           // Foreign key to Forecast
	Forecast       Forecast        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`                          // Belongs to Forecast
	ForecastLevels []ForecastLevel `gorm:"foreignKey:ForecastSetID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // One-to-many relationship with ForecastLevel
}

// Table: forecast_levels
type ForecastLevel struct {
	ID            int         `gorm:"primaryKey;autoIncrement;column:id"`
	ForecastSetID int         `gorm:"column:forecast_set_id;not null"`              // Foreign key to ForecastSet
	ForecastSet   ForecastSet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Belongs to ForecastSet
	PronoID       int64       `gorm:"column:prono_id;not null"`                     // Use int64 for large integers
	TimeProno     time.Time   `gorm:"column:time_prono;not null"`
	Valor         float64     `gorm:"column:valor;not null"`
}

// Struct for responseHeader in the forecast json
type ForecastResponseHeader struct {
	ID                int    `json:"id"` // Primary key
	VarId             int    `json:"varid"`
	Request           string `json:"request"`
	EstacionTabla     string `json:"estacion_tabla"`
	CorId             int    `json:"corid"`
	TimeEnd           string `json:"timeEnd"`
	SeriesId          int    `json:"seriesid"`
	ProcNombre        string `json:"proc_nombre"`
	UnitId            int    `json:"unitid"`
	SiteCode          int    `json:"sitecode"`
	EstacionNombre    string `json:"estacion_nombre"`
	RedNombre         string `json:"red_nombre"`
	CalId             int    `json:"calid"`
	CalName           string `json:"cal_name"`
	ResponseTimestamp string `json:"responseTimestamp"`
	TimeStart         string `json:"timeStart"`
	VarNombre         string `json:"var_nombre"`
	CalModel          string `json:"cal_model"`
	QueryUrl          string `json:"queryUrl"`
	ModelId           int    `json:"model_id"`
	CreationTime      string `json:"creationTime"`
	UnitNombre        string `json:"unit_nombre"`
	ForecastDate      string `json:"forecastdate"`
}

// Struct for data entries in the forecast json
type ForecastDataEntry struct {
	ID        int     `json:"id"` // Primary key
	PronoId   int64   `json:"prono_id"`
	TimeStart string  `json:"timestart"`
	TimeEnd   string  `json:"timeend"`
	Valor     float64 `json:"valor"`
}

// Struct that encompasses the whole forecast JSON structure for unmarshalling
type ForecastResponse struct {
	ResponseHeader ForecastResponseHeader `json:"responseHeader"`
	Data           []ForecastDataEntry    `json:"data"`
}
