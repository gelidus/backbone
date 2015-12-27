package user

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/asaskevich/govalidator"
)

// User is main entity for any user
type Model struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Name           string        `bson:"name" json:"name"`
	ProfilePicture string        `bson:"profile_picture" json:"profile_picture"`
	Email          string        `bson:"email" json:"email"`
	Phone          string        `bson:"phone,omitempty" json:"phone"`
	About          string        `bson:"about,omitempty" json:"about"`
	Password       string        `bson:"password" json:"-"`
}

// Valid will validate the whole user struct
func (s *Model) Valid() (bool, error) {
	return govalidator.ValidateStruct(s)
}