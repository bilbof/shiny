package proxy

import (
  "log"
  "net/url"
  "net/http"
  "strings"
)

// The proxy package contains utility functions for request proxying.
// At the moment it handles requests and forwards them on to a given URL.

// ProxyServer is an HTTP Handler that takes an incoming request and
// sends it to another server, proxying the response back to the
// client.
type ProxyServer struct {
  Url *url.URL
  // ResponseModifier is a function that takes a response and request.
  // It should be used for e.g. caching responses.
  // It should not attempt to respond to the request itself.
  ResponseModifier func(*http.Response, *http.Request) error
}

func (p ProxyServer) RequestHandler(w http.ResponseWriter, req *http.Request) {
  p.newReverseProxy().ServeHTTP(w, req)
}

func (p ProxyServer) newReverseProxy() (*ReverseProxy) {
  log.Print("Handler caught this")
  return &ReverseProxy{
    Director: p.director,
    ModifyResponse: p.ResponseModifier,
  }
}

// The default director, as in the Go 'net/httputil' package.
// https://golang.org/src/net/http/httputil/reverseproxy.go?s=3330:3391#L98
func (p ProxyServer) director(req *http.Request) {
  target := p.Url
  targetQuery := target.RawQuery
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
	if targetQuery == "" || req.URL.RawQuery == "" {
		req.URL.RawQuery = targetQuery + req.URL.RawQuery
	} else {
		req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
	}
	if _, ok := req.Header["User-Agent"]; !ok {
		// explicitly disable User-Agent so it's not set to default value
		req.Header.Set("User-Agent", "")
	}
}

// A Util function from the Go 'net/httputil' package.
// https://golang.org/src/net/http/httputil/reverseproxy.go?s=3330:3391#L98
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
