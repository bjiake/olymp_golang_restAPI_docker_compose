package regionType

type RegionType struct {
	ID   int64  `json:"id" bson:"_id"`
	Type string `json:"type" bson:"type"`
}

type NewRegionType struct {
	Type string `json:"type" bson:"type"`
}
