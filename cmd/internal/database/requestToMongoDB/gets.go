package requestsToMongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/account"
	"go_project/cmd/internal/models/region"
	"go_project/cmd/internal/models/regionType"
	"go_project/cmd/internal/models/weather"
	"go_project/cmd/internal/models/weatherForecast"
	"golang.org/x/net/context"
	"log"
)

func GetRegionByID(id int64) (region.Region, error) {
	collection, err := getCollection(config.CollectionRegions)
	if err != nil {
		log.Println(err)
		return region.Region{}, err
	}

	var result region.Region

	filter := bson.D{{"_id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return region.Region{}, err
	}

	log.Printf("Found a single Region document: %+v\n", result)
	return result, nil
}

func GetRegionTypeByID(id int64) (regionType.RegionType, error) {
	collection, err := getCollection(config.CollectionRegionType)
	if err != nil {
		log.Println(err)
		return regionType.RegionType{}, err
	}

	var result regionType.RegionType

	filter := bson.D{{"_id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return regionType.RegionType{}, err
	}

	log.Println("Found a single document: %+v\n", result)
	return result, nil
}

func GetWeatherByID(id int64) (weather.Weather, error) {
	collection, err := getCollection(config.CollectionWeather)
	if err != nil {
		log.Println(err)
		return weather.Weather{}, err
	}

	var result weather.Weather

	filter := bson.D{{"_id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return weather.Weather{}, err
	}

	log.Printf("Found a single document: %+v\n", result)
	return result, nil
}

func GetWeatherForecastByID(id int64) (weatherForecast.WeatherForecast, error) {
	collection, err := getCollection(config.CollectionWeatherForecast)
	if err != nil {
		log.Println(err)
		return weatherForecast.WeatherForecast{}, err
	}

	var result weatherForecast.WeatherForecast

	filter := bson.D{{"regionId", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return weatherForecast.WeatherForecast{}, err
	}

	log.Printf("Found a single document: %+v\n", result)
	return result, nil
}

func GetAccountByID(id int64) (account.AccountInfo, error) {
	collection, err := getCollection(config.CollectionAccount)
	if err != nil {
		log.Println(err)
		return account.AccountInfo{}, err
	}

	var accountExpect account.Account

	filter := bson.D{{"_id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&accountExpect)
	if err != nil {
		log.Println(err)
		return account.AccountInfo{}, err
	}
	var result account.AccountInfo
	result.FirstName = accountExpect.FirstName
	result.LastName = accountExpect.LastName
	result.ID = accountExpect.ID
	result.Email = accountExpect.Email

	log.Printf("GET single account documenct: %+v\n", result)
	return result, nil
}

func GetLastIDByCollection(collectionName string) (int64, error) {
	collection, err := getCollection(collectionName)
	if err != nil {
		return 0, err
	}

	var result bson.M

	filter := bson.D{{}}
	options := options.FindOne().SetSort(bson.D{{"_id", -1}})
	err = collection.FindOne(context.TODO(), filter, options).Decode(&result)
	if err != nil {
		return 0, err
	}

	log.Printf("Found last document account ID: %d\n", result["_id"])
	return result["_id"].(int64), nil
}

func CheckExistAccount(current account.AccountLogin) (account.Account, error) {
	collection, err := getCollection(config.CollectionAccount)
	if err != nil {
		return account.Account{}, err
	}

	filter := bson.D{{"email", current.Email}, {"password", current.Password}}
	var result account.Account
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return account.Account{}, err
	}

	log.Printf("Found exist account by: %+v\n", result)
	return result, nil
}
