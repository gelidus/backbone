package preregistration

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	*mgo.Collection
}

func NewMongoRepository(preregistrations *mgo.Collection) *MongoRepository {
	return &MongoRepository{Collection: preregistrations}
}

func (repo *MongoRepository) Create(model *Model) (error) {
	model.ID = bson.NewObjectId()
	_, err := model.Valid()
	if err != nil {
		return err
	}

	return repo.Insert(model)
}

func (repo *MongoRepository) ReadBy(selector map[string]interface{}) (*Model, error) {
	model := &Model{}

	return model, repo.Find(selector).One(model)
}

func (repo *MongoRepository) ReadByEmail(email string) (*Model, error) {
	return repo.ReadBy(map[string]interface{}{
		"email": email,
	})
}

func (repo *MongoRepository) UpdateBy(selector map[string]interface{}, updater map[string]interface{}) (error) {
	return repo.Update(selector, updater)
}

