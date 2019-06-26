package main

import (
	"fmt"

	"github.com/balaji39mit/Programs/go-training/oops"
)

func main() {
	h := oops.House{oops.Kitchen{5, 10}, 20, 30}
	fmt.Println(h.GetPlates())
	fmt.Println(h.GetRooms())
	fmt.Println(h.GetSize())
	fmt.Println(h.GetSqFt())
}
