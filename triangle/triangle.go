package triangle

type Kind int

const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

func KindFromSides(a, b, c float64) Kind {
	UniqueSides := map[float64]bool{a: true, b: true, c: true}
	err := a <= 0 || b <= 0 || c <= 0 || a+b < c || b+c < a || a+c < b
	if !err && len(UniqueSides) == Equ {
		return Kind(Equ) // Equilateral
	}
	if !err && len(UniqueSides) == Iso {
		return Kind(Iso) // Isosceles
	}
	if !err && len(UniqueSides) == Sca {
		return Kind(Sca) // Scalene
	}
	return NaT
}
