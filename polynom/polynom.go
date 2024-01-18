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
	// Ensure the resulting polynomial has enough space
	result := make(Polynom, max(p.Grad(), q.Grad())+1)

	// Perform polynomial addition by adding corresponding coefficients
	for i, val := range p {
		result[i] += val
	}
	for i, val := range q {
		result[i] += val
	}

	return result
}

func (p Polynom) Mul(q Polynom) Polynom {
	// Calculate the degree of the resulting polynomial
	deg := p.Grad() + q.Grad()

	// Initialize the resulting polynomial with the appropriate degree
	result := make(Polynom, deg+1)

	// Perform polynomial multiplication using nested loops
	for i := 0; i <= p.Grad(); i++ {
		for j := 0; j <= q.Grad(); j++ {
			result[i+j] += p[i] * q[j]
		}
	}

	return result
}

func (p Polynom) Eval(x float64) float64 {
	// Evaluate the polynomial using Horner's method
	result := 0.0
	for i := p.Grad(); i >= 0; i-- {
		result = result*x + p[i]
	}

	return result
}
