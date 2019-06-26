package oops

type Kitchen struct {
	Plates int
	Size   int
}

type House struct {
	Kitchen
	Rooms      int
	SquareFeet int
}

func (k *Kitchen) GetPlates() int {
	return k.Plates
}

func (k Kitchen) GetSize() int {
	return k.Size
}

func (h House) GetRooms() int {
	return h.Rooms
}

func (h *House) GetSqFt() int {
	return h.SquareFeet
}
func init() {

}
