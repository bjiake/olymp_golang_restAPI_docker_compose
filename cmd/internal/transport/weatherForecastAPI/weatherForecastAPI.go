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

func PutWeatherForecast(context *gin.Context) {
	loginId, err := config.GetIntParam(context.Cookie("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "invalid login"})
		log.Println("invalid login")
		return
	}
	_, err = requestsToMongoDB.GetAccountByID(loginId)
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "this account does not exist"})
		log.Println("this account does not exist")
		return
	}

	id, err := config.GetIntParam(context.Param("forecastId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	var updatedDoc weatherForecast.NewPutWeatherForecast

	err = context.BindJSON(&updatedDoc)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": "Invalid request"})
		log.Println(err)
		return
	}
	if !isValidPut(updatedDoc) {
		context.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}

	err = requestsToMongoDB.PutWeatherForecast(id, updatedDoc)
	if err != nil {
		if err.Error() == "Region with this Latitude and Longitude already exists" {
			context.IndentedJSON(409, gin.H{"error": err.Error()})
		} else {
			context.IndentedJSON(404, gin.H{"error": err.Error()})
		}
	}
	result, _ := requestsToMongoDB.GetWeatherForecastByID(id)

	context.IndentedJSON(200, result)
}

func PostWeatherByID(c *gin.Context) {
	loginIdInt, err := config.GetIntParam(c.Cookie("id"))
	if err != nil {
		log.Println("error invalid login")
		c.IndentedJSON(401, gin.H{"error": "invalid login"})
		return
	}
	_, err = requestsToMongoDB.GetAccountByID(loginIdInt)
	if err != nil {
		c.IndentedJSON(401, gin.H{"error": "this account does not exist"})
		log.Println("this account does not exist")
		return
	}

	var newWeatherForecast weatherForecast.NewPostWeatherForecast
	if err := c.BindJSON(&newWeatherForecast); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		log.Println("invalid request")
		return
	}
	if !isValidPost(newWeatherForecast) {
		c.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}
	var resultWeatherForecast weatherForecast.WeatherForecast

	lastID, err := requestsToMongoDB.GetLastIDByCollection(config.CollectionWeatherForecast)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": "internal server error cannot get last id"})
		return
	}

	resultWeatherForecast.ID = lastID + 1
	resultWeatherForecast.RegionID = newWeatherForecast.RegionID
	resultWeatherForecast.WeatherCondition = newWeatherForecast.WeatherCondition
	resultWeatherForecast.DateTime = newWeatherForecast.DateTime
	resultWeatherForecast.Temperature = newWeatherForecast.Temperature

	err = requestsToMongoDB.AddWeatherForecast(resultWeatherForecast)
	if err != nil {
		if err.Error() == "region with this id not found" {
			c.IndentedJSON(404, "region with this id not found")
			log.Println(err)
			return
		} else {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
	}

	c.JSON(200, resultWeatherForecast)
	log.Printf("User created: %+v\n", resultWeatherForecast)
}

func DeleteRegionTypeByID(context *gin.Context) {
	loginIdInt, err := config.GetIntParam(context.Cookie("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	_, err = requestsToMongoDB.GetAccountByID(loginIdInt)
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "this account does not exist"})
		log.Println("this account does not exist")
		return
	}
	idInt, err := config.GetIntParam(context.Param("forecastId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	err = requestsToMongoDB.DeleteByIDAndCollection(idInt, config.CollectionWeatherForecast)
	if err != nil {
		context.IndentedJSON(404, gin.H{"message": err.Error()})
		log.Println(err)
		return
	}
	context.IndentedJSON(200, gin.H{})
	log.Println("WeatherForecast deleted")
	return
}

func isValidPut(n weatherForecast.NewPutWeatherForecast) bool {
	// Check if dateTime is in ISO-8601 format
	_, err := time.Parse(time.RFC3339, n.DateTime.Time().Format(time.RFC3339))
	if err != nil {
		return false
	}

	// Check if weatherCondition is not one of the allowed values
	allowedConditions := map[string]bool{
		"CLEAR":   true,
		"CLOUDY":  true,
		"RAIN":    true,
		"SNOW":    true,
		"FOG":     true,
		"STORM":   true,
		"UNKNOWN": true, // Add "UNKNOWN" as an allowed value for this example
	}
	if _, ok := allowedConditions[n.WeatherCondition]; !ok {
		return false
	}

	return true
}

func isValidPost(n weatherForecast.NewPostWeatherForecast) bool {
	// Check if dateTime is in ISO-8601 format
	_, err := time.Parse(time.RFC3339, n.DateTime.Time().Format(time.RFC3339))
	if err != nil {
		return false
	}
	if n.RegionID <= 0 {
		return false
	}

	// Check if weatherCondition is not one of the allowed values
	allowedConditions := map[string]bool{
		"CLEAR":   true,
		"CLOUDY":  true,
		"RAIN":    true,
		"SNOW":    true,
		"FOG":     true,
		"STORM":   true,
		"UNKNOWN": true, // Add "UNKNOWN" as an allowed value for this example
	}
	if _, ok := allowedConditions[n.WeatherCondition]; !ok {
		return false
	}

	return true
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
