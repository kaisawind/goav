package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>
import "C"
import "unsafe"

const (
	AvSampleFmtNone = int(C.AV_SAMPLE_FMT_NONE)
	AvSampleFmtU8   = int(C.AV_SAMPLE_FMT_U8)
	AvSampleFmtS16  = int(C.AV_SAMPLE_FMT_S16)
	AvSampleFmtS32  = int(C.AV_SAMPLE_FMT_S32)
	AvSampleFmtFlt  = int(C.AV_SAMPLE_FMT_FLT)
	AvSampleFmtDbl  = int(C.AV_SAMPLE_FMT_DBL)

	AvSampleFmtU8P  = int(C.AV_SAMPLE_FMT_U8P)
	AvSampleFmtS16P = int(C.AV_SAMPLE_FMT_S16P)
	AvSampleFmtS32P = int(C.AV_SAMPLE_FMT_S32P)
	AvSampleFmtFltp = int(C.AV_SAMPLE_FMT_FLTP)
	AvSampleFmtDblp = int(C.AV_SAMPLE_FMT_DBLP)
	AvSampleFmtS64  = int(C.AV_SAMPLE_FMT_S64)
	AvSampleFmtS64P = int(C.AV_SAMPLE_FMT_S64P)

	AvSampleFmtNb = int(C.AV_SAMPLE_FMT_NB)
)

func AvGetSampleFmtName(sampleFmt int) string {
	return C.GoString(C.av_get_sample_fmt_name((C.enum_AVSampleFormat)(sampleFmt)))
}

func AvSamplesAlloc(data **uint8, lineSize *int, nbChannels, nbSamples, sampleFmt, align int) int {
	return int(C.av_samples_alloc((**C.uint8_t)(unsafe.Pointer(data)), (*C.int)(unsafe.Pointer(lineSize)), C.int(nbChannels), C.int(nbSamples), (C.enum_AVSampleFormat)(sampleFmt), C.int(align)))
}
