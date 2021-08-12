package avutil

//#include <libavutil/rational.h>
import "C"
import "strconv"

func (r Rational) Num() int {
	return int(r.num)
}

func (r *Rational) SetNum(num int) {
	r.num = C.int(num)
}

func (r Rational) Den() int {
	return int(r.den)
}

func (r *Rational) SetDen(den int) {
	r.den = C.int(den)
}

func (r Rational) ToDouble() float64 {
	if r.Num() == 0 || r.Den() == 0 {
		return 0
	}
	return float64(r.Num()) / float64(r.Den())
}

func (r Rational) String() string {
	if r.Num() == 0 || r.Den() == 0 {
		return "0"
	}
	return strconv.Itoa(r.Num()) + "/" + strconv.Itoa(r.Den())
}

func NewRational(num, den int) Rational {
	r := Rational{}
	r.SetNum(num)
	r.SetDen(den)
	return r
}
