package main

import (
	"errors"
	"fmt"
)

func spm(a int, b int) (int, int, int, float64, error) {
	if b == 0 {
		return 0, 0, 0, 0, errors.New("Divide by Zero")
	}
	return a + b, a * b, a - b, float64(a) / float64(b), nil
}

func main() {
	s, p, m, d, err := spm(10, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(s, p, m, d)
}
