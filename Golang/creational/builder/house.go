package builder

type House struct {
  windowType string
  doorType string
  floor int
}

func (h *House) GetWindowType() string {
  return h.windowType
}

func (h *House) GetDoorType() string {
  return h.doorType
}

func (h *House) GetNumFloor() int {
  return h.floor
}
