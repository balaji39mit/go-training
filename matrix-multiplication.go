package main

func main() {
	a := []int{1, 2, 3}
	m := map[int]int{
		0: 4,
		1: 5,
		2: 6,
	}
	s := make([]int, len(a), len(a))
	for i, v := range a {
		s[i] = v * m[i]
	}

	//m1 := [][]int{{1, 2, 3}}
	//
	// m2 := [][]int{{4}, {5}, {6}}

}
