package float

import "text/template"

var baseTmpl = template.Must(template.New("generated").Parse(`
type Float float64

func ZeroFloat() Float {
	return 0
}

func MaxFloat() Float {
	return math.MaxFloat64
}

func MinFloat() Float {
	return math.SmallestNonzeroFloat64
}

func (rcv Float) E() Float {
	return math.E
}

func (rcv Float) Pi() Float {
	return math.Pi
}

func (rcv Float) Phi() Float {
	return math.Phi
}

func ToFloat(x float64) Float {
	return (Float)(x)
}

// Max returns the larger of x or y.
func (rcv Float) Max(y Float) Float {
	return rcv.curry2(math.Max)(y)
}

// Min returns the smaller of x or y.
func (rcv Float) Min(y Float) Float {
	return rcv.curry2(math.Min)(y)
}

// Dim returns the maximum of x-y or 0.
func (rcv Float) Dim(y Float) Float {
	return rcv.curry2(math.Dim)(y)
}

func (rcv Float) Abs() Float {
	return rcv.Apply(math.Abs)
}

func (rcv Float) Sin() Float {
	return rcv.Apply(math.Sin)
}

// Asin returns the arcsine, in radians, of x.
func (rcv Float) Asin() Float {
	return rcv.Apply(math.Asin)
}

func (rcv Float) Cos() Float {
	return rcv.Apply(math.Cos)
}

// Acosh returns the inverse hyperbolic cosine of x.
func (rcv Float) Acos() Float {
	return rcv.Apply(math.Acos)
}

// Tan returns the tangent of the radian argument x.
func (rcv Float) Tan() Float {
	return rcv.Apply(math.Tan)
}

// Returns radian angle between -pi/2 and +pi/2 whose tangent is x.
func (rcv Float) Atan() Float {
	return rcv.Apply(math.Atan)
}
// Sincos returns Sin(x), Cos(x).
func (rcv Float) Sincos() (Float, Float) {
	x, y := math.Sincos(rcv.Unwrap())
	return ToFloat(x), ToFloat(y)
}

// Sqrt returns the square root of x.
func (rcv Float) Sqrt() Float {
	return rcv.Apply(math.Sqrt)
}

// Cbrt returns the cube root of x.
func (rcv Float) Cbrt() Float {
	return rcv.Apply(math.Cbrt)
}

func (rcv Float) Pow(y Float) Float {
	return rcv.curry2(math.Pow)(y)
}

// Floor returns the greatest integer value less than or equal to x.
func (rcv Float) Floor() Float {
	return rcv.Apply(math.Floor)
}

// Ceil returns the least integer value greater than or equal to x.
func (rcv Float) Ceil() Float {
	return rcv.Apply(math.Ceil)
}

// Pow10 returns 10**n, the base-10 exponential of n.
func (rcv Float) Pow10(n int) Float {
	return ToFloat(math.Pow10(n))
}

// Remainder returns the IEEE 754 floating-point remainder of x/y.
func (rcv Float) Remainder(y Float) Float {
	return rcv.curry2(math.Remainder)(y)
}

// Log returns the natural logarithm of x.
func (rcv Float) Log() Float {
	return rcv.Apply(math.Log)
}

// Log10 returns the decimal logarithm of x.
func (rcv Float) Log10() Float {
	return rcv.Apply(math.Log10)
}

// Log2 returns the binary logarithm of x.
func (rcv Float) Log2() Float {
	return rcv.Apply(math.Log2)
}

// Log1p returns the natural logarithm of 1 plus its argument x.
func (rcv Float) Log1p() Float {
	return rcv.Apply(math.Log1p)
}

func (rcv Float) Multiply(x Float) Float {
	return rcv * x
}

func (rcv Float) Add(x Float) Float {
	return rcv + x
}

func (rcv Float) Subtract(x Float) Float {
	return rcv - x
}

func (rcv Float) Unwrap() float64 {
	return (float64)(rcv)
}

func (rcv Float) Println() {
	fmt.Println(rcv)
}

func (rcv Float) MkString() String {
	return String(fmt.Sprintf("%v", rcv.Unwrap()))
}

func (rcv Float) ApplyF(fn func(Float) Float) Float {
	return fn(rcv)
}

func (rcv Float) Apply(fn func(float64) float64) Float {
	return ToFloat(fn(rcv.Unwrap()))
}

// curried and 1st function applied with rcv
func (rcv Float) curry2(fn func(float64, float64) float64) func (Float) Float {
	f := curry2(fn)(rcv.Unwrap())
	return wrapFloat11(f)
}

func wrapFloat11(fn func(float64) float64) func (Float) Float {
	return func(x Float) Float {
		return ToFloat(fn(x.Unwrap()))
	}
}

func curry2(fn func(float64, float64) float64) func (float64) func (float64) float64 {
	return func(x float64) func(float64) float64 {
		return func(y float64) float64 {
			return fn(x, y)
		}
	}
}
`))

var importsTemplate = template.Must(template.New("generated").Parse(`
import (
	{{range $index, $import := .Imports}}"{{$import}}"
	{{end}}
)
`))
