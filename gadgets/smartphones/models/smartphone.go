package models

// Smartphone is a struct that represents a smartphone
type Smartphone struct {
	Id            int64
	Name          string
	Price         int
	CountryOrigin string
	Os            string
}

// CreateSmartphoneCMD is a struct that represents a smartphone
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	Os            string `json:"os"`
}
