package transport

import (
	"github.com/gin-gonic/gin"
	"go_project/cmd/internal/transport/accountAPI"
	"go_project/cmd/internal/transport/regionAPI"
	regionTypeAPI "go_project/cmd/internal/transport/regionTypeAPI"
	"go_project/cmd/internal/transport/weatherAPI"
	weatherForecastAPI "go_project/cmd/internal/transport/weatherForecastAPI"
	"log"
)

func Transport() {
	router := gin.Default()

	//Методы для account
	router.POST("/registration", accountAPI.Registration)
	router.POST("/login", accountAPI.Login)
	router.GET("/accounts/:accountId", accountAPI.GetAccountByID)
	router.PUT("/accounts/:accountId", accountAPI.PutAccountByID)
	router.DELETE("/accounts/:accountId", accountAPI.DeleteAccountByID)

	//Методы для Region
	router.GET("/region/:regionId", regionAPI.GetRegionByID)
	router.POST("/region", regionAPI.PostRegion)
	router.PUT("/region/:regionId", regionAPI.PutRegionByID)
	router.DELETE("region/:regionId", regionAPI.DeleteRegionByID)

	//RegionType
	router.GET("/region/type/:typeId", regionTypeAPI.GetRegionTypeByID)

	//Weather
	router.GET("/region/weather/:regionId", weatherAPI.GetWeatherById)

	//WeatherForecast
	router.GET("/region/weather/forecast/:forecastId", weatherForecastAPI.GetWeatherForecastByID)

	err := router.Run(":8088")
	if err != nil {
		log.Fatal(err)
		return
	}
}
