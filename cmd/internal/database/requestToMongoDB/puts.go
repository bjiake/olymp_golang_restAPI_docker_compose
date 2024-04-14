package requestsToMongoDB

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/account"
	"go_project/cmd/internal/models/region"
	"go_project/cmd/internal/models/regionType"
	"go_project/cmd/internal/models/weather"
	"go_project/cmd/internal/models/weatherForecast"
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
	var existingAccount region.Region
	err = collection.FindOne(context.TODO(), filter).Decode(&existingAccount)
	if err != nil {
		return errors.New("Cannot found this region by regionId")
	}
	err = collection.FindOne(context.TODO(), bson.M{"latitude": newRegion.Latitude, "longitude": newRegion.Longitude}).Decode(&existingAccount)
	if err == nil && id != existingAccount.ID {
		// A user with the same Latitude and Longitude already exists
		return errors.New("Region with this Latitude and Longitude already exists")
	}

	filter = bson.M{"_id": id}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func PutRegionType(id int64, newRegion regionType.NewRegionType) error {
	collection, err := getCollection(config.CollectionRegionType)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": newRegion}
	var existingRegionType regionType.RegionType
	err = collection.FindOne(context.TODO(), filter).Decode(&existingRegionType)
	if err != nil {
		return errors.New("Cannot found this regionType by typeId")
	}
	err = collection.FindOne(context.TODO(), bson.M{"type": newRegion.Type}).Decode(&existingRegionType)
	if err == nil && id != existingRegionType.ID {
		// A user with the same Latitude and Longitude already exists
		return errors.New("Region with this Type already exists")
	}

	filter = bson.M{"_id": id}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func PutWeather(newWeather weather.Weather) error {
	collection, err := getCollection(config.CollectionWeather)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": newWeather.ID}
	update := bson.M{"$set": newWeather}

	var existingWeather weather.Weather
	err = collection.FindOne(context.TODO(), filter).Decode(&existingWeather)
	if err != nil {
		log.Println(err)
		return err
	}

	collectionWeatherForecast, err := getCollection(config.CollectionWeatherForecast)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, value := range newWeather.WeatherForecast {
		var existingWeatherForecast weatherForecast.WeatherForecast
		err = collectionWeatherForecast.FindOne(context.TODO(), bson.M{"_id": value}).Decode(&existingWeatherForecast)
		if err != nil {
			return errors.New("cannot found forecast id on add weather")
		} else if !errors.Is(err, mongo.ErrNoDocuments) {
			// An error occurred while querying the database
			return err
		}
	}

	filter = bson.M{"_id": newWeather.ID}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func PutWeatherForecast(id int64, updatedWeatherForecast weatherForecast.NewPutWeatherForecast) error {
	collection, err := getCollection(config.CollectionWeatherForecast)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedWeatherForecast}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
