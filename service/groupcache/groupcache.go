package groupcache

import "github.com/golang/groupcache"

var (
	pool *groupcache.HTTPPool = nil

	// pool size constants
	Size64MB int64 = 64 << 20
	Size32MB int64 = 32 << 20
	Size16MB int64 = 16 << 20
	Size8MB int64 = 8 << 20
)

// Initialize will startup the groupcache pool
// with the provided self ip address
func Initialize(self string) {
	pool = groupcache.NewHTTPPool(self)
}

func Group(name string) *groupcache.Group {
	group := groupcache.NewGroup(name, Size64MB, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			return nil
		}))

	return group
}
