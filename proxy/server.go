package proxy

import (
  "log"
  "net/url"
  "net/http"
  "net/http/httputil"
)

// The proxy package contains utility functions for request proxying.
// At the moment it handles requests and forwards them on to a given URL.

type ProxyServer struct {
  Url *url.URL
}

func (p ProxyServer) newReverseProxy() (*httputil.ReverseProxy) {
  log.Print("Handler caught this")
  return httputil.NewSingleHostReverseProxy(p.Url)
}

func (p ProxyServer) RequestHandler(w http.ResponseWriter, req *http.Request) {
  server := p.newReverseProxy()
  server.ServeHTTP(w, req)
}
