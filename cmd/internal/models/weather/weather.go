package weather

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Weather struct {
	ID                  int64              `json:"id" bson:"_id"`
	RegionName          string             `json:"regionName"`
	Temperature         float64            `json:"temperature"`
	Humidity            float64            `json:"humidity"`
	WindSpeed           float64            `json:"windSpeed"`
	WeatherCondition    string             `json:"weatherCondition"`
	PrecipitationAmount float64            `json:"precipitationAmount"`
	MeasurementDateTime primitive.DateTime `json:"measurementDateTime"`
	WeatherForecast     []int64            `json:"weatherForecastAPI" bson:"weatherForecastAPI"`
}
