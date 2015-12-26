package session

import (
	"time"
	"gopkg.in/redis.v3"
	"gopkg.in/mgo.v2/bson"
)

// MongoRepository is the implementation of the session
// repository for the mongo database
type RedisRepository struct {
	*redis.Client
}

// configurable parts
var (
	prefix = "session_"
	expiration = time.Duration(time.Hour * 24 * 3)
)

// NewRedisRepository will create the new repository instance for
// the user entity
func NewRedisRepository(users *redis.Client) *RedisRepository {
	return &RedisRepository{Client: users}
}

// ReadByUser will return the session based on the links id provided
func (repo *RedisRepository) ReadByLink(link string) (*Session, error) {
	session := &Session{}
	token, err := repo.Get(link).Result()
	if err != nil {
		return nil, err
	}

	session.Token = token
	session.Link = bson.ObjectIdHex(link)
	session.Expiration, _ = repo.TTL(link).Result()

	return session, nil
}

// Create will insert new session entity into the database
// Warning: this will not check for the entity duplication
func (repo *RedisRepository) Create(session *Session) (error) {
	return repo.Set(session.Link.Hex(), session.Token, expiration).Err()
}