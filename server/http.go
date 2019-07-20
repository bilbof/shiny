package server

import (
    "flag"
    "net/http"
    // "log"
    // "fmt"
    "net/url"
    "github.com/bilbof/shiny/proxy"
)

// The server package runs the HTTP server that listens to all requests.
// At the moment it just forwards requests to the proxy server.

type Server struct {
  Port  string
  BackendUrl string
}

func (s Server) ListenAndServe() (error) {
  backendUrl, err := url.Parse(s.BackendUrl)
  if err != nil {
    return err
  }
  p := proxy.ProxyServer {
    Url: backendUrl,
  }

  http.Handle("/", http.HandlerFunc(p.RequestHandler))

  var addr = flag.String("addr", s.Port, "http service address")
  flag.Parse()
  return http.ListenAndServe(*addr, nil)
}
