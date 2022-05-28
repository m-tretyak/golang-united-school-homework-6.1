package golang_united_school_homework

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return 3.0 * t.Side
}

//acc. to https://oeis.org/A002194/constant
const sqrt3 float64 = 1.7320508075688772935274463415058723669428052538103806280558069794519330169088000370811461867572485756756261414154
const sqrt3by4 = 0.25 * sqrt3

func (t Triangle) CalcArea() float64 {
	return sqrt3by4 * t.Side * t.Side
}
