package requestsToMongoDB

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/region"
	"golang.org/x/net/context"
	"log"
)

func DeleteByIDAndCollection(id int64, collectionName string) error {
	collection, err := getCollection(collectionName)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.D{{"_id", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if deleteResult.DeletedCount == 0 {
		log.Println(err)
		return errors.New("No document found")
	}
	log.Printf("Deleted: %v document in the mods collection\n", deleteResult)
	return nil
}

func DeleteRegionTypeByID(id int64) error {
	collectionRegion, err := getCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		return err
	}
	var expectedRegion region.Region
	filter := bson.D{{"regionType", id}}
	err = collectionRegion.FindOne(context.TODO(), filter).Decode(&expectedRegion)
	if err == nil {
		return errors.New("This regionType is parent regionType for someone")
	}
	err = DeleteByIDAndCollection(id, config.CollectionRegionType)
	if err != nil {
		return err
	}
	return nil
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

func DeleteWeatherById(weatherId int64, regionId int64) error {
	var existingRegion region.Region
	existingRegion, err = GetRegionByID(regionId)
	if err != nil {
		log.Println(err)
		return err
	}
	filter := bson.D{{"regionName", existingRegion.Name}}
	collectionWeather, err := getCollection(config.CollectionWeather)
	if err != nil {
		log.Println(err)
		return err
	}

	err = collectionWeather.FindOne(context.TODO(), filter).Decode(&existingRegion)
	if err != nil {
		log.Println(err)
		return errors.New("This region does not exist")
	}

	filter = bson.D{{"_id", weatherId}}
	deleteResult, err := collectionWeather.DeleteOne(context.TODO(), filter)
	if deleteResult.DeletedCount == 0 {
		log.Println(err)
		return errors.New("This weather does not exist")
	}
	log.Printf("Deleted: %v document in the mods collection\n", deleteResult)
	return nil
}
