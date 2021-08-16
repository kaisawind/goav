package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/opt.h>
import "C"
import "unsafe"

// AvOptSetDefaults Set the values of all AVOption fields to their default values.
func AvOptSetDefaults(s unsafe.Pointer) {
	C.av_opt_set_defaults(s)
}
