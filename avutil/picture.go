package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/imgutils.h>
import "C"
import "unsafe"

func AvAllocateImageBuffer(size int) []uint8 {
	p := AvMalloc(uintptr(size))
	if p == nil {
		return nil
	}
	return (*[1 << 30]uint8)(p)[:size]
}

func AvFreeImageBuffer(buffer []uint8) {
	AvFree(unsafe.Pointer(&buffer[0]))
}

// AvImageGetBufferSize Return the size in bytes of the amount of data required to store an image with the given parameters.
func AvImageGetBufferSize(pixFmt PixelFormat, width, height, align int) int {
	return int(C.av_image_get_buffer_size((C.enum_AVPixelFormat)(pixFmt), C.int(width), C.int(height), C.int(align)))
}

// AvImageFillArrays Setup the data pointers and linesizes based on the specified image parameters and the provided array.
func AvImageFillArrays(dstData [8]*uint8, dstLinesize [8]int32, src []uint8,
	pixFmt PixelFormat, width, height, align int) int {
	cData := (**C.uint8_t)(unsafe.Pointer(&dstData[0]))
	cLinesize := (*C.int)(unsafe.Pointer(&dstLinesize[0]))
	cSrc := (*C.uint8_t)(unsafe.Pointer(&src[0]))
	cPixFmt := (C.enum_AVPixelFormat)(pixFmt)

	return int(C.av_image_fill_arrays(cData, cLinesize, cSrc, cPixFmt, C.int(width), C.int(height), C.int(align)))
}

// AvImageAlloc Allocate an image with size w and h and pixel format pix_fmt, and fill pointers and linesizes accordingly.
// The allocated image buffer has to be freed by using AvFreeP(&pointers[0]).
func AvImageAlloc(pointers [8]*uint8, linesizes [8]int32, w, h int, pix_fmt PixelFormat, align int) int {
	cPointers := (**C.uint8_t)(unsafe.Pointer(&pointers[0]))
	cLinesizes := (*C.int)(unsafe.Pointer(&linesizes[0]))
	cPixFmt := (C.enum_AVPixelFormat)(pix_fmt)

	return int(C.av_image_alloc(cPointers, cLinesizes, C.int(w), C.int(h), cPixFmt, C.int(align)))
}
