package region

type Region struct {
	ID             int64   `json:"id" bson:"_id" required:"true"`
	RegionType     int64   `json:"regionType"`
	AccountId      int64   `json:"accountId"`
	Name           string  `json:"name"`
	ParentRegionId string  `json:"parentRegion"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}

type NewRegion struct {
	RegionType     int64   `json:"regionType" bson:"regionType"`
	Name           string  `json:"name" bson:"name" required:"true"`
	ParentRegionId string  `json:"parentRegion" bson:"parentRegion"`
	Latitude       float64 `json:"latitude" bson:"latitude"`
	Longitude      float64 `json:"longitude" bson:"longitude"`
}
