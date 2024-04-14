package weatherForecast

import "go.mongodb.org/mongo-driver/bson/primitive"

type WeatherForecast struct {
	ID               int64              `json:"id" bson:"_id" validate:"required"`
	DateTime         primitive.DateTime `json:"dateTime" bson:"dateTime" validate:"required"`
	Temperature      float64            `json:"temperature" bson:"temperature" validate:"required"`
	WeatherCondition string             `json:"weatherCondition" bson:"weatherCondition" validate:"required"`
	RegionID         int64              `json:"regionId" bson:"regionId" validate:"required"`
}

type NewPutWeatherForecast struct {
	DateTime         primitive.DateTime `json:"dateTime" bson:"dateTime" validate:"required"`
	Temperature      float64            `json:"temperature" bson:"temperature" validate:"required"`
	WeatherCondition string             `json:"weatherCondition" bson:"weatherCondition" validate:"required"`
}

type NewPostWeatherForecast struct {
	DateTime         primitive.DateTime `json:"dateTime" bson:"dateTime" validate:"required"`
	Temperature      float64            `json:"temperature" bson:"temperature" validate:"required"`
	RegionID         int64              `json:"regionId" bson:"regionId" validate:"required"`
	WeatherCondition string             `json:"weatherCondition" bson:"weatherCondition" validate:"required"`
}
