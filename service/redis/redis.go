package redis

import "gopkg.in/redis.v3"

func Initialize(opts *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opts)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
