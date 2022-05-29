package entities

type City struct {
	City string `json:"city_name"`
}
type CityResponse struct {
	ID   uint   `json:"id"`
	City string `json:"city_name"`
}
type AddCity struct {
	City string `json:"city_name" validate:"required"`
}
