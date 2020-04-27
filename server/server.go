package server

import (
  "flag"
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
  log.Print("✨ Shiny listening at ", s.Port, " ✨")
  http.Handle("/", http.HandlerFunc(s.requestHandler))
  var addr = flag.String("addr", s.Port, "http service address")
  flag.Parse()
  return http.ListenAndServe(*addr, nil)
}

// requestHandler takes a request and passes it either to the cache
// or to the proxy server
func (s Server) requestHandler(rw http.ResponseWriter, req *http.Request)  {
  item, err := cache.Get(req)

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
    cache.ServeHTTP(item, rw, req)
  }
}

// responseModifier is a function provided to ReverseProxy.ModifyResponse.
// It receives the origin server response, and is where we can hook in
// to perform actions such as caching the response.
func responseModifier(res *http.Response, req *http.Request) error {
  res.Header["S-Cache"] = []string{"MISS"}

  if err := cache.CacheResponse(req, res); err != nil {
    // Log the cache error, but don't die because of it.
    fmt.Errorf("Error: failed to cache response for ", req.Method, " ", req.URL.Path, "?", req.URL.RawQuery)
  }
  return nil
}
