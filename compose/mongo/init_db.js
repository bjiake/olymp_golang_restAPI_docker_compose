var account = db.getCollection("account")
var account1 = {
    "_id": 1,
    firstName: "John",
    lastName: "Doe",
    email: "john.doe@example.com",
    password: "secret123"
}
var account2 = {
    "_id": 2,
    firstName: "Jane",
    lastName: "Doe",
    email: "jane.doe@example.com",
    password: "password456"
}
var account3 ={
    "_id": 3,
    firstName: "Bob",
    lastName: "Smith",
    email: "bob.smith@example.com",
    password: "qwerty789"
}
var account4 = {
    "_id": 4,
    firstName: "Alice",
    lastName: "Johnson",
    email: "alice.johnson@example.com",
    password: "letmein123"
}
var account5 = {
    "_id": 5,
    firstName: "Charlie",
    lastName: "Brown",
    email: "charlie.brown@example.com",
    password: "12345678"
}

account.insertMany([account1,account2,account3,account4,account5])

var region = db.getCollection("region")

var region1 = {
    "_id": Long(1),
    regionType: Long(1), // Change this to long
    accountId: Long(123),
    name: "North America",
    parentRegion: "",
    latitude: 37.0902,
    longitude: -95.7129
}
var region2 = {
    "_id": Long(2),
    regionType:  Long(2),
    accountId: Long(123),
    name: "United States",
    parentRegion: "North America",
    latitude: 37.0902,
    longitude: -95.7129
}
var region3 = {
    "_id": Long(3),
    regionType:  Long(3),
    accountId: Long(123),
    name: "New York",
    parentRegion: "United States",
    latitude: 40.7128,
    longitude: -74.0060
}
var region4 = {
    "_id": Long(4),
    regionType:  Long(2),
    accountId: Long(123),
    name: "United States",
    parentRegion: "North America",
    latitude: 37.0902,
    longitude: -95.7129
}
var region5 = {
    "_id": Long(5),
    regionType:  Long(3),
    accountId: Long(123),
    name: "Florida",
    parentRegion: "United States",
    latitude: 27.6648,
    longitude: -81.5158
}
region.insertMany([region1,region2,region3,region4,region5])

var regionTypes = db.getCollection("regionType")
var region1 = {
    "_id": Long(1),
    type: "Country"
}
var region2 = {
    "_id": Long(2),
    type: "State"
}
var region3 = {
    "_id": Long(3),
    type: "City"
}
regionTypes.insertMany([region1, region2, region3])

var weather = db.getCollection("weather")
var weather1 = {
    "_id": Long(1),
    regionName: "New York",
    temperature: Double(25.5),
    humidity: Double(50.0),
    windSpeed: Double(5.0),
    weatherCondition: "CLEAR",
    precipitationAmount: Double(0.0),
    measurementDateTime: ISODate("2022-03-01T12:00:00.000Z"),
    weatherForecast: [
        Long(1), Long(2), Long(3)
    ]
}
var weather2 = {
    "_id": Long(2),
    regionName: "Los Angeles",
    temperature: Double(20.0),
    humidity: Double(30.0),
    windSpeed: Double(3.0),
    weatherCondition: "CLEAR",
    precipitationAmount: Double(0.0),
    measurementDateTime: ISODate("2022-03-01T12:00:00.000Z"),
    weatherForecast: [
        Long(1), Long(2), Long(3)
    ]
}
var weather3 = {
    "_id": Long(3),
    regionName: "Chicago",
    temperature: Double(10.0),
    humidity: Double(70.0),
    windSpeed: Double(7.0),
    weatherCondition: "STORM",
    precipitationAmount: Double(10.0),
    measurementDateTime: ISODate("2022-03-01T12:00:00.000Z"),
    weatherForecast: [
        Long(1), Long(2), Long(3)
    ]
}
var weather4 = {
    "_id": Long(4),
    regionName: "Miami",
    temperature: Double(30.0),
    humidity: Double(80.0),
    windSpeed: Double(4.0),
    weatherCondition: "RAIN",
    precipitationAmount: Double(5.0),
    measurementDateTime: ISODate("2022-03-01T12:00:00.000Z"),
    weatherForecast: [
        Long(1), Long(2), Long(3)
    ]
}
var weather5 ={
    "_id": Long(5),
    regionName: "Seattle",
    temperature: Double(15.0),
    humidity: Double(60.0),
    windSpeed: Double(6.0),
    weatherCondition: "CLOUDY",
    precipitationAmount: Double(2.0),
    measurementDateTime: ISODate("2022-03-01T12:00:00.000Z"),
    weatherForecast: [
        Long(1), Long(2), Long(3)
    ]
}
weather.insertMany([weather1, weather2, weather3,weather4, weather5])

var weatherForecast = db.getCollection("weatherForecast")

var weatherForecast1 = {
    "_id": Long(1),
    dateTime: ISODate("2022-03-02T12:00:00.000Z"),
    temperature: Double(25.5),
    weatherCondition: "CLEAR",
    regionId: Long(1)
}

var weatherForecast2 = {
    "_id": Long(2),
    dateTime: ISODate("2022-03-03T12:00:00.000Z"),
    temperature: Double(20.0),
    weatherCondition: "CLOUDY",
    regionId: Long(2)
}
var weatherForecast3 = {
    "_id": Long(3),
    dateTime: ISODate("2022-03-04T12:00:00.000Z"),
    temperature: Double(10.0),
    weatherCondition: "RAIN",
    regionId: Long(1)
}

var weatherForecast4 = {
    "_id": Long(4),
    dateTime: ISODate("2022-03-05T12:00:00.000Z"),
    temperature: Double(5.0),
    weatherCondition: "SNOW",
    regionId: Long(1)
}
var weatherForecast5 = {
    "_id": Long(5),
    dateTime: ISODate("2022-03-06T12:00:00.000Z"),
    temperature: Double(15.0),
    weatherCondition: "FOG",
    regionId: Long(1)
}

weatherForecast.insertMany([weatherForecast1, weatherForecast2, weatherForecast3, weatherForecast4, weatherForecast5])