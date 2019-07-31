package server

import (
  "bytes"
  "flag"
  "io/ioutil"
  "log"
  "fmt"
  "net/http"
  "net/url"
  "github.com/bilbof/shiny/proxy"
  "github.com/bilbof/shiny/cache"
)

// The server package runs the HTTP server that listens to all requests.
// At the moment it just forwards requests to the proxy server.

type Server struct {
  Port  string
  BackendUrl string
}

func (s Server) ListenAndServe() (error) {
  log.Print("Listening at ", s.Port)
  http.Handle("/", http.HandlerFunc(s.requestHandler))
  var addr = flag.String("addr", s.Port, "http service address")
  flag.Parse()
  return http.ListenAndServe(*addr, nil)
}

// requestHandler takes a request and passes it either to the cache
// or to the proxy server
func (s Server) requestHandler(rw http.ResponseWriter, req *http.Request)  {
  log.Print("requestHandler")
  key := RequestHash(req)
  log.Print("looking up in cache")
  item, err := cache.Get(key)

  if err != nil {
    log.Print("Error:", err)
  }

  if err != nil {
    fmt.Errorf("Error:", err)
    backendUrl, err := url.Parse(s.BackendUrl)
    if err != nil {
      return
    }
    p := proxy.ProxyServer {
      Url: backendUrl,
      ResponseModifier: responseModifier,
    }
    p.RequestHandler(rw, req)
  } else {
    log.Print("No error: gonna serve from cache!!! ", key)
    cache.ServeHTTP(item, rw, req)
  }
}

// responseModifier is a function provided to ReverseProxy.ModifyResponse.
// It receives the origin server response, and is where we can hook in
// to perform actions such as caching the response.
func responseModifier(res *http.Response, req *http.Request) error {
  // This is where hooks like caching shall happen
  log.Print("ResponseModifier ", req.Method, " ", req.URL.Path, "?", req.URL.RawQuery)
  // TODO: Optionally Cache request responses.
  res.Header["Shiny-Cache"] = []string{"MISS"}

  if err := cacheResponse(res, req); err != nil {
    // Log the cache error, but don't die because of it.
    fmt.Errorf("Error: ResponseModifier failed to cache response")
  }
  return nil
}

func cacheResponse(res *http.Response, req *http.Request) error {
  body, err := ioutil.ReadAll(res.Body)
  res.Body = ioutil.NopCloser(bytes.NewBuffer(body))

  if err != nil {
    fmt.Errorf("Error: Failed to read response Body")
    return err
  }

  key := RequestHash(req)

  item := cache.CacheableResponse{
    Key: key,
    Body: body,
    Header: res.Header,
  }

  if err := cache.Set(key, item.Bytes()); err != nil {
    fmt.Errorf("Error: Failed to set cache", err)
    return err
  }

  return nil
}
