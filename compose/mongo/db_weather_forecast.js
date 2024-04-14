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