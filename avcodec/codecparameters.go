package avcodec

//#include <libavcodec/avcodec.h>
import "C"

func ParametersToContext(codecContext *Context, codecParameters *CodecParameters) int {
	return int(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(codecContext), (*C.struct_AVCodecParameters)(codecParameters)))
}

func ParametersFromContext(codecParameters *CodecParameters, codecContext *Context) int {
	return int(C.avcodec_parameters_from_context((*C.struct_AVCodecParameters)(codecParameters), (*C.struct_AVCodecContext)(codecContext)))
}

func ParametersCopy(out, in *CodecParameters) int {
	return int(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(out), (*C.struct_AVCodecParameters)(in)))
}
