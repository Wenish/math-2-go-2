package main

import (
	"fmt"
	"polynom"
)

func main() {
	fmt.Println("Hallo")

	var p1a, p1b, p2, p3, p4 polynom.Polynom

	p1a = polynom.NewPolynomByValues(-1, +1, -1, +1)
	fmt.Printf("p1a: %v, grad: %d\n", p1a, p1a.Grad())

	p1b = polynom.NewPolynomByValues(+1, -1, +1, -1)
	fmt.Printf("p1b: %v, grad: %d\n", p1b, p1b.Grad())

	p2 = polynom.NewPolynom(5)
	p2[0] = 1
	p2[1] = 2
	p2[2] = 3
	p2[3] = 4
	p2[4] = 5
	p2[5] = 6
	fmt.Printf("p2 : %v, grad: %d\n", p2, p2.Grad())

	p3 = p1a.Add(p1b)
	fmt.Printf("p3 : %v, grad: %d\n", p3, p3.Grad())

	p4 = p1a.Mul(p2)
	fmt.Printf("p4 : %v, grad: %d\n", p4, p4.Grad())

	fmt.Printf("p4(2.56): %.5f\n", p4.Eval(2.56))
}
