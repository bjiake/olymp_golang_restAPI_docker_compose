package weatherAPI

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_project/cmd/internal/config"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/models/region"
	"go_project/cmd/internal/models/weather"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetWeatherById(context *gin.Context) {
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

	idRegion, err := config.GetIntParam(context.Param("regionId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	result, err := requestsToMongoDB.GetWeatherByID(idRegion)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	context.IndentedJSON(http.StatusOK, result[0])
	return
}

func PostWeather(c *gin.Context) {
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

	var newWeather weather.NewWeather
	if err := c.BindJSON(&newWeather); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		log.Println("invalid request")
		return
	}
	if !isValidData(newWeather) {
		c.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}
	var resultWeather weather.Weather

	lastID, err := requestsToMongoDB.GetLastIDByCollection(config.CollectionWeather)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": "internal server error cannot get last id"})
		return
	}

	resultWeather.ID = lastID + 1
	lastRegion, err := requestsToMongoDB.GetRegionByID(newWeather.RegionId)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(409, gin.H{"error": "Cannot find the region"})
		return
	}
	resultWeather.RegionName = lastRegion.Name
	resultWeather.Temperature = newWeather.Temperature
	resultWeather.Humidity = newWeather.Humidity
	resultWeather.WindSpeed = newWeather.WindSpeed
	resultWeather.WeatherCondition = newWeather.WeatherCondition
	resultWeather.PrecipitationAmount = newWeather.PrecipitationAmount
	resultWeather.MeasurementDateTime = newWeather.MeasurementDateTime
	resultWeather.WeatherForecast = newWeather.WeatherForecast

	err = requestsToMongoDB.AddWeather(resultWeather)
	if err != nil {
		if err.Error() == "weather with this WeatherForecastID with region does not exist" {
			c.IndentedJSON(404, "weather with this WeatherForecastID with region does not exist")
			log.Println(err)
			return
		} else {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
	}

	c.JSON(201, resultWeather)
	log.Printf("User created: %+v\n", resultWeather)
}
func isValidData(weather weather.NewWeather) bool {
	validConditions := []string{"CLEAR", "CLOUDY", "RAIN", "SNOW", "FOG", "STORM"}

	if weather.RegionId <= 0 {
		return false
	}

	_, err := time.Parse(time.RFC3339, weather.MeasurementDateTime.Time().Format(time.RFC3339))
	if err != nil {
		return false
	}
	if weather.MeasurementDateTime.Time().IsZero() {
		return false
	}

	if weather.WindSpeed < 0 {
		return false
	}

	if !contains(validConditions, weather.WeatherCondition) {
		return false
	}

	if weather.PrecipitationAmount < 0 {
		return false
	}

	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func PutWeather(context *gin.Context) {
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

	id, err := config.GetIntParam(context.Param("regionId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	var recieveWeather weather.NewWeather

	err = context.BindJSON(&recieveWeather)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": "Invalid request"})
		log.Println(err)
		return
	}
	currentRegion, err := requestsToMongoDB.GetRegionByID(recieveWeather.RegionId)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": "Cannot find the region"})
		log.Println(err)
		return
	}
	var newRegion region.NewRegion
	currentRegion.Name = recieveWeather.RegionName
	err = requestsToMongoDB.PutRegion(recieveWeather.RegionId, newRegion)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": "Cannot find the region"})
		log.Println(err)
		return
	}
	if !isValidData(recieveWeather) || recieveWeather.RegionName == "" || strings.TrimSpace(recieveWeather.RegionName) == "" {
		context.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}
	var resultWeather weather.Weather
	resultWeather.ID = id
	resultWeather.RegionName = recieveWeather.RegionName
	resultWeather.Temperature = recieveWeather.Temperature
	resultWeather.Humidity = recieveWeather.Humidity
	resultWeather.WindSpeed = recieveWeather.WindSpeed
	resultWeather.WeatherCondition = recieveWeather.WeatherCondition
	resultWeather.PrecipitationAmount = recieveWeather.PrecipitationAmount
	resultWeather.MeasurementDateTime = recieveWeather.MeasurementDateTime
	resultWeather.WeatherForecast = recieveWeather.WeatherForecast

	err = requestsToMongoDB.PutWeather(resultWeather)
	if err != nil {
		if err.Error() == "cannot found forecast id on add weather" {
			context.IndentedJSON(404, gin.H{"error": err.Error()})
		} else {
			context.IndentedJSON(404, gin.H{"error": err.Error()})
		}
	}

	context.IndentedJSON(200, resultWeather)
}

func DeleteWeatherByID(context *gin.Context) {
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

	idWeather, err := config.GetIntParam(context.Param("weatherId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	idRegion, err := config.GetIntParam(context.Param("regionId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	err = requestsToMongoDB.DeleteWeatherById(idWeather, idRegion)
	if err != nil {
		context.IndentedJSON(404, gin.H{"message": err.Error()})
		log.Println(err)
		return
	}
	context.IndentedJSON(200, gin.H{})
	log.Println("Region deleted")
	return
}

func SearchWeather(c *gin.Context) {
	var weatherSearch weather.SearchWeather
	startDateTimeString := c.Query("startDateTime")
	startDateTime, err := time.Parse(time.RFC3339, startDateTimeString)
	if err != nil && startDateTimeString != "" {
		log.Println("Invalid start datetime format")
		c.IndentedJSON(400, gin.H{"error": "Invalid start datetime format"})
		return
	} else if startDateTimeString != "" {
		weatherSearch.StartDateTime = primitive.NewDateTimeFromTime(startDateTime)
	} else {
		weatherSearch.StartDateTime = primitive.NewDateTimeFromTime(time.Time{})
	}
	endDateTimeString := c.Query("endDateTime")
	endDateTime, err := time.Parse(time.RFC3339, endDateTimeString)
	if err != nil && endDateTimeString != "" {
		log.Println("Invalid end datetime format")
		c.IndentedJSON(400, gin.H{"error": "Invalid end datetime format"})
		return
	} else if endDateTimeString != "" {
		weatherSearch.EndDateTime = primitive.NewDateTimeFromTime(endDateTime)
	} else {
		weatherSearch.EndDateTime = primitive.NewDateTimeFromTime(time.Time{})
	}

	weatherConditionString := c.Query("weatherCondition")
	if weatherConditionString != "" {
		validConditions := []string{"CLEAR", "CLOUDY", "RAIN", "SNOW", "FOG", "STORM"}
		if !contains(validConditions, weatherConditionString) {
			log.Println("Not a valid condition")
			c.IndentedJSON(400, gin.H{"error": "Not a valid condition"})
			return
		}
		weatherSearch.WeatherCondition = weatherConditionString
	}

	weatherSearch.RegionId, _ = config.GetIntParam(c.Query("regionId"), nil)
	formString := c.Query("form")
	if formString == "" {
		log.Println("invalid form")
		c.IndentedJSON(400, gin.H{"error": "invalid form"})
		return
	}

	regionIdString := c.Query("regionId")
	regionIdInt, err := strconv.ParseInt(regionIdString, 10, 64)
	if err == nil && regionIdInt <= 0 {
		log.Println("invalid regionId")
		c.IndentedJSON(400, gin.H{"error": "invalid regionId"})
		return
	} else {
		weatherSearch.RegionId = regionIdInt
	}

	formInt, err := strconv.ParseInt(formString, 10, 64)
	if err != nil || formInt < 0 {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": "invalid form"})
		return
	}
	weatherSearch.Form = formInt

	sizeInt, err := config.GetIntParam(c.Query("size"), nil)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": "size must be greater than zero"})
		return
	}
	weatherSearch.Size = sizeInt

	accountsInfo, err := requestsToMongoDB.GetSearchWeather(weatherSearch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accountsInfo)
}
