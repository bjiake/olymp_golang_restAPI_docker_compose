package weatherAPI

import (
	"github.com/gin-gonic/gin"
	"go_project/cmd/internal/config"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"log"
	"net/http"
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

	context.IndentedJSON(http.StatusOK, result)
	return
}
