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
	loginIdInt, err := getIntParam(context.Query("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	result, err := requestsToMongoDB.GetAccountByID(loginIdInt)
	if err != nil {
		context.IndentedJSON(404, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	idInt, err := getIntParam(context.Param("accountId"))
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

func Login(c *gin.Context) {
	var inputUser account.AccountLogin
	if err := c.BindJSON(&inputUser); err != nil {
		c.IndentedJSON(401, gin.H{"error": "Invalid request"})
		return
	}

	result, err := requestsToMongoDB.CheckExistAccount(inputUser)
	if err != nil {
		c.IndentedJSON(401, gin.H{"error": "Invalid login data"})
		return
	}

	c.IndentedJSON(200, result.ID)
}

func Registration(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		log.Println("Already registered")
		c.IndentedJSON(403, gin.H{"id": "already_registered"})
		return
	}

	var newAccountRegistration account.AccountRegistration
	if err := c.BindJSON(&newAccountRegistration); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		log.Println("invalid request")
		return
	}
	if !isValidAccountRegister(newAccountRegistration) {
		c.IndentedJSON(400, gin.H{"error": "Invalid data"})
		log.Println("invalid data")
		return
	}
	var newAccount account.Account

	lastID, err := requestsToMongoDB.GetLastIDAccount()
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
		log.Println(err)
		return
	}
	if err != nil {
		if errors.Is(err, errors.New("email already exists")) {
			c.IndentedJSON(409, "Email already exists")
		} else {
			c.IndentedJSON(500, gin.H{"error": err.Error()})
		}
	}

	c.JSON(201, gin.H{"message": "User created"})
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

func getIntParam(value string) (int64, error) {
	loginId := value
	if loginId == "" {
		log.Println("Need to login to account api")
		return -1, errors.New("need to login to account api")
	}

	loginIdInt, err := strconv.ParseInt(loginId, 10, 64)
	if err != nil || loginIdInt <= 0 {
		log.Println(err)
		return -1, errors.New("invalid id value")
	}
	return loginIdInt, nil
}

// Поделать put
func PutAccountByID(context *gin.Context) {
	loginId, err := getIntParam(context.Query("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	id, err := getIntParam(context.Param("accountId"))
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	if loginId != id {
		context.IndentedJSON(403, gin.H{"error": "Try to delete not your account"})
		log.Println(err)
		return
	}

}

func DeleteAccountByID(context *gin.Context) {
	loginIdInt, err := getIntParam(context.Query("id"))
	if err != nil {
		context.IndentedJSON(401, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	idInt, err := getIntParam(context.Param("accountId"))
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
	requestsToMongoDB.DeleteByID(loginIdInt, config.CollectionAccount)
	context.IndentedJSON(200, gin.H{"message": "User deleted"})
}
