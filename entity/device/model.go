package device

// Device is entity that represents information about
// users' device
type Model struct {
	Agent   string `json:"agent" bson:"agent"`
	Browser string `json:"browser" bson:"browser"`
	OS      string `json:"os" bson:"os"`
	Device  string `json:"device" bson:"device"`
}
