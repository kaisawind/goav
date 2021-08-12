package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

const (
	AvPixFmtBgr24    = C.AV_PIX_FMT_BGR24
	AvPixFmtNone     = C.AV_PIX_FMT_NONE
	AvPixFmtRgb24    = C.AV_PIX_FMT_RGB24
	AvPixFmtRgba     = C.AV_PIX_FMT_RGBA
	AvPixFmtYuv420P  = C.AV_PIX_FMT_YUV420P
	AvPixFmtYuvj420P = C.AV_PIX_FMT_YUVJ420P
)

// PixelFormatFromString returns a pixel format from a string
func PixelFormatFromString(i string) PixelFormat {
	switch i {
	case "bgr24":
		return AvPixFmtBgr24
	case "rgb24":
		return AvPixFmtRgb24
	case "rgba":
		return AvPixFmtRgba
	case "yuv420p":
		return AvPixFmtYuv420P
	case "yuvj420p":
		return AvPixFmtYuvj420P
	default:
		return -1
	}
}
