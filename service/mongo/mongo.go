package mongo

import "gopkg.in/mgo.v2"

// default configuration
var (
	defaultConfig = mgo.DialInfo{
		Addrs: []string{"localhost"},
		Database: "test",
	}

	session *mgo.Session = nil
)

// Initialize will try to connect to the database
func Initialize(config *mgo.DialInfo) (*mgo.Database, error) {
	var err error

	// if no configuration was provided, reset it to default
	if config == nil {
		config = &defaultConfig
	}

	session, err = mgo.DialWithInfo(config)
	if err != nil {
		return nil, err
	}

	return Database(config.Database), nil
}

// Database will return wanted database from the
// currently connected session
func Database(name string) *mgo.Database {
	if session == nil {
		return nil
	}

	return session.DB(name)
}