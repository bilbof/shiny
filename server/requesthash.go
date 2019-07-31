package server

import (
  "crypto/sha256"
  "encoding/base64"
  "net/http"
  "log"
)

// Generates a hash for a request.

func RequestHash(req *http.Request) string {
  k := key(req)
  h := sha256.New()
  h.Write([]byte(k))
  return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func key(req *http.Request) string {
  log.Print(req.Method + req.URL.Path + req.URL.RawQuery)
  return req.Method + req.URL.Path + req.URL.RawQuery
}
