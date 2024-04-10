package requestsToMongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"log"
)

func DeleteByID(id int64, collectionName string) {
	collection, err := getCollection(collectionName)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted: %v document in the mods collection\n", deleteResult)
}
