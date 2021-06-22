package edgegrid

import (
	"time"

	"github.com/patrickmn/go-cache"
	gocache "github.com/patrickmn/go-cache"
)

// Config struct provides all the necessary fields to
// create authorization header, debug is optional
type (
	profileCache gocache.Cache
)

// Init initializes cache
//

// See: InitCache()
func InitCache() (gocache.Cache, error) {
	cache := cache.New(5*time.Minute, 10*time.Minute)
	return *cache, nil
}
