package weatherForecast

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeatherForecast struct {
	ID               int64              `json:"id" bson:"_id"`
	DateTime         primitive.DateTime `json:"dateTime"`
	Temperature      float64            `json:"temperature"`
	WeatherCondition string             `json:"weatherCondition"`
	RegionID         int64              `json:"regionId"`
}
