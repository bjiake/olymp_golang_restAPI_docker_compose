package weatherforecast

import (
	"github.com/gin-gonic/gin"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/models/weatherForecast"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetWeatherForecastByID(context *gin.Context) {
	id := context.Param("forecastId")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id value"})
		log.Println(err)
		return
	}
	result, err := requestsToMongoDB.GetWeatherForecastByID(idInt)
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
