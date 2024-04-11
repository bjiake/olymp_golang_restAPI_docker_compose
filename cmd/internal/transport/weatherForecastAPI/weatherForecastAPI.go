package weatherforecast

import (
	"github.com/gin-gonic/gin"
	"go_project/cmd/internal/config"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/models/weatherForecast"
	"log"
	"net/http"
	"time"
)

func GetWeatherForecastByID(context *gin.Context) {
	loginIdInt, err := config.GetIntParam(context.Cookie("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "invalid login"})
		log.Println(err)
		return
	}
	_, err = requestsToMongoDB.GetAccountByID(loginIdInt)
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "this account does not exist"})
		log.Println("this account does not exist")
		return
	}

	regionId, err := config.GetIntParam(context.Param("forecastId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	result, err := requestsToMongoDB.GetWeatherForecastByID(regionId)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	if !checkWeatherCondition(result) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		log.Println("invalid data")
		return
	}

	context.IndentedJSON(http.StatusOK, result)
	return
}

func checkWeatherCondition(weatherForecast weatherForecast.WeatherForecast) bool {
	var result = false
	weatherConditions := [...]string{"CLEAR", "CLOUDY", "RAIN", "SNOW", "FOG", "STORM"}
	for _, item := range weatherConditions {
		if weatherForecast.WeatherCondition == item {
			result = true
			break
		}
	}
	return result
}

func checkDateISO(date string) bool {
	result := true
	_, err := time.Parse(time.RFC3339, date)
	if err != nil {
		result = false
	}
	return result
}
