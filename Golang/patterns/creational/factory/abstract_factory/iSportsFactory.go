package abstract_factory

type ISportsfactory interface {
  MakeShoe() IShoe
  MakeShort() IShort
}

func GetSportsFactory(brand string) ISportsfactory {
  if brand == "adidas" {
    return &Adidas{}
  }
  if brand == "nike" {
    return &Nike{}
  }
  return nil
}
