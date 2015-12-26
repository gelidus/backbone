package session

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestSessionValidity(t *testing.T) {
	link := bson.NewObjectId()
	session := Session{bson.NewObjectId(), link, "abcd", time.Duration(50)}

	// correct session
	validity, err := session.Valid()
	assert.Equal(t, true, validity)
	assert.Equal(t, nil, err)

	// empty session
	emptySession := Session{}
	validity, err = emptySession.Valid()
	assert.Equal(t, false, validity)
	assert.NotEqual(t, nil, err)

	// empty token
	session.Token = ""
	validity, err = session.Valid()
	assert.Equal(t, false, validity)
	assert.NotEqual(t, nil, err)
}