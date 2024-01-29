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
	// Array erstellen, welches genügen platz hat für das neue Polynom
	result := make(Polynom, max(p.Grad(), q.Grad())+1)

	// das basis polynom zum neuen polynom anfügen
	for i, val := range p {
		result[i] += val
	}
	// das polynom welches addiert werden soll zum neuen polynom anfügen
	for i, val := range q {
		result[i] += val
	}

	// Entferne Nullwerte aus dem result
	finalResult := removeZeroValues(result)

	// Sicherstellen, dass mindestens ein Element im finalResult bleibt
	if len(finalResult) == 0 {
		finalResult = append(finalResult, 0)
	}

	// das neue zusammengezählte polynom zurückgeben
	return finalResult
}

// Funktion zum Entfernen von Nullwerten aus einem Polynom
func removeZeroValues(p Polynom) Polynom {
	var result Polynom
	for _, val := range p {
		if val != 0 {
			result = append(result, val)
		}
	}
	return result
}

func (p Polynom) Mul(q Polynom) Polynom {
	// "Der Grad des resultierenden Polynoms ist die Summe der Grade der Ausgangspolynome – Ende." S. Mühlebach (18. Januar 2024)
	gradNewPolynom := p.Grad() + q.Grad()

	// Initialize the resulting polynomial with the appropriate degree
	result := make(Polynom, gradNewPolynom+1)

	// Perform polynomial multiplication using nested loops
	for i := 0; i <= p.Grad(); i++ {
		for j := 0; j <= q.Grad(); j++ {
			result[i+j] += p[i] * q[j]
		}
	}

	return result
}

func (p Polynom) Eval(x float64) float64 {
	// Mit Horner's Methode das Polynom evaluieren
	result := 0.0
	for i := p.Grad(); i >= 0; i-- {
		result = result*x + p[i]
	}

	return result
}
