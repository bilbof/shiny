package cache

import (
  "log"
  "net/http"
  "strings"
)

func ServeHTTP(item CacheableResponse, rw http.ResponseWriter, req *http.Request)  {
  log.Print(req.Method, " ", req.URL.Path, "?", req.URL.RawQuery, ": served from cache")

  for k, v := range item.Header {
    rw.Header().Set(k, strings.Join(v, ","))
  }
  rw.Header().Set("Shiny-Cache", "HIT")

  rw.Write(item.Body)
  return
}
