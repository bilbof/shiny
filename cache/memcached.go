package cache

import (
  "log"
  "github.com/bradfitz/gomemcache/memcache"
)

// Provides a wrapper around Brad Fitz's memcache package

// https://godoc.org/github.com/bradfitz/gomemcache/memcache#Client.Get
func Get(key string) (c CacheableResponse, err error) {
  item, err := new().Get(key)
  if err != nil {
    return CacheableResponse{}, err
  }
  return BytesToCacheableResponse(item.Value), nil
}

// https://godoc.org/github.com/bradfitz/gomemcache/memcache#Client.Set
func Set(key string, body []byte) error {
  log.Print("caching request as: ", key)
  return new().Set(&memcache.Item{
    Key: key,
    Value: body,
    // TODO Expiration, Flags
  })
}

func new() (*memcache.Client) {
  // TODO Use config
  c := GetConfig()
  return memcache.New(c.Cache.Urls...)
}
