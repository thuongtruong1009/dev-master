package abstract_factory

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
  return &adidasShoe{
    shoe: shoe{
      logo: "adidas",
      size: 14,
    },
  }
}

func (a *Adidas) MakeShort() IShort {
  return &adidasShort{
    short: short{
      logo: "adidas",
      size: 14,
    },
  }
}
