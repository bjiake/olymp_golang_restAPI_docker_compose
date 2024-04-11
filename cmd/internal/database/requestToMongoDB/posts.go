package requestsToMongoDB

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/account"
	"go_project/cmd/internal/models/region"
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

func AddRegion(newRegion region.Region) error {
	collection, err := getCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if a Latitude and Longitude already exist
	var existingAccount account.Account
	err = collection.FindOne(context.TODO(), bson.M{"latitude": newRegion.Latitude, "longitude": newRegion.Longitude}).Decode(&existingAccount)
	if err == nil {
		// A user with the same Latitude and Longitude already exists
		return errors.New("Region with this Latitude and Longitude already exists")
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		// An error occurred while querying the database
		return err
	}

	// Add the new account
	_, err = collection.InsertOne(context.TODO(), newRegion)
	if err != nil {
		return err
	}

	return nil
}
