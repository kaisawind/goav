// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avdevice deals with the input and output devices provided by the libavdevice library
// The libavdevice library provides the same interface as libavformat.
// Namely, an input device is considered like a demuxer, and an output device like a muxer,
// and the interface and generic device options are the same provided by libavformat
package avdevice

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"
import (
	"unsafe"
)

type (
	Rect                  C.struct_AVDeviceRect
	CapabilitiesQuery     C.struct_AVDeviceCapabilitiesQuery
	Info                  C.struct_AVDeviceInfo
	InfoList              C.struct_AVDeviceInfoList
	InputFormat           C.struct_AVInputFormat
	OutputFormat          C.struct_AVOutputFormat
	AvFormatContext       C.struct_AVFormatContext
	AvAppToDevMessageType C.enum_AVAppToDevMessageType
	AvDevToAppMessageType C.enum_AVDevToAppMessageType
)

// Version unsigned 	avdevice_version (void)
func Version() uint {
	return uint(C.avdevice_version())
}

// Configuration Return the libavdevice build-time configuration.
func Configuration() string {
	return C.GoString(C.avdevice_configuration())
}

// License Return the libavdevice license.
func License() string {
	return C.GoString(C.avdevice_license())
}

// RegisterAll Initialize libavdevice and register all the input and output devices.
func RegisterAll() {
	C.avdevice_register_all()
}

// AppToDevControlMessage Send control message from application to device.
func AppToDevControlMessage(s *AvFormatContext, m AvAppToDevMessageType, da int, d uintptr) int {
	return int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s), (C.enum_AVAppToDevMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

// DevToAppControlMessage Send control message from device to application.
func DevToAppControlMessage(fcxt *AvFormatContext, m AvDevToAppMessageType, da int, d uintptr) int {
	return int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(fcxt), (C.enum_AVDevToAppMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

// ListDevices List devices.
func ListDevices(s *AvFormatContext, d **InfoList) int {
	return int(C.avdevice_list_devices((*C.struct_AVFormatContext)(s), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(d))))
}

// FreeListDevices Convenient function to free result of avdeviceListDevices().
func FreeListDevices(d **InfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(d)))
}
