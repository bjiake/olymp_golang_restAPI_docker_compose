package requestsToMongoDB

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_project/cmd/internal/config"
	"go_project/cmd/internal/models/account"
	"go_project/cmd/internal/models/region"
	"go_project/cmd/internal/models/regionType"
	"go_project/cmd/internal/models/weather"
	"go_project/cmd/internal/models/weatherForecast"
	"golang.org/x/net/context"
	"log"
	"regexp"
	"time"
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

func GetWeatherByID(id int64) ([]weather.Weather, error) {
	collectionWeatherForecast, err := getCollection(config.CollectionWeatherForecast)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var weatherForecasts []weatherForecast.WeatherForecast
	filter := bson.D{{"regionId", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collectionWeatherForecast.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = cursor.All(ctx, &weatherForecasts); err != nil {
		return nil, err
	}

	collectionWeather, err := getCollection(config.CollectionWeather)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	weatherForecastIds := make([]int64, 0)
	for _, value := range weatherForecasts {
		weatherForecastIds = append(weatherForecastIds, value.ID)
	}
	filterWeather := bson.M{"weatherForecast": bson.M{"$all": weatherForecastIds}}

	var result []weather.Weather
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err = collectionWeather.Find(ctx, filterWeather)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
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

func GetSearchAccount(currentAccount account.Search) ([]account.AccountInfo, error) {
	collection, err := getCollection(config.CollectionAccount)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	emailQuery := regexp.MustCompile(currentAccount.Email)
	firstNameQuery := regexp.MustCompile(currentAccount.FirstName)
	lastNameQuery := regexp.MustCompile(currentAccount.LastName)
	filter := bson.M{}
	if emailQuery.String() != "" {
		filter["$and"] = []bson.M{{"email": bson.M{"$regex": emailQuery.FindString(currentAccount.Email)}}}
	}
	if firstNameQuery.String() != "" {
		if len(filter) == 0 {
			filter["$and"] = []bson.M{}
		}
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"firstName": bson.M{"$regex": firstNameQuery.FindString(currentAccount.FirstName)}})
	}
	if lastNameQuery.String() != "" {
		if len(filter) == 0 {
			filter["$and"] = []bson.M{}
		}
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"lastName": bson.M{"$regex": lastNameQuery.FindString(currentAccount.LastName)}})
	}

	if len(filter) == 0 {
		filter = bson.M{}
	} else if len(filter["$and"].([]bson.M)) == 1 {
		filter = filter["$and"].([]bson.M)[0]
	}

	findOptions := options.Find()
	findOptions.SetSkip(currentAccount.Form * currentAccount.Size)
	findOptions.SetLimit(currentAccount.Size)

	var results []account.AccountInfo
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetSearchWeather(currentWeather weather.SearchWeather) ([]weather.Weather, error) {
	collectionWeather, err := getCollection(config.CollectionWeather)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	weatherConditionQuery := regexp.MustCompile(currentWeather.WeatherCondition)
	utcStartDateTime := currentWeather.StartDateTime.Time()
	startDateTimeDB := primitive.NewDateTimeFromTime(utcStartDateTime)
	startDateTimeQuery := regexp.MustCompile(startDateTimeDB.Time().Format("2006-01-02T15:04:05Z07:00"))

	utcEndDateTime := currentWeather.EndDateTime.Time()
	endDateTimeDB := primitive.NewDateTimeFromTime(utcEndDateTime)
	endDateTimeQuery := regexp.MustCompile(endDateTimeDB.Time().Format("2006-01-02T15:04:05Z07:00"))
	var regionNameString string
	if currentWeather.RegionId > 0 {
		expectedRegion, err := GetRegionByID(currentWeather.RegionId)
		if err != nil {
			log.Println("cannot find region")
			return nil, errors.New("cannot find region")
		}
		regionNameString = expectedRegion.Name
	}
	var regionNameQuery *regexp.Regexp
	if regionNameString != "" {
		regionNameQuery = regexp.MustCompile(regionNameString)
	}

	filter := bson.M{}
	if weatherConditionQuery.String() != "" {
		if len(filter) == 0 {
			filter["$and"] = []bson.M{}
		}
		filter["$and"] = []bson.M{{"weatherCondition": bson.M{"$regex": weatherConditionQuery.FindString(currentWeather.WeatherCondition)}}}
	}
	if startDateTimeQuery.String() != "" && startDateTimeQuery.String() != "�" && startDateTimeQuery.String() != "0001-01-01T07:00:00+07:00" {
		if len(filter) == 0 {
			filter["$and"] = []bson.M{}
		}
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"measurementDateTime": bson.M{"$gte": startDateTimeDB}})
	}
	if endDateTimeQuery.String() != "" && endDateTimeQuery.String() != "�" && endDateTimeQuery.String() != "0001-01-01T07:00:00+07:00" {
		if len(filter) == 0 {
			filter["and"] = []bson.M{}
		}
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"measurementDateTime": bson.M{"$lte": endDateTimeDB}})
	}
	if regionNameQuery != nil && regionNameQuery.String() != "" {
		if len(filter) == 0 {
			filter["$and"] = []bson.M{}
		}
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{"regionName": bson.M{"$regex": regionNameQuery.FindString(regionNameString)}})
	}

	if len(filter) == 0 {
		filter = bson.M{}
	} else if len(filter["$and"].([]bson.M)) == 1 {
		filter = filter["$and"].([]bson.M)[0]
	}

	findOptions := options.Find()
	findOptions.SetSkip(currentWeather.Form * currentWeather.Size)
	findOptions.SetLimit(currentWeather.Size)

	var results []weather.Weather
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collectionWeather.Find(ctx, filter, findOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
