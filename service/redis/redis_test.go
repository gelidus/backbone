package redis

import (
	"testing"
	"github.com/gelidus/backbone/mock"
	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {
	client, err := Initialize(&mock.Redis)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, client)
}
