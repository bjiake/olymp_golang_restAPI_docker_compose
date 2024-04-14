package weather

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Weather struct {
	ID                  int64              `json:"id" validate:"required" bson:"_id"`
	RegionName          string             `json:"regionName" bson:"regionName" validate:"required"`
	Temperature         float64            `json:"temperature" bson:"temperature" validate:"required"`
	Humidity            float64            `json:"humidity" bson:"humidity" validate:"required"`
	WindSpeed           float64            `json:"windSpeed" bson:"windSpeed" validate:"required"`
	WeatherCondition    string             `json:"weatherCondition" bson:"weatherCondition" validate:"required"`
	PrecipitationAmount float64            `json:"precipitationAmount" bson:"precipitationAmount" validate:"required"`
	MeasurementDateTime primitive.DateTime `json:"measurementDateTime" bson:"measurementDateTime" validate:"required"`
	WeatherForecast     []int64            `json:"weatherForecast" bson:"weatherForecast" validate:"required"`
}

type NewWeather struct {
	RegionId            int64              `json:"regionId" bson:"_id" validate:"required"`
	RegionName          string             `json:"regionName" validate:"required"`
	Temperature         float64            `json:"temperature" bson:"temperature" validate:"required"`
	Humidity            float64            `json:"humidity" bson:"humidity" validate:"required"`
	WindSpeed           float64            `json:"windSpeed" bson:"windSpeed" validate:"required"`
	WeatherCondition    string             `json:"weatherCondition" bson:"weatherCondition" validate:"required"`
	PrecipitationAmount float64            `json:"precipitationAmount" bson:"precipitationAmount" validate:"required"`
	MeasurementDateTime primitive.DateTime `json:"measurementDateTime" bson:"measurementDateTime" validate:"required"`
	WeatherForecast     []int64            `json:"weatherForecast" bson:"weatherForecast" validate:"required"`
}

type SearchWeather struct {
	StartDateTime    primitive.DateTime `json:"startDateTime"`
	EndDateTime      primitive.DateTime `json:"endDateTime"`
	RegionId         int64              `json:"regionId"`
	WeatherCondition string             `json:"weatherCondition"`
	Form             int64              `json:"form"`
	Size             int64              `json:"size"`
}
