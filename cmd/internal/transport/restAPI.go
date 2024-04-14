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
	router.GET("/accounts/search", accountAPI.SearchAccount)

	//Методы для Region
	router.GET("/region/:regionId", regionAPI.GetRegionByID)
	router.POST("/region", regionAPI.PostRegion)
	router.PUT("/region/:regionId", regionAPI.PutRegionByID)
	router.DELETE("region/:regionId", regionAPI.DeleteRegionByID)

	//RegionType
	router.GET("/region/types/:typeId", regionTypeAPI.GetRegionTypeByID)
	router.POST("/region/types", regionTypeAPI.PostRegionType)
	router.PUT("/region/types/:typeId", regionTypeAPI.PutRegionTypeByID)
	router.DELETE("region/types/:typeId", regionTypeAPI.DeleteRegionTypeByID)

	//Weather
	router.GET("/region/weather/:regionId", weatherAPI.GetWeatherById)
	router.POST("/region/weather", weatherAPI.PostWeather)
	router.PUT("/region/weather/:regionId", weatherAPI.PutWeather)
	router.DELETE("/region/:regionId/weather/:weatherId", weatherAPI.DeleteWeatherByID)

	//тестировать
	router.GET("/region/weather/search", weatherAPI.SearchWeather)

	//WeatherForecast
	router.GET("/region/weather/forecast/:forecastId", weatherForecastAPI.GetWeatherForecastByID)
	router.PUT("/region/weather/forecast/:forecastId", weatherForecastAPI.PutWeatherForecast)
	router.POST("/region/weather/forecast", weatherForecastAPI.PostWeatherByID)
	router.DELETE("/region/weather/forecast/:forecastId", weatherForecastAPI.DeleteRegionTypeByID)

	err := router.Run(":8088")
	if err != nil {
		log.Fatal(err)
		return
	}
}
