package requestsToMongoDB

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/account"
	"golang.org/x/net/context"
	"log"
)

func AddAccount(newAccount account.Account) error {
	collection, err := getCollection(config.CollectionAccount)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if a user with the same email already exists
	var existingAccount account.Account
	err = collection.FindOne(context.TODO(), bson.M{"email": newAccount.Email}).Decode(&existingAccount)
	if err == nil {
		// A user with the same email already exists
		return errors.New("email already exists")
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		// An error occurred while querying the database
		return err
	}

	// Add the new account
	_, err = collection.InsertOne(context.TODO(), newAccount)
	if err != nil {
		return err
	}

	return nil
}
