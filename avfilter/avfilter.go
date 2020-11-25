// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avfilter contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"
)

type (
	Filter     C.struct_AVFilter
	Context    C.struct_AVFilterContext
	Link       C.struct_AVFilterLink
	Graph      C.struct_AVFilterGraph
	Input      C.struct_AVFilterInOut
	Pad        C.struct_AVFilterPad
	Dictionary C.struct_AVDictionary
	Class      C.struct_AVClass
	MediaType  C.enum_AVMediaType
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
func AvfilterPadCount(p *Pad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

// AvfilterPadGetName Get the name of an Pad.
func AvfilterPadGetName(p *Pad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

// AvfilterPadGetType Get the type of an Pad.
func AvfilterPadGetType(p *Pad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

// AvfilterLink Link two filters together.
func AvfilterLink(s *Context, sp uint, d *Context, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

// AvfilterLinkFree Free the link in *link, and set its pointer to NULL.
func AvfilterLinkFree(l **Link) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

// AvfilterConfigLinks Negotiate the media format, dimensions, etc of all inputs to a filter.
func AvfilterConfigLinks(f *Context) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

// AvfilterProcessCommand Make the filter instance process a command.
func AvfilterProcessCommand(f *Context, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

// AvfilterInitStr Initialize a filter with the supplied parameters.
func (ctx *Context) AvfilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

// AvfilterInitDict Initialize a filter with the supplied dictionary of options.
func (ctx *Context) AvfilterInitDict(o **Dictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

// AvfilterFree Free a filter context.
func (ctx *Context) AvfilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

// AvfilterInsertFilter Insert a filter in the middle of an existing link.
func AvfilterInsertFilter(l *Link, f *Context, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

// AvfilterGetClass_get_class
func AvfilterGetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

// AvfilterInoutAlloc Allocate a single Input entry.
func AvfilterInoutAlloc() *Input {
	return (*Input)(C.avfilter_inout_alloc())
}

// AvfilterInoutFree Free the supplied list of Input and set *inout to NULL.
func AvfilterInoutFree(i *Input) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}
