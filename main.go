package main

import (
	"fmt"
	"polynom"
)

func main() {
	fmt.Println("Hallo")

	var p1, p2, p3, p4 polynom.Polynom

	p1 = polynom.NewPolynom(3)
	p1[0] = -1
	p1[1] = +1
	p1[2] = -1
	p1[3] = +1
	fmt.Printf("p1: %v\n", p1)

	p2 = polynom.NewPolynom(5)
	p2[0] = 1
	p2[1] = 2
	p2[2] = 3
	p2[3] = 4
	p2[4] = 5
	p2[5] = 6
	fmt.Printf("p2: %v\n", p2)

	p3 = p1.Add(p2)
	fmt.Printf("p3: %v\n", p3)

	p4 = p1.Mul(p2)
	fmt.Printf("p4: %v\n", p4)
}
