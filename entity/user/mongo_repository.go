package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	*mgo.Collection
}

// NewMongoRepository will create the new repository instance for
// the user entity
func NewMongoRepository(users *mgo.Collection) *MongoRepository {
	return &MongoRepository{Collection: users}
}

// ReadByID will return the user based on the provided ID
func (repo *MongoRepository) ReadByID(id string) (*Model, error) {
	user := &Model{}

	return user, repo.FindId(bson.ObjectIdHex(id)).One(user)
}

// ReadByEmail will return the user based on the email provided
func (repo *MongoRepository) ReadByEmail(email string) (*Model, error) {
	user := &Model{}

	return user, repo.Find(bson.M{"email": email}).One(user)
}

// Create will insert new user entity into the database
// Warning: this will not check for the entity duplication
func (repo *MongoRepository) Create(user *Model) (error) {
	return repo.Insert(user);
}