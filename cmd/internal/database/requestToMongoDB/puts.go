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

func PutAccount(id int64, updatedAccount account.AccountRegistration) error {
	collection, err := getCollection(config.CollectionAccount)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedAccount}
	var existingAccount account.Account
	err = collection.FindOne(context.TODO(), bson.M{"email": updatedAccount.Email}).Decode(&existingAccount)
	if err == nil && id != existingAccount.ID {
		// A user with the same email already exists
		err = errors.New("emailAlreadyExist")
		log.Println(err)
		return err
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		// An error occurred while querying the database
		return err
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func PutRegion(id int64, newRegion region.NewRegion) error {
	collection, err := getCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": newRegion}
	var existingAccount account.Account
	err = collection.FindOne(context.TODO(), filter).Decode(&existingAccount)
	if err != nil {
		return errors.New("Cannot found this region by regionId")
	}
	err = collection.FindOne(context.TODO(), bson.M{"latitude": newRegion.Latitude, "longitude": newRegion.Longitude}).Decode(&existingAccount)
	if err == nil && id != existingAccount.ID {
		// A user with the same Latitude and Longitude already exists
		return errors.New("Region with this Latitude and Longitude already exists")
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
