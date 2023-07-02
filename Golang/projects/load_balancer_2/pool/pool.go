package pool

type ServerPool struct {
  backends []*Backend
  current unint64
}

func (s *ServerPool) AddBackend(backend *Backend) {
  s.Backend = append(s.Backend, backend)
}

func (s *ServerPool) NextIndex() int {
  return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}

func (s *ServerPool) MarkBackendStatus(backendUrl *url.URL, alive bool) {
  for _, b := range s.backends {
    if b.URL.String() == backendUrl.String()  {
      b.SetAlive(alive)
      break
    }
  }
}

func (s *ServerPool) GetNextPeer() *Backend {
  next := s.NextIndex()
  l := len(s.backends) + next
  for i := next; i < l; i++ {
    idx := i % len(s.backends)
    if s.backends[idx].IsAlive() {
      if i != next {
        atomic.StoreUint64(&s.current, uint64(idx))
      }
      return s.backends[idx]
    }
  }
  return nil
}

func (s *ServerPool) HealthCheck() {
  for _, b := range s.backends {
    status := "up"
    alive := isBackendAlive(b.URL)
    b.SetAlive(alive)
    if !alive {
      status = "down"
    }
    log.Printf("%s [%s]\n", b.URL, status)
  }
}
