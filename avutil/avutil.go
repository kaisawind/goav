// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

import "C"

/*
#cgo pkg-config: libavutil
#include <libavutil/avutil.h>
#include <libavutil/channel_layout.h>
#include <libavutil/pixdesc.h>
#include <stdlib.h>
#include <errno.h>

unsigned major(unsigned v) {
	return AV_VERSION_MAJOR(v);
}

unsigned minor(unsigned v) {
	return AV_VERSION_MINOR(v);
}

unsigned micro(unsigned v) {
	return AV_VERSION_MICRO(v);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type (
	Options       C.struct_AVOptions
	AvTree        C.struct_AVTree
	Rational      C.struct_AVRational
	MediaType     C.enum_AVMediaType
	AvPictureType C.enum_AVPictureType
	PixelFormat   C.enum_AVPixelFormat
	File          C.FILE
)

const (
	AvTimeBase   = C.AV_TIME_BASE
	AvNoPtsValue = C.AV_NOPTS_VALUE
)

var AvTimeBaseQ Rational = NewRational(1, AvTimeBase)

const (
	AvMediaTypeUnknown    = C.AVMEDIA_TYPE_UNKNOWN
	AvMediaTypeVideo      = C.AVMEDIA_TYPE_VIDEO
	AvMediaTypeAudio      = C.AVMEDIA_TYPE_AUDIO
	AvMediaTypeData       = C.AVMEDIA_TYPE_DATA
	AvMediaTypeSubtitle   = C.AVMEDIA_TYPE_SUBTITLE
	AvMediaTypeAttachment = C.AVMEDIA_TYPE_ATTACHMENT
	AvMediaTypeNb         = C.AVMEDIA_TYPE_NB
)

// MediaTypeFromString returns a media type from a string
func MediaTypeFromString(i string) MediaType {
	switch i {
	case "audio":
		return AvMediaTypeAudio
	case "subtitle":
		return AvMediaTypeSubtitle
	case "video":
		return AvMediaTypeVideo
	default:
		return -1
	}
}

const (
	AvChFrontLeft    = 0x1
	AvChFrontRight   = 0x2
	AvChLayoutStereo = 0x3 //(AvChFrontLeft | AvChFrontRight)
)

const (
	AvErrorEagain    = -(C.EAGAIN)
	AvErrorEio       = -(C.EIO)
	AvErrorEof       = C.AVERROR_EOF
	AvErrorEperm     = -(C.EPERM)
	AvErrorEpipe     = -(C.EPIPE)
	AvErrorEtimedout = -(C.ETIMEDOUT)
)

const (
	MaxAverrorStrLen       = 255
	MaxChannelLayoutStrLen = 64
)

const (
	AvPictureTypeNone = C.AV_PICTURE_TYPE_NONE // Undefined
	AvPictureTypeI    = C.AV_PICTURE_TYPE_I
	AvPictureTypeB    = C.AV_PICTURE_TYPE_B
	AvPictureTypeP    = C.AV_PICTURE_TYPE_P
	AvPictureTypeS    = C.AV_PICTURE_TYPE_S
	AvPictureTypeSi   = C.AV_PICTURE_TYPE_SI
	AvPictureTypeSp   = C.AV_PICTURE_TYPE_SP
	AvPictureTypeBi   = C.AV_PICTURE_TYPE_BI
)

func Major(v uint) uint {
	return uint(C.major(C.uint(v)))
}

func Minor(v uint) uint {
	return uint(C.minor(C.uint(v)))
}

func Micro(v uint) uint {
	return uint(C.micro(C.uint(v)))
}

func AvVersion(v uint) string {
	return fmt.Sprintf("%d.%d.%d", Major(v), Minor(v), Micro(v))
}

// Version Return the LIBAvUTIL_VERSION_INT constant.
func Version() uint {
	return uint(C.avutil_version())
}

func VersionInfo() string {
	return C.GoString(C.av_version_info())
}

// Configuration Return the libavutil build-time configuration.
func Configuration() string {
	return C.GoString(C.avutil_configuration())
}

// License Return the libavutil license.
func License() string {
	return C.GoString(C.avutil_license())
}

// AvGetMediaTypeString Return a string describing the media_type enum, NULL if media_type is unknown.
func AvGetMediaTypeString(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

// AvGetPictureTypeChar Return a single letter to describe the given picture type pict_type.
func AvGetPictureTypeChar(pt AvPictureType) byte {
	return byte(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

// AvXIfNull Return x default pointer in case p is NULL.
func AvXIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

// AvIntListLengthForSize Compute the length of an integer list.
func AvIntListLengthForSize(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

// AvFopenUtf8 Open a file using a UTF-8 filename.
func AvFopenUtf8(p, m string) *File {
	cp := C.CString(p)
	defer C.free(unsafe.Pointer(cp))
	cm := C.CString(m)
	defer C.free(unsafe.Pointer(cm))
	f := C.av_fopen_utf8(cp, cm)
	return (*File)(f)
}

// AvGetTimeBaseQ Return the fractional representation of the internal time base.
func AvGetTimeBaseQ() Rational {
	return (Rational)(C.av_get_time_base_q())
}

func AvGetChannelLayoutNbChannels(channelLayout uint64) int {
	return int(C.av_get_channel_layout_nb_channels(C.uint64_t(channelLayout)))
}

func AvGetPixFmtName(pixFmt PixelFormat) string {
	s := C.av_get_pix_fmt_name((C.enum_AVPixelFormat)(pixFmt))
	if s == nil {
		return fmt.Sprintf("unknown pixel format with value %d", pixFmt)
	}
	return C.GoString(s)
}

func AvGetChannelLayoutString(nbChannels int, channelLayout uint64) string {
	bufSize := C.size_t(MaxChannelLayoutStrLen)
	buf := (*C.char)(C.malloc(bufSize))
	if buf == nil {
		return fmt.Sprintf("unknown channel layout with code %d", channelLayout)
	}
	defer C.free(unsafe.Pointer(buf))
	C.av_get_channel_layout_string(buf, C.int(bufSize), C.int(nbChannels), C.uint64_t(channelLayout))
	return C.GoString(buf)
}

func AvStrerr(errcode int) string {
	errbufSize := C.size_t(MaxAverrorStrLen)
	errbuf := (*C.char)(C.malloc(errbufSize))
	if errbuf == nil {
		return fmt.Sprintf("unknown error with code %d", errcode)
	}
	defer C.free(unsafe.Pointer(errbuf))
	ret := C.av_strerror(C.int(errcode), errbuf, errbufSize)
	if ret < 0 {
		return fmt.Sprintf("unknown error with code %d", errcode)
	}
	return C.GoString(errbuf)
}
