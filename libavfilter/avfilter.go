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
func AvfilterPadCount(p *AvFilterPad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

// AvfilterPadGetName Get the name of an AvFilterPad.
func AvfilterPadGetName(p *AvFilterPad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

// AvfilterPadGetType Get the type of an AvFilterPad.
func AvfilterPadGetType(p *AvFilterPad, pi int) AvMediaType {
	return (AvMediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

// AvfilterLink Link two filters together.
func AvfilterLink(s *AvFilterContext, sp uint, d *AvFilterContext, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

// AvfilterLinkFree Free the link in *link, and set its pointer to NULL.
func AvfilterLinkFree(l **AvFilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

// AvfilterConfigLinks Negotiate the media format, dimensions, etc of all inputs to a filter.
func AvfilterConfigLinks(f *AvFilterContext) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

// AvfilterProcessCommand Make the filter instance process a command.
func AvfilterProcessCommand(f *AvFilterContext, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

// AvfilterInitStr Initialize a filter with the supplied parameters.
func (ctx *AvFilterContext) AvfilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

// AvfilterInitDict Initialize a filter with the supplied dictionary of options.
func (ctx *AvFilterContext) AvfilterInitDict(o **AvDictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

// AvfilterFree Free a filter context.
func (ctx *AvFilterContext) AvfilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

// AvfilterInsertFilter Insert a filter in the middle of an existing link.
func AvfilterInsertFilter(l *AvFilterLink, f *AvFilterContext, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
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
func AvfilterInoutFree(i *AvFilterInput) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}
