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