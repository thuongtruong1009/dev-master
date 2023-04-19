package builder

type IBuilder interface {
  setWindowType()
  setDoorType()
  setNumFloor()
  getHouse() House
}

func GetBuilder(builderType string) IBuilder {
  switch builderType {
    case "normal":
      return &normalBuilder{}
    case "igloo":
      return &iglooBuilder{}
    }
  return nil
}

