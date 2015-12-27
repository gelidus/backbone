package preregistration

import (
	"github.com/stretchr/testify/suite"
	"github.com/gelidus/backbone/service/mongo"
	"github.com/gelidus/backbone/mock"
	"github.com/stretchr/testify/assert"
"testing"
)

type MongoRepositorySuite struct {
	suite.Suite

	repo *MongoRepository
}

func (suite *MongoRepositorySuite) SetupTest() {
	client, err := mongo.Initialize(&mock.Mongo)
	assert.Equal(suite.T(), nil, err)

	suite.repo = NewMongoRepository(client.C("test_preregistration"))
	assert.NotEqual(suite.T(), nil, suite.repo)

	// drop the test collection if it exists
	assert.Equal(suite.T(), nil, suite.repo.DropCollection())
}

func (suite *MongoRepositorySuite) TestCreate() {
	model := &Model{
		Email: "gelidus@hotmail.sk",
		Residence: "Brno",
	}

	err := suite.repo.Create(model)
	assert.Equal(suite.T(), nil, err)
}

func (suite *MongoRepositorySuite) TestReadBy() {
	model, err := suite.repo.ReadBy(map[string]interface{}{
		"residence": "Brno",
	})

	assert.Equal(suite.T(), nil, err)
	assert.NotEqual(suite.T(), nil, model)
	assert.Equal(suite.T(), "gelidus@hotmail.sk", model.Email)
	assert.Equal(suite.T(), "Brno", model.Residence)
}

func (suite *MongoRepositorySuite) TestReadByEmail() {
	model, err := suite.repo.ReadByEmail("gelidus@hotmail.sk")

	assert.Equal(suite.T(), nil, err)
	assert.NotEqual(suite.T(), nil, model)
	assert.Equal(suite.T(), "gelidus@hotmail.sk", model.Email)
	assert.Equal(suite.T(), "Brno", model.Residence)
}

func (suite *MongoRepositorySuite) TestUpdateBy() {

}

func TestMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(MongoRepositorySuite))
}