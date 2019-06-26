package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	s := a[:2:2]
	fmt.Println(a, s, len(s), cap(s))
	s = append(s, 4)
	fmt.Println(a, s, len(s), cap(s))
	//s = append(s, 3)
	//fmt.Println(a, s, len(s), cap(s))
}
