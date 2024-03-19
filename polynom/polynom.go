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
	// Type Polynom Slice erstellen, welches genügen platz hat für das neue Polynom
	// mit max nimt es den grösseren grad wert als initial grid für das neue polynom
	gradNewPolynom := max(p.Grad(), q.Grad())
	result := NewPolynom(gradNewPolynom)

	// das basis polynom zum neuen polynom anfügen
	for i, val := range p {
		result[i] += val
	}
	// das polynom welches addiert werden soll zum neuen polynom anfügen
	for i, val := range q {
		result[i] += val
	}

	for i := len(result); i >= 1; i-- {
		if result[i-1] != 0 {
			return result[:i]
		}
	}

	return result[:1]
}

func (p Polynom) Mul(q Polynom) Polynom {
	// "Der Grad des resultierenden Polynoms ist die Summe der Grade der Ausgangspolynome – Ende." S. Mühlebach (18. Januar 2024)
	gradNewPolynom := p.Grad() + q.Grad()

	// Array erstellen, welches genügen platz hat für das neue Polynom
	result := NewPolynom(gradNewPolynom)

	for i := 0; i <= p.Grad(); i++ {
		for j := 0; j <= q.Grad(); j++ {
			// p[i] * q[j] ergeben multipliziert die summe welche bei dem entsprechenden koeffizent index addiert werden muss
			// i+j ergibt den koeffizent index
			result[i+j] += p[i] * q[j]
		}
	}

	return result
}

func (p Polynom) Eval(x float64) float64 {
	// Mit Horner's Methode das Polynom evaluieren
	sum := p[0]
	for i := 1; i < len(p); i++ {
		sum = sum*x + p[i]
	}

	return sum
}
