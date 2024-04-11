package requestsToMongoDB

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/region"
	"golang.org/x/net/context"
	"log"
)

func DeleteByIDAndCollection(id int64, collectionName string) {
	collection, err := getCollection(collectionName)
	if err != nil {
		log.Println(err)
	}

	filter := bson.D{{"_id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Deleted: %v document in the mods collection\n", deleteResult)
}

func DeleteRegionById(id int64) error {
	collection, err := getCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		return err
	}
	filter := bson.D{{"_id", id}}

	var existingRegion region.Region
	err = collection.FindOne(context.TODO(), filter).Decode(&existingRegion)
	if err != nil {
		log.Println(err)
		return errors.New("This region does not exist")
	}

	filter = bson.D{{"parentRegion", existingRegion.Name}}
	err = collection.FindOne(context.TODO(), filter).Decode(&existingRegion)
	if err == nil && id != existingRegion.ID {
		return errors.New("This region is parent region for someone")
	}

	filter = bson.D{{"_id", id}}
	deleteResult, _ := collection.DeleteOne(context.TODO(), filter)

	log.Printf("Deleted: %v document in the mods collection\n", deleteResult)
	return nil
}
