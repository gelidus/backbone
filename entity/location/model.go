package location

// Location is pair of latitude and longitude values
type Model struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

