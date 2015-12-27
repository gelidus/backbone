package session

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoRepository is the implementation of the session
// repository for the mongo database
type MongoRepository struct {
	*mgo.Collection
}

// NewMongoRepository will create the new repository instance for
// the user entity
func NewMongoRepository(users *mgo.Collection) *MongoRepository {
	return &MongoRepository{Collection: users}
}

// ReadByUser will return the session based on the links id provided
func (repo *MongoRepository) ReadByLink(link string) (*Model, error) {
	session := &Model{}

	return session, repo.Find(bson.M{"link": bson.ObjectIdHex(link)}).One(session)
}

// Create will insert new session entity into the database
// Warning: this will not check for the entity duplication
func (repo *MongoRepository) Create(session *Model) (error) {
	return repo.Insert(session);
}