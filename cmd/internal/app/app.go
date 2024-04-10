package app

import (
	"go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/transport"
)

func Start() {
	requestsToMongoDB.GetRegionByID(2)
	requestsToMongoDB.GetRegionTypeByID(3)
	requestsToMongoDB.GetWeatherByID(4)
	requestsToMongoDB.GetWeatherForecastByID(2)
	requestsToMongoDB.GetAccountByID(2)
	transport.Transport()
	requestsToMongoDB.CloseConnection()
}
