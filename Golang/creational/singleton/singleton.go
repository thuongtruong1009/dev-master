package singleton

type Singleton interface {
  AddOne() int
}

type singleton struct {
  count int
}

var instance *singleton

// init value when package is loaded
func init() {
  instance = &singleton{count: 100}
}

func GetInstance() Singleton {
  return instance
}

func (s *singleton) AddOne() int {
  s.count++
  return s.count
}
