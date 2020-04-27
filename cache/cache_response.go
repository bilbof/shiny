package cache

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strconv"
)

func CacheResponse(req *http.Request, res *http.Response) error {
  cacheable, expires, err := CanCache(req, res)

  if err != nil { return err }
  if !cacheable { return nil }

  body, err := ioutil.ReadAll(res.Body)
  res.Body = ioutil.NopCloser(bytes.NewBuffer(body))

  if err != nil { return err }

  i := CacheableResponse{
    Key: RequestHash(req),
    Body: body,
    Header: res.Header,
    Expiration: expires,
  }

  if err := Set(i.Key, i.Bytes(), i.Expiration); err != nil {
    fmt.Errorf("Error: Failed to set cache", err)
    return err
  }

  return nil
}
