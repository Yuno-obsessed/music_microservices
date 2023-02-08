package entity

type City struct {
	Id        string  `json:"city_id"`
	CountryId Country `json:"country_id"`
	CityName  string  `json:"city_name"`
}
