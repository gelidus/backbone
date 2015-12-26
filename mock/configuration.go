package mock
import (
	"gopkg.in/redis.v3"
	"gopkg.in/mgo.v2"
)

var (
	Redis = redis.Options{
		Addr: "192.168.99.100:6379",
		Password: "",
		DB: 0,
	}
	Mongo = mgo.DialInfo{
		Addrs: []string{"192.168.99.100"},
		Database: "test",
	}
)
