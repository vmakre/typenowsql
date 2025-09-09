package models

import "time"

type City struct {
	City_id     int       `json:"cityid"`
	City        string    `json:"city_name"`
	Last_update time.Time `json:"last_update"`
	Country_id  int       `json:"country_id"`
}

type CreateCityRequest struct {
	City string `json:"city_name" validate:"required"`
	//Email string `json:"email" validate:"required,email"`
}
