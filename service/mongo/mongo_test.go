package mongo

import (
	"testing"
	"github.com/gelidus/backbone/mock"
	"github.com/stretchr/testify/assert"
)

func TestMongoConnection(t *testing.T) {
	db, err := Initialize(&mock.Mongo)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, db)
}