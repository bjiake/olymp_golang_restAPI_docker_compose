package accountAPI

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_project/cmd/internal/config"
	requestsToMongoDB "go_project/cmd/internal/database/requestToMongoDB"
	"go_project/cmd/internal/models/account"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type AccountRequest struct {
	CurrentPage int               `json:"current_Page"`
	Data        []account.Account `json:"data"`
	PrevPage    string            `json:"prev_page"`
	NextPage    string            `json:"next_page"`
	TotalPage   int               `json:"total_Page"`
}

func NewRequestAPI(currentPage int, data []account.Account, prevPage string, nextPage string, totalPage int) *AccountRequest {
	return &AccountRequest{CurrentPage: currentPage, Data: data, PrevPage: prevPage, NextPage: nextPage, TotalPage: totalPage}
}

func GetAccountByID(context *gin.Context) {
	loginIdInt, err := config.GetIntParam(context.Cookie("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	result, err := requestsToMongoDB.GetAccountByID(loginIdInt)
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "this account does not exist"})
		log.Println("this account does not exist")
		return
	}

	idInt, err := config.GetIntParam(context.Param("accountId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	result, err = requestsToMongoDB.GetAccountByID(idInt)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	context.IndentedJSON(http.StatusOK, result)
	return
}

// Login Handler
func Login(c *gin.Context) {
	//request account
	var inputUser account.AccountLogin
	if err := c.BindJSON(&inputUser); err != nil {
		c.IndentedJSON(401, gin.H{"error": "Invalid request"})
		return
	}

	//Check loginId
	result, err := requestsToMongoDB.CheckExistAccount(inputUser)
	if err != nil {
		c.IndentedJSON(401, gin.H{"error": "Invalid login data"})
		return
	}

	c.IndentedJSON(200, gin.H{"id": result.ID})
}

// Registration Handler
func Registration(c *gin.Context) {
	//Check login
	_, err := config.GetIntParam(c.Cookie("id"))
	if err == nil {
		log.Println("Already registered")
		c.IndentedJSON(403, gin.H{"id": "already_registered"})
		return
	}

	//handle request
	var newAccountRegistration account.AccountRegistration
	if err := c.BindJSON(&newAccountRegistration); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		log.Println("invalid request")
		return
	}

	//check valid data
	if !isValidAccountRegister(newAccountRegistration) {
		c.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}

	var newAccount account.Account
	lastID, err := requestsToMongoDB.GetLastIDByCollection(config.CollectionAccount)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(500, gin.H{"error": "internal server error cannot get last id"})
		return
	}

	newAccount.ID = lastID + 1
	newAccount.Email = newAccountRegistration.Email
	newAccount.FirstName = newAccountRegistration.FirstName
	newAccount.LastName = newAccountRegistration.LastName
	newAccount.Password = newAccountRegistration.Password

	err = requestsToMongoDB.AddAccount(newAccount)
	if err != nil {
		if !errors.Is(err, errors.New("email already exists")) {
			c.IndentedJSON(409, "Email already exists")
			log.Println(err)
			return
		} else {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
	}

	infoAccount, _ := requestsToMongoDB.GetAccountByID(newAccount.ID)

	c.JSON(201, infoAccount)
	log.Printf("User created: %+v\n", newAccount)
}

func isValidAccountRegister(newAccountRegistration account.AccountRegistration) bool {
	if newAccountRegistration.FirstName == "" || strings.TrimSpace(newAccountRegistration.FirstName) == "" {
		return false
	}

	if newAccountRegistration.LastName == "" || strings.TrimSpace(newAccountRegistration.LastName) == "" {
		return false
	}

	if newAccountRegistration.Email == "" || strings.TrimSpace(newAccountRegistration.Email) == "" || !isValidEmail(newAccountRegistration.Email) {
		return false
	}

	if newAccountRegistration.Password == "" || strings.TrimSpace(newAccountRegistration.Password) == "" {
		return false
	}

	return true
}

func isValidEmail(email string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return match
}

func PutAccountByID(context *gin.Context) {
	loginId, err := config.GetIntParam(context.Cookie("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	_, err = requestsToMongoDB.GetAccountByID(loginId)
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": "this account does not exist"})
		log.Println("this account does not exist")
		return
	}
	id, err := config.GetIntParam(context.Param("accountId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	if loginId != id {
		context.IndentedJSON(403, gin.H{"error": "Try to update not your account"})
		log.Println(err)
		return
	}
	var updatedDoc account.AccountRegistration
	err = context.BindJSON(&updatedDoc)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": "Invalid request"})
		log.Println(err)
		return
	}
	if !isValidAccountRegister(updatedDoc) {
		context.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}

	err = requestsToMongoDB.PutAccount(id, updatedDoc)
	if err != nil {
		if err.Error() == "emailAlreadyExist" {
			context.IndentedJSON(409, gin.H{"error": err.Error()})
		} else {
			context.IndentedJSON(404, gin.H{"error": err.Error()})
		}
	}

	context.IndentedJSON(200, updatedDoc)
}

func DeleteAccountByID(context *gin.Context) {
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
	idInt, err := config.GetIntParam(context.Param("accountId"), nil)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	if loginIdInt != idInt {
		context.IndentedJSON(403, gin.H{"error": "Try to delete not your account"})
		log.Println(err)
		return
	}
	err = requestsToMongoDB.DeleteByIDAndCollection(loginIdInt, config.CollectionAccount)
	if err != nil {
		context.IndentedJSON(404, gin.H{"message": err.Error()})
		log.Println(err)
		return
	}
	context.IndentedJSON(200, gin.H{})
	log.Println("user deleted")
	return
}

func SearchAccount(c *gin.Context) {
	var accountSearch account.Search
	accountSearch.FirstName = c.Query("firstName")
	accountSearch.LastName = c.Query("lastName")
	accountSearch.Email = c.Query("email")
	formString := c.Query("form")
	if formString == "" {
		log.Println("invalid form")
		c.IndentedJSON(400, gin.H{"error": "invalid form"})
		return
	}

	formInt, err := strconv.ParseInt(formString, 10, 64)
	if err != nil || formInt < 0 {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": "invalid form"})
		return
	}
	accountSearch.Form = formInt
	sizeInt, err := config.GetIntParam(c.Query("size"), nil)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(400, gin.H{"error": "size must be greater than zero"})
		return
	}
	accountSearch.Size = sizeInt

	accountsInfo, err := requestsToMongoDB.GetSearchAccount(accountSearch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accountsInfo)
}
