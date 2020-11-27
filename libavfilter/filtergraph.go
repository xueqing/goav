// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"
)

// AvfilterGraphAlloc Allocate a filter graph.
func AvfilterGraphAlloc() *AvFilterGraph {
	return (*AvFilterGraph)(C.avfilter_graph_alloc())
}

// AvfilterGraphAllocFilter Create a new filter instance in a filter graph.
func (g *AvFilterGraph) AvfilterGraphAllocFilter(filter *AvFilter, name string) *AvFilterContext {
	return (*AvFilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g),
		(*C.struct_AVFilter)(filter), C.CString(name)))
}

// AvfilterGraphGetFilter Get a filter instance identified by instance name from graph.
func (g *AvFilterGraph) AvfilterGraphGetFilter(name string) *AvFilterContext {
	return (*AvFilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g),
		C.CString(name)))
}

// AvfilterGraphSetAutoConvert Enable or disable automatic format conversion inside the graph.
func (g *AvFilterGraph) AvfilterGraphSetAutoConvert(flags uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(flags))
}

// AvfilterGraphConfig Check validity and configure all the links and formats in the graph.
func (g *AvFilterGraph) AvfilterGraphConfig(logCtx int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g),
		unsafe.Pointer(&logCtx)))
}

// AvfilterGraphFree Free a graph, destroy its links, and set *graph to NULL.
func (g *AvFilterGraph) AvfilterGraphFree() {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(g)))
}

// AvfilterGraphParse Add a graph described by a string to a graph.
func (g *AvFilterGraph) AvfilterGraphParse(filters string, inputs, outputs *AvFilterInOut, logCtx int) int {
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g),
		C.CString(filters), (*C.struct_AVFilterInOut)(inputs),
		(*C.struct_AVFilterInOut)(outputs), unsafe.Pointer(&logCtx)))
}

// AvfilterGraphParsePtr Add a graph described by a string to a graph.
func (g *AvFilterGraph) AvfilterGraphParsePtr(filters string, inputs, outputs **AvFilterInOut, logCtx int) int {
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g),
		C.CString(filters), (**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(outputs)), unsafe.Pointer(&logCtx)))
}

// AvfilterGraphParse2 Add a graph described by a string to a graph.
func (g *AvFilterGraph) AvfilterGraphParse2(filters string, inputs, outputs **AvFilterInOut) int {
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g),
		C.CString(filters), (**C.struct_AVFilterInOut)(unsafe.Pointer(inputs)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(outputs))))
}

// AvfilterGraphSendCommand Send a command to one or more filter instances.
func (g *AvFilterGraph) AvfilterGraphSendCommand(target, cmd, arg, res string, resLen, flags int) int {
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g),
		C.CString(target), C.CString(cmd), C.CString(arg), C.CString(res),
		C.int(resLen), C.int(flags)))
}

// AvfilterGraphQueueCommand Queue a command for one or more filter instances.
func (g *AvFilterGraph) AvfilterGraphQueueCommand(target, cmd, arg, res string, flags int, ts C.double) int {
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g),
		C.CString(target), C.CString(cmd), C.CString(arg), C.int(flags), ts))
}

// AvfilterGraphDump Dump a graph into a human-readable string representation.
func (g *AvFilterGraph) AvfilterGraphDump(options string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g),
		C.CString(options)))
}

// AvfilterGraphRequestOldestlink Request a frame on the oldest sink
func (g *AvFilterGraph) AvfilterGraphRequestOldestlink() int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}

// AvfilterGraphCreateFilter Create and add a filter instance into an existing graph.
func AvfilterGraphCreateFilter(filterContext **AvFilterContext, filter *AvFilter, name, args string, opaque int, g *AvFilterGraph) int {
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(filterContext)),
		(*C.struct_AVFilter)(filter), C.CString(name), C.CString(args),
		unsafe.Pointer(&opaque), (*C.struct_AVFilterGraph)(g)))
}
