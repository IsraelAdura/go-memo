package gomemo

import (
	"time"
)

type memo struct {
	Expires time.Time
	Result  interface{}
}

// Cache map to hold all the cached values
var Cache = make(map[string]*memo)

// Memoize accepts a cacheKey,a function to execute and an expiry period
// check if cache key already exists in the map i.e call is already memoized
// if it has expired or does not exist the function is executed
// and its result and set expiration time stored in the cache
func Memoize(cacheKey string, caller func() interface{}, expires uint) interface{} {
	if expires == 0 {
		return caller()
	}

	memoized := Cache[cacheKey]
	if memoized == nil || memoized.Expires.Before(time.Now()) {
		result := caller()
		if result != nil {
			Cache[cacheKey] = &memo{
				Expires: time.Now().Add(time.Duration(expires) * time.Second),
				Result:  result,
			}
		}
		return result
	}
	return memoized.Result
}

// UnMemoize remove cacheKey from cache
func UnMemoize(cacheKey string) {
	delete(Cache, cacheKey)
}

// ClearCache delete all entries from the cache
func ClearCache() {
	for cacheKey := range Cache {
		delete(Cache, cacheKey)
	}
}

// Get value
func Get(cacheKey string) interface{} {
	memoized := Cache[cacheKey]
	if memoized != nil && memoized.Expires.After(time.Now()) {
		return memoized.Result
	}
	return nil
}
