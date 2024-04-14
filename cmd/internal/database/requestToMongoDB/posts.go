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
	} else if errors.Is(err, mongo.ErrNoDocuments) {
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
	var existingAccount region.Region
	err = collection.FindOne(context.TODO(), bson.M{"latitude": newRegion.Latitude, "longitude": newRegion.Longitude}).Decode(&existingAccount)
	if err == nil {
		// A user with the same Latitude and Longitude already exists
		return errors.New("Region with this Latitude and Longitude already exists")
	} else if errors.Is(err, mongo.ErrNoDocuments) {
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

func AddRegionType(newRegion regionType.RegionType) error {
	collection, err := getCollection(config.CollectionRegionType)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if a Latitude and Longitude already exist
	var existingAccount region.Region
	err = collection.FindOne(context.TODO(), bson.M{"type": newRegion.Type}).Decode(&existingAccount)
	if err == nil {
		// A user with the same Latitude and Longitude already exists
		return errors.New("Region with this Type already exists")
	} else if errors.Is(err, mongo.ErrNoDocuments) {
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

func AddWeather(newWeather weather.Weather) error {
	collectionWeather, err := getCollection(config.CollectionWeather)
	if err != nil {
		log.Println(err)
		return err
	}
	collectionRegion, err := getCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		return err
	}
	collectionWeatherForecast, err := getCollection(config.CollectionWeatherForecast)
	if err != nil {
		log.Println(err)
		return err
	}

	isFoundRegion := false
	for _, value := range newWeather.WeatherForecast {
		var existingWeatherForecast weatherForecast.WeatherForecast
		err = collectionWeatherForecast.FindOne(context.TODO(), bson.M{"_id": value}).Decode(&existingWeatherForecast)

		var existingRegion region.Region
		err = collectionRegion.FindOne(context.TODO(), bson.M{"name": newWeather.RegionName}).Decode(&existingRegion)
		if existingRegion.ID == value {
			isFoundRegion = true
		}
	}
	if !isFoundRegion {
		return errors.New("weather with this WeatherForecastID with region does not exist")
	}

	// Add the new account
	_, err = collectionWeather.InsertOne(context.TODO(), newWeather)
	if err != nil {
		return err
	}

	return nil
}

func AddWeatherForecast(newWeatherForecast weatherForecast.WeatherForecast) error {
	collectionWeatherForecast, err := getCollection(config.CollectionWeatherForecast)
	if err != nil {
		log.Println(err)
		return err
	}
	collectionRegion, err := getCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check if a Latitude and Longitude already exist
	var existingAccount region.Region
	err = collectionRegion.FindOne(context.TODO(), bson.M{"_id": newWeatherForecast.RegionID}).Decode(&existingAccount)
	if err != nil {
		// A user with the same Latitude and Longitude already exists
		return errors.New("region with this id not found")
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		// An error occurred while querying the database
		return err
	}

	// Add the new account
	_, err = collectionWeatherForecast.InsertOne(context.TODO(), newWeatherForecast)
	if err != nil {
		return err
	}

	return nil
}
