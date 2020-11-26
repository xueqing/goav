// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package libavfilter contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package libavfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"
)

type (
	AvFilter        C.struct_AVFilter
	AvFilterContext C.struct_AVFilterContext
	AvFilterLink    C.struct_AVFilterLink
	AvFilterGraph   C.struct_AVFilterGraph
	AvFilterInput   C.struct_AVFilterInOut
	AvFilterPad     C.struct_AVFilterPad
	AvDictionary    C.struct_AVDictionary
	AvClass         C.struct_AVClass
	AvMediaType     C.enum_AVMediaType
)

// AvfilterInitStr Initialize a filter with the supplied parameters.
func (ctx *AvFilterContext) AvfilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

// AvfilterInitDict Initialize a filter with the supplied dictionary of options.
func (ctx *AvFilterContext) AvfilterInitDict(options **AvDictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvfilterFree Free a filter context.
func (ctx *AvFilterContext) AvfilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

// AvfilterVersion Return the LIBAvFILTER_VERSION_INT constant.
func AvfilterVersion() uint {
	return uint(C.avfilter_version())
}

// AvfilterConfiguration Return the libavfilter build-time configuration.
func AvfilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

// AvfilterLicense Return the libavfilter license.
func AvfilterLicense() string {
	return C.GoString(C.avfilter_license())
}

// AvfilterPadCount Get the number of elements in a NULL-terminated array of Pads (e.g.
func AvfilterPadCount(pads *AvFilterPad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(pads)))
}

// AvfilterPadGetName Get the name of an AvFilterPad.
func AvfilterPadGetName(pads *AvFilterPad, padIdx int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(pads),
		C.int(padIdx)))
}

// AvfilterPadGetType Get the type of an AvFilterPad.
func AvfilterPadGetType(pads *AvFilterPad, padIdx int) AvMediaType {
	return (AvMediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(pads),
		C.int(padIdx)))
}

// AvfilterLink Link two filters together.
func AvfilterLink(src *AvFilterContext, srcPad uint, dst *AvFilterContext, dstPad uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(src), C.uint(srcPad),
		(*C.struct_AVFilterContext)(dst), C.uint(dstPad)))
}

// AvfilterLinkFree Free the link in *link, and set its pointer to NULL.
func AvfilterLinkFree(link **AvFilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(link)))
}

// AvfilterConfigLinks Negotiate the media format, dimensions, etc of all inputs to a filter.
func AvfilterConfigLinks(filter *AvFilterContext) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(filter)))
}

// AvfilterProcessCommand Make the filter instance process a command.
func AvfilterProcessCommand(filter *AvFilterContext, cmd, arg, res string, resLen, flags int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(filter), C.CString(cmd),
		C.CString(arg), C.CString(res), C.int(resLen), C.int(flags)))
}

// AvfilterInsertFilter Insert a filter in the middle of an existing link.
func AvfilterInsertFilter(link *AvFilterLink, filter *AvFilterContext, filtSrcPadIdx, filtDstPadIdx uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(link),
		(*C.struct_AVFilterContext)(filter), C.uint(filtSrcPadIdx), C.uint(filtDstPadIdx)))
}

// AvfilterGetClass Return avfilter_get_class
func AvfilterGetClass() *AvClass {
	return (*AvClass)(C.avfilter_get_class())
}

// AvfilterInoutAlloc Allocate a single Input entry.
func AvfilterInoutAlloc() *AvFilterInput {
	return (*AvFilterInput)(C.avfilter_inout_alloc())
}

// AvfilterInoutFree Free the supplied list of Input and set *inout to NULL.
func AvfilterInoutFree(inout *AvFilterInput) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(inout)))
}
