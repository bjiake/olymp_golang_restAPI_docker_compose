package regionTypeAPI

import (
	"github.com/gin-gonic/gin"
	"go_project/cmd/internal/config"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/models/regionType"
	"log"
	"net/http"
	"strings"
)

// GetRegionTypeByID invalid acc to add
func GetRegionTypeByID(context *gin.Context) {
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

	idRegion, err := config.GetIntParam(context.Param("typeId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	result, err := requestsToMongoDB.GetRegionTypeByID(idRegion)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	context.IndentedJSON(http.StatusOK, result)
	return
}

func PostRegionType(c *gin.Context) {
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

	var newRegion regionType.NewRegionType
	if err := c.BindJSON(&newRegion); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		log.Println("invalid request")
		return
	}
	if newRegion.Type == "" || strings.TrimSpace(newRegion.Type) == "" {
		c.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}
	var resultRegion regionType.RegionType

	lastID, err := requestsToMongoDB.GetLastIDByCollection(config.CollectionRegionType)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": "internal server error cannot get last id"})
		return
	}

	resultRegion.ID = lastID + 1
	resultRegion.Type = newRegion.Type

	err = requestsToMongoDB.AddRegionType(resultRegion)
	if err != nil {
		if err.Error() == "Region with this Type already exists" {
			c.IndentedJSON(409, "Region with this Type already exists")
			log.Println(err)
			return
		} else {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
	}

	c.JSON(201, resultRegion)
	log.Printf("User created: %+v\n", resultRegion)
}

func PutRegionTypeByID(context *gin.Context) {
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

	id, err := config.GetIntParam(context.Param("typeId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	var updatedDoc regionType.NewRegionType

	err = context.BindJSON(&updatedDoc)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": "Invalid request"})
		log.Println(err)
		return
	}
	if updatedDoc.Type == "" || strings.TrimSpace(updatedDoc.Type) == "" {
		context.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}

	err = requestsToMongoDB.PutRegionType(id, updatedDoc)
	if err != nil {
		if err.Error() == "Region with this Type already exists" {
			context.IndentedJSON(409, gin.H{"error": err.Error()})
		} else {
			context.IndentedJSON(404, gin.H{"error": err.Error()})
		}
	}
	info, _ := requestsToMongoDB.GetRegionTypeByID(id)

	context.IndentedJSON(200, info)
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
	idInt, err := config.GetIntParam(context.Param("typeId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	err = requestsToMongoDB.DeleteRegionTypeByID(idInt)
	if err != nil && err.Error() == "This regionType is parent regionType for someone" {
		context.IndentedJSON(400, "This regionType is parent regionType for someone")
	} else if err != nil {
		context.IndentedJSON(404, gin.H{"message": err.Error()})
		log.Println(err)
		return
	}
	context.IndentedJSON(200, gin.H{})
	log.Println("Region deleted")
	return
}
