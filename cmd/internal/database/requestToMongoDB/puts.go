package requestsToMongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/account"
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
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
