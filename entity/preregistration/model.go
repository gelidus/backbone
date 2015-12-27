package preregistration

import (
	"github.com/asaskevich/govalidator"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/gelidus/backbone/entity/device"
)

// PreRegistration is invitation model that stores only critical data for the
// pre-registered user. This model is populated by the invitation form on the
// landing page
type Model struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Email     string        `json:"email" bson:"email" valid:"required, email"`
	Residence string        `json:"residence" bson:"residence" valid:"required"`
	device.Model

	Invited   bool      `bson:"invited,omitempty" json:"invited"`
	InvitedAt time.Time `bson:"invited_at,omitempty" json:"invited_at"`
}

// Valid checks the provided email and residence
func (r *Model) Valid() (bool, error) {
	return govalidator.ValidateStruct(r)
}

