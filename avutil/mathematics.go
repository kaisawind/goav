package avutil

//#include <libavutil/mathematics.h>
import "C"

const (
	AvRoundNearInf    = C.AV_ROUND_NEAR_INF
	AvRoundPassMinmax = C.AV_ROUND_PASS_MINMAX
)

func AvRescaleQ(a int64, bq Rational, cq Rational) int64 {
	return (int64)(C.av_rescale_q(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq)))
}

func AvRescaleQRnd(a int64, bq Rational, cq Rational, flags uint32) int64 {
	return (int64)(C.av_rescale_q_rnd(C.int64_t(a), (C.struct_AVRational)(bq), (C.struct_AVRational)(cq), flags))
}
