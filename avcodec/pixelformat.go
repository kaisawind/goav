// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"

	"github.com/asticode/goav/avutil"
)

// PixFmtToCodecTag Return a value representing the fourCC code associated to the pixel format pix_fmt, or 0 if no associated fourCC code can be found.
func PixFmtToCodecTag(pixFmt avutil.PixelFormat) uint {
	return uint(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(pixFmt)))
}

// FindBestPixFmtOfList Find the best pixel format to convert to given a certain source pixel format.
func FindBestPixFmtOfList(pixFmtList *avutil.PixelFormat, srcPixFmt avutil.PixelFormat, a int, l *int) avutil.PixelFormat {
	return (avutil.PixelFormat)(C.avcodec_find_best_pix_fmt_of_list((*C.enum_AVPixelFormat)(pixFmtList), (C.enum_AVPixelFormat)(srcPixFmt), C.int(a), (*C.int)(unsafe.Pointer(l))))
}
