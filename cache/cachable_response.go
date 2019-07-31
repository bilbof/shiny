package cache

import (
  "bytes"
  "net/http"
  "encoding/gob"
)

// CacheableResponse is a middleman data structure that sits between
// http.Response and cache Items.

type CacheableResponse struct {
  Key string
  Body []byte
  Header http.Header
}

func (c CacheableResponse) Bytes() []byte {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(&c)
	return b.Bytes()
}

func BytesToCacheableResponse(b []byte) CacheableResponse {
	var c CacheableResponse
	dec := gob.NewDecoder(bytes.NewReader(b))
	dec.Decode(&c) // todo catch error
	return c
}
