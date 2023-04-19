package builder

type normalBuilder struct {
    windowType string
    doorType string
    floor int
}

func newNormalBuilder() *normalBuilder {
  return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
  b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
  b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
  b.floor = 2
}

func (b *normalBuilder) getHouse() House {
  return House{
    windowType: b.windowType,
    doorType: b.doorType,
    floor: b.floor,
  }
}
