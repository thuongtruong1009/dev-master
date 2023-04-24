package abstract_factory

type IShort interface {
  setLogo(logo string)
  setSize(size int)
  GetLogo() string
  GetSize() int
}

type short struct {
  logo string
  size int
}

func (s *short) setLogo(logo string) {
  s.logo = logo
}

func (s *short) setSize(size int) {
  s.size = size
}

func (s *short) GetLogo() string {
  return s.logo
}

func (s *short) GetSize() int {
  return s.size
}
