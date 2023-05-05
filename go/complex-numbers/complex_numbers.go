package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	// Real part of the complex num (a)
	real float64
	// Imaginary part of the complex num (b)
	img float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.img
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		real: n1.real + n2.real,
		img:  n1.img + n2.img,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real: n1.real - n2.real,
		img:  n1.img - n2.img,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	// (a+ib)(x+iy) = (ax - by) + i(ay+xb)
	return Number{
		real: n1.real*n2.real - n1.img*n2.img,
		img:  n2.real*n1.img + n1.real*n2.img,
	}
}

func (n Number) Times(factor float64) Number {
	// k(a+ib) -> ka + kbi
	return Number{
		real: n.real * factor,
		img:  n.img * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	/*
		a+ib					(a+ib)(x-iy)				(ax + by) + i(xb - ay)					(ax + by)		(xb - ay)
		------   ->> 	-------------  ->>	--------------------------  ->> --------- + --------- i
		x+iy					(x+iy)(x-iy)								x²+y²											x²+y²				x²+y²
	*/
	x2y2 := math.Pow(n2.real, 2) + math.Pow(n2.img, 2)
	return Number{
		real: (n1.real*n2.real + n1.img*n2.img) / x2y2,
		img:  (n2.real*n1.img - n1.real*n2.img) / x2y2,
	}
}

func (n Number) Conjugate() Number {
	return Number{
		real: n.real,
		img:  -n.img,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.real, 2) + math.Pow(n.img, 2))
}

func (n Number) Exp() Number {
	// e^(a + i * b) = e^a * e^(i * b) = e^a * (cos(b) + i * sin(b)) = e^a*cos(b) + e^a*sin(b)*i
	ea := math.Exp(n.real)
	return Number{
		real: ea * math.Cos(n.img),
		img:  ea * math.Sin(n.img),
	}
}
