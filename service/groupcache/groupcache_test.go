package groupcache

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGroupCacheCreate(t *testing.T) {
	self := "10.0.0.1"
	Initialize(self)

	assert.NotEqual(t, nil, pool)

	group := Group("sessions")
	assert.NotEqual(t, nil, group)
}

func TestGroupGet(t *testing.T) {

}
