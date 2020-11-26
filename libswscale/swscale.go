// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package libswscale performs highly optimized image scaling and colorspace and pixel format conversion operations.
//Rescaling: is the process of changing the video size. Several rescaling options and algorithms are available.
//Pixel format conversion: is the process of converting the image format and colorspace of the image.
package libswscale

//#cgo pkg-config: libswscale libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <string.h>
//#include <stdint.h>
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

type (
	SwsContext    C.struct_SwsContext
	SwsFilter     C.struct_SwsFilter
	SwsVector     C.struct_SwsVector
	AvClass       C.struct_AVClass
	AvPixelFormat C.enum_AVPixelFormat
)

// SwscaleVersion Return the LIBSWSCALE_VERSION_INT constant.
func SwscaleVersion() uint {
	return uint(C.swscale_version())
}

// SwscaleConfiguration Return the libswscale build-time configuration.
func SwscaleConfiguration() string {
	return C.GoString(C.swscale_configuration())
}

// SwscaleLicense Return the libswscale license.
func SwscaleLicense() string {
	return C.GoString(C.swscale_license())
}

// SwsGetcoefficients Return a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
func SwsGetcoefficients(colorspace int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(colorspace))))
}

// SwsIssupportedinput Return a positive value if pix_fmt is a supported input format, 0 otherwise.
func SwsIssupportedinput(pixFmt AvPixelFormat) int {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(pixFmt)))
}

// SwsIssupportedoutput Return a positive value if pix_fmt is a supported output format, 0 otherwise.
func SwsIssupportedoutput(pixFmt AvPixelFormat) int {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(pixFmt)))
}

// SwsIssupportedendiannessconversion Return a positive value if an endianness conversion for pix_fmt is supported, 0 otherwise.
func SwsIssupportedendiannessconversion(pixFmt AvPixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(pixFmt)))
}

// SwsScale Scale the image slice in srcSlice and put the resulting scaled slice in the image in dst.
func SwsScale(ctx *SwsContext, srcSlice *uint8, srcStride int, srcSliceY, srcSliceH int, dst *uint8, dstStride int) int {
	cCtx := (*C.struct_SwsContext)(unsafe.Pointer(ctx))
	cSrcSlice := (*C.uint8_t)(unsafe.Pointer(srcSlice))
	cSrcStride := (*C.int)(unsafe.Pointer(&srcStride))
	cDst := (*C.uint8_t)(unsafe.Pointer(dst))
	cDstStride := (*C.int)(unsafe.Pointer(&dstStride))
	return int(C.sws_scale(cCtx, &cSrcSlice, cSrcStride,
		C.int(srcSliceY), C.int(srcSliceH), &cDst, cDstStride))
}

// SwsScale2 refer SwsScale
func SwsScale2(ctx *SwsContext, srcSlice [8]*uint8, srcStride [8]int32, srcSliceY, srcSliceH int, dst [8]*uint8, dstStride [8]int32) int {
	cCtx := (*C.struct_SwsContext)(unsafe.Pointer(ctx))
	cSrcSlice := (**C.uint8_t)(unsafe.Pointer(&srcSlice[0]))
	cSrcStride := (*C.int)(unsafe.Pointer(&srcStride[0]))
	cDst := (**C.uint8_t)(unsafe.Pointer(&dst[0]))
	cDstStride := (*C.int)(unsafe.Pointer(&dstStride))
	return int(C.sws_scale(cCtx, cSrcSlice, cSrcStride,
		C.int(srcSliceY), C.int(srcSliceH), cDst, cDstStride))
}

// SwsSetcolorspacedetails Return -1 if not supported
func SwsSetcolorspacedetails(ctx *SwsContext, invTable *int, srcRange int, table *int, dstRange, brightness, contrast, saturation int) int {
	cInvTable := (*C.int)(unsafe.Pointer(invTable))
	cTable := (*C.int)(unsafe.Pointer(table))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctx),
		cInvTable, C.int(srcRange), cTable, C.int(dstRange),
		C.int(brightness), C.int(contrast), C.int(saturation)))
}

// SwsGetcolorspacedetails Return -1 if not supported
func SwsGetcolorspacedetails(ctx *SwsContext, invTable, srcRange, table, dstRange, brightness, contrast, saturation *int) int {
	cInvTable := (**C.int)(unsafe.Pointer(invTable))
	cSrcRange := (*C.int)(unsafe.Pointer(srcRange))
	cTable := (**C.int)(unsafe.Pointer(table))
	cDstRange := (*C.int)(unsafe.Pointer(dstRange))
	cBrightness := (*C.int)(unsafe.Pointer(brightness))
	cContrast := (*C.int)(unsafe.Pointer(contrast))
	cSaturation := (*C.int)(unsafe.Pointer(saturation))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctx),
		cInvTable, cSrcRange, cTable, cDstRange, cBrightness, cContrast, cSaturation))
}

// SwsGetdefaultfilter Return default fileter
func SwsGetdefaultfilter(lumaGBlur, chromaGBlur, lumaSharpen, chromaSharpen, chromaHShift, chromaVShift float32, verbose int) *SwsFilter {
	return (*SwsFilter)(unsafe.Pointer(C.sws_getDefaultFilter(
		C.float(lumaGBlur), C.float(chromaGBlur), C.float(lumaSharpen), C.float(chromaSharpen),
		C.float(chromaHShift), C.float(chromaVShift), C.int(verbose))))
}

// SwsFreefilter free struct
func SwsFreefilter(filter *SwsFilter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(filter))
}

// SwsConvertpalette8topacked32 Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
func SwsConvertpalette8topacked32(src, dst *uint8, numPixels int, palette *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(src), (*C.uint8_t)(dst),
		C.int(numPixels), (*C.uint8_t)(palette))
}

// SwsConvertpalette8topacked24 Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
func SwsConvertpalette8topacked24(src, dst *uint8, numPixels int, palette *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(src), (*C.uint8_t)(dst),
		C.int(numPixels), (*C.uint8_t)(palette))
}

// SwsGetClass Get the Class for swsContext.
func SwsGetClass() *AvClass {
	return (*AvClass)(C.sws_get_class())
}
