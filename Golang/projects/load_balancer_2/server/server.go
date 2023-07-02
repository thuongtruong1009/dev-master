package server

import (
  "net/http/httputil"
  "net/url"
  "sync"
)

type Backend struct {
  URL *url.URL
  Alive bool
  mux sync.RWMutex
  ReverseProxy *httputil.ReverseProxy
}

func  (b *Backend) SetAlive(alive bool) {
  b.mux.RLock()
  b.Alive = alive
  b.mux.RUnlock()
}

func (b *Backend) IsAlive() (alive bool) {
  b.mux.RLock()
  alive = b.Alive
  b.mux.RUnlock()
  return
}
