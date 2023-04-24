package abstract_factory

type Nike struct {
}

func (n *Nike) MakeShoe() IShoe {
  return &nikeShoe{
    shoe: shoe{
      logo: "Nike",
      size: 14,
    },
  }
}

func (n *Nike) MakeShort() IShort {
  return &nikeShort{
    short: short{
      logo: "Nike",
      size: 14,
    },
  }
}
