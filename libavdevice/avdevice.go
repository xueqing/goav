// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package libavdevice deals with the input and output devices provided by the libavdevice library
// The libavdevice library provides the same interface as libavformat.
// Namely, an input device is considered like a demuxer, and an output device like a muxer,
// and the interface and generic device options are the same provided by libavformat
package libavdevice

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"
import (
	"unsafe"
)

type (
	AvDeviceRect              C.struct_AVDeviceRect
	AvDeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery
	AvDeviceInfo              C.struct_AVDeviceInfo
	AvDeviceInfoList          C.struct_AVDeviceInfoList
	AvInputFormat             C.struct_AVInputFormat
	AvOutputFormat            C.struct_AVOutputFormat
	AvFormatContext           C.struct_AVFormatContext
	AvDictionary              C.struct_AVDictionary
	AvAppToDevMessageType     C.enum_AVAppToDevMessageType
	AvDevToAppMessageType     C.enum_AVDevToAppMessageType
)

// AvdeviceVersion Return the LIBAVDEVICE_VERSION_INT constant.
func AvdeviceVersion() uint {
	return uint(C.avdevice_version())
}

// AvdeviceConfiguration Return the libavdevice build-time configuration.
func AvdeviceConfiguration() string {
	return C.GoString(C.avdevice_configuration())
}

// AvdeviceLicense Return the libavdevice license.
func AvdeviceLicense() string {
	return C.GoString(C.avdevice_license())
}

// AvdeviceRegisterAll Initialize libavdevice and register all the input and output devices.
func AvdeviceRegisterAll() {
	C.avdevice_register_all()
}

// AvdeviceAppToDevControlMessage Send control message from application to device.
func AvdeviceAppToDevControlMessage(fctx *AvFormatContext, typ AvAppToDevMessageType, data int, dataSize uintptr) int {
	return int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(fctx),
		(C.enum_AVAppToDevMessageType)(typ), unsafe.Pointer(&data), C.size_t(dataSize)))
}

// AvdeviceDevToAppControlMessage Send control message from device to application.
func AvdeviceDevToAppControlMessage(fctx *AvFormatContext, typ AvDevToAppMessageType, data int, dataSize uintptr) int {
	return int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(fctx),
		(C.enum_AVDevToAppMessageType)(typ), unsafe.Pointer(&data), C.size_t(dataSize)))
}

// AvdeviceCapabilitiesCreate Initialize capabilities probing API based on AvOption API.
func AvdeviceCapabilitiesCreate(caps **AvDeviceCapabilitiesQuery, fctx *AvFormatContext, deviceOptions **AvDictionary) int {
	return int(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)),
		(*C.struct_AVFormatContext)(fctx), (**C.struct_AVDictionary)(unsafe.Pointer(deviceOptions))))
}

// AvdeviceCapabilitiesFree Free resources created by avdevice_capabilities_create()
func AvdeviceCapabilitiesFree(caps **AvDeviceCapabilitiesQuery, fctx *AvFormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(caps)),
		(*C.struct_AVFormatContext)(fctx))
}

// AvdeviceListDevices List devices.
func AvdeviceListDevices(fctx *AvFormatContext, deviceList **AvDeviceInfoList) int {
	return int(C.avdevice_list_devices((*C.struct_AVFormatContext)(fctx),
		(**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList))))
}

// AvdeviceFreeListDevices Convenient function to free result of avdeviceListDevices().
func AvdeviceFreeListDevices(deviceList **AvDeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(deviceList)))
}

// //int 	avdevice_list_input_sources (struct AvInputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// //List devices.
// func AvdeviceListInputSources(d *AvInputFormat, dv string, do *AvDictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(d), C.CString(dv), (*C.struct_AVDictionary)(do), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }

// //int 	avdevice_list_output_sinks (struct AvOutputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// func AvdeviceListOutputSinks(d *AvOutputFormat, dn string, di *AvDictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(d), C.CString(dn), (*C.struct_AVDictionary)(di), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }
