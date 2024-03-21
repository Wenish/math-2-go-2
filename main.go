package main

import (
	"fmt"
	"polynom"
)

func main() {
	fmt.Println("Hallo to the polynom program")

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

	p5 := polynom.NewPolynomByValues(1, 1)

	p6 := p5.Mul(p5).Mul(p5).Mul(p5).Mul(p5)
	fmt.Printf("p6 : %v, grad: %d\n", p6, p6.Grad())

	fmt.Printf("p4(2.56): %.5f\n", p4.Eval(2.56))

	p7 := polynom.NewPolynomByValues(+2, -6, +2, -1)
	fmt.Printf("p7(2): %.5f\n", p7.Eval(1))

	p8 := polynom.NewPolynom(6)
	p9 := polynom.NewPolynomByValues(43, 3, 6)
	p10 := p8.Mul(p9)

	fmt.Printf("10 : %v, grad: %d\n", p10, p10.Grad())
}
