package polynom

type Polynom []float64

func NewPolynom(grad int) Polynom {
	return make([]float64, grad+1)
}

func NewPolynomByValues(a ...float64) Polynom {
	p := make([]float64, len(a))
	copy(p, a)
	return p
}

func (p Polynom) Grad() int {
	return len(p) - 1
}

func (p Polynom) Eq(q Polynom) bool {
	if p.Grad() != q.Grad() {
		return false
	}
	for i := range p {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}

func (p Polynom) Add(q Polynom) Polynom {
	// TO DO
	return NewPolynom(3)
}

func (p Polynom) Mul(q Polynom) Polynom {
	// TO DO
	return NewPolynom(3)
}

func (p Polynom) Eval(x float64) float64 {
	// TO DO
	return 53
}
