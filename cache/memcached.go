package cache

import (
  "net/http"
  "github.com/bradfitz/gomemcache/memcache"
)

// Provides a wrapper around Brad Fitz's memcache package

// https://godoc.org/github.com/bradfitz/gomemcache/memcache#Client.Get
func Get(req *http.Request) (c CacheableResponse, err error) {
  key := RequestHash(req)
  item, err := new().Get(key)
  if err != nil {
    return CacheableResponse{}, err
  }
  return BytesToCacheableResponse(item.Value), nil
}

// https://godoc.org/github.com/bradfitz/gomemcache/memcache#Client.Set
func Set(key string, body []byte, expiration int32) error {
  return new().Set(&memcache.Item{
    Key: key,
    Value: body,
    Expiration: expiration,
  })
}

func new() *memcache.Client {
  c := GetConfig()
  return memcache.New(c.Cache.Urls...)
}
