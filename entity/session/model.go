package session

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/asaskevich/govalidator"
)

// Session entity handles the session storing of the linking
// objects. Any object id can be provided as the linked object
type Model struct {
	ID         bson.ObjectId `bson:"_id" json:"id" valid:"required"`
	Link       bson.ObjectId `bson:"link" json:"link" valid:"required"`
	Token      string `bson:"token" json:"token" valid:"required"`
	Expiration time.Duration `bson:"expiration" json:"expiration"`
}

// Valid will validate the whole Session struct
func (s *Model) Valid() (bool, error) {
	return govalidator.ValidateStruct(s)
}