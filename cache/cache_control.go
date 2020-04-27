package cache

import (
  "math"
  "net/http"
  "time"

  "github.com/pquerna/cachecontrol"
)

// Given a request and response, determine whether an object can be cached,
// and if so, for how long.
func CanCache(req *http.Request, res *http.Response) (bool, int32, error) {
  excuses, expires, err := cachecontrol.CachableResponse(req, res, cachecontrol.Options{})

  if err != nil {
    return false, 0, err
  }

  expiry := int32(math.Max(0, time.Until(expires).Seconds()))

  return len(excuses) == 0 || expiry > 1, expiry, nil
}
