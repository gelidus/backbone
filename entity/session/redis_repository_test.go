package session

import (
	"testing"
	client "gopkg.in/redis.v3"
	"github.com/gelidus/backbone/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/gelidus/backbone/service/redis"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type RedisSuite struct {
	suite.Suite
	Client *client.Client
	Repo *RedisRepository

	link bson.ObjectId
}

func (suite *RedisSuite) SetupTest() {
	client, err := redis.Initialize(&mock.Redis)
	assert.Equal(suite.T(), nil, err)
	assert.NotEqual(suite.T(), nil, client)

	suite.Repo = NewRedisRepository(client)
	assert.NotEqual(suite.T(), nil, suite.Repo)

	suite.link = bson.NewObjectId()
	suite.Repo.Del(suite.link.Hex()) // remove the link if it exists
}

func (suite *RedisSuite) TestGetSet() {
	// non existing key
	sess, err := suite.Repo.ReadByLink(suite.link.Hex())
	assert.Equal(suite.T(), (*Model)(nil), sess)
	assert.Equal(suite.T(), client.Nil, err)

	var expiration time.Duration = time.Hour * 1
	token := "hello world"

	// setting of key
	err = suite.Repo.Create(&Model{Link: suite.link, Token: token, Expiration: expiration})
	assert.Equal(suite.T(), nil, err)

	// get the existing key
	sess, err = suite.Repo.ReadByLink(suite.link.Hex())
	assert.Equal(suite.T(), nil, err)
	assert.NotEqual(suite.T(), (*Model)(nil), sess)
	assert.Equal(suite.T(), token, sess.Token)
}

func TestRedisSuite(t *testing.T) {
	suite.Run(t, new(RedisSuite))
}