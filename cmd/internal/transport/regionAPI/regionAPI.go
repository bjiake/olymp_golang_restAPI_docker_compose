package regionAPI

import (
	"github.com/gin-gonic/gin"
	"go_project/cmd/internal/config"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/models/region"
	"log"
	"net/http"
	"strings"
)

func GetRegionByID(context *gin.Context) {
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

	result, err := requestsToMongoDB.GetRegionByID(idRegion)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	context.IndentedJSON(http.StatusOK, result)
	return
}

func PostRegion(c *gin.Context) {
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

	var newRegion region.NewRegion
	if err := c.BindJSON(&newRegion); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		log.Println("invalid request")
		return
	}
	if !isValidData(newRegion) {
		c.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}
	var resultRegion region.Region

	lastID, err := requestsToMongoDB.GetLastIDByCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": "internal server error cannot get last id"})
		return
	}

	resultRegion.ID = lastID + 1
	resultRegion.RegionType = newRegion.RegionType
	resultRegion.ParentRegionId = newRegion.ParentRegionId
	resultRegion.Name = newRegion.Name
	resultRegion.Latitude = newRegion.Latitude
	resultRegion.Longitude = newRegion.Longitude
	resultRegion.AccountId = loginIdInt

	err = requestsToMongoDB.AddRegion(resultRegion)
	if err != nil {
		if err.Error() == "Region with this Latitude and Longitude already exists" {
			c.IndentedJSON(409, "Region with this Latitude and Longitude already exists")
			log.Println(err)
			return
		} else {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
	}

	c.JSON(201, gin.H{"message": "Region created"})
	log.Printf("User created: %+v\n", resultRegion)
}

func PutRegionByID(context *gin.Context) {
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

	var updatedDoc region.NewRegion

	err = context.BindJSON(&updatedDoc)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": "Invalid request"})
		log.Println(err)
		return
	}
	if !isValidData(updatedDoc) {
		context.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}

	err = requestsToMongoDB.PutRegion(id, updatedDoc)
	if err != nil {
		if err.Error() == "Region with this Latitude and Longitude already exists" {
			context.IndentedJSON(409, gin.H{"error": err.Error()})
		} else {
			context.IndentedJSON(404, gin.H{"error": err.Error()})
		}
	}
	infoRegion, _ := requestsToMongoDB.GetRegionByID(id)

	context.IndentedJSON(200, infoRegion)
}

func isValidData(newRegion region.NewRegion) bool {
	// 0 т.к в golang Нету nil Для float64
	if newRegion.Latitude == 0 || newRegion.Longitude == 0 || newRegion.Name == "" || strings.TrimSpace(newRegion.Name) == "" {
		return false
	}
	return true
}

func DeleteRegionByID(context *gin.Context) {
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
	idInt, err := config.GetIntParam(context.Param("regionId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	err = requestsToMongoDB.DeleteRegionById(idInt)
	if err != nil && err.Error() == "This region does not exist" {
		context.IndentedJSON(404, gin.H{"error": "this region does not exist"})
		log.Println("this region does not exist")
		return
	} else if err != nil && err.Error() == "This region is parent region for someone" {
		context.IndentedJSON(400, gin.H{"error": "This region is parent region for someone"})
		log.Println("This region is parent region for someone")
		return
	}
	context.IndentedJSON(200, gin.H{})
	log.Println("Region deleted")
}
