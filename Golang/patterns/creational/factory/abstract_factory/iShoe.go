package abstract_factory

type IShoe interface {
  setLogo(logo string)
  setSize(size int)
  GetLogo() string
  GetSize() int
}

type shoe struct {
  logo string
  size int
}

func (s *shoe) setLogo(logo string) {
  s.logo = logo
}

func (s *shoe) setSize(size int) {
  s.size = size
}

func (s *shoe) GetLogo() string {
  return s.logo
}

func (s *shoe) GetSize() int {
  return s.size
}
