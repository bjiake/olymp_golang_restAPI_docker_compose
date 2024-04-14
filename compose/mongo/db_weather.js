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
