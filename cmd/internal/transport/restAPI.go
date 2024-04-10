package transport

import (
	"github.com/gin-gonic/gin"
	"go_project/cmd/internal/transport/accountAPI"
	regionTypeAPI "go_project/cmd/internal/transport/regionTypeAPI"
	weatherForecastAPI "go_project/cmd/internal/transport/weatherForecastAPI"
)

func Transport() {
	router := gin.Default()

	router.POST("/registration", accountAPI.Registration)
	router.POST("/login", accountAPI.Login)
	router.GET("/accounts/:accountId", accountAPI.GetAccountByID)
	router.PUT("/accounts/:accountId", accountAPI.PutAccountByID)
	router.GET("/region/type/:id", regionTypeAPI.GetRegionTypeByID)
	router.GET("/region/weather/forecast/:forecastId", weatherForecastAPI.GetWeatherForecastByID)

	router.DELETE("/accounts/:accountId", accountAPI.DeleteAccountByID)

	router.Run(":8088")
}
