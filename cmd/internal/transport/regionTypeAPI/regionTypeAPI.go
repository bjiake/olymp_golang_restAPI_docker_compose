package regionTypeAPI

import (
	"github.com/gin-gonic/gin"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"log"
	"net/http"
	"strconv"
)

// GetRegionTypeByID invalid acc to add
func GetRegionTypeByID(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id value"})
		log.Println(err)
		return
	}
	result, err := requestsToMongoDB.GetRegionTypeByID(idInt)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	context.IndentedJSON(http.StatusOK, result)
	return
}
