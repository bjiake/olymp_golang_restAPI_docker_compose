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