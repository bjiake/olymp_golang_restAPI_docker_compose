package region

type Region struct {
	ID             int64   `json:"id" bson:"_id"`
	RegionType     int64   `json:"regionTypeAPI"`
	AccountId      int64   `json:"accountId"`
	Name           string  `json:"name"`
	ParentRegionId string  `json:"parentRegionId"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}
