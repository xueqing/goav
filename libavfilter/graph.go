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
func (g *AvFilterGraph) AvfilterGraphAllocFilter(f *AvFilter, n string) *AvFilterContext {
	return (*AvFilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), C.CString(n)))
}

// AvfilterGraphGetFilter Get a filter instance identified by instance name from graph.
func (g *AvFilterGraph) AvfilterGraphGetFilter(n string) *AvFilterContext {
	return (*AvFilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), C.CString(n)))
}

// AvfilterGraphSetAutoConvert Enable or disable automatic format conversion inside the graph.
func (g *AvFilterGraph) AvfilterGraphSetAutoConvert(f uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(f))
}

// AvfilterGraphConfig Check validity and configure all the links and formats in the graph.
func (g *AvFilterGraph) AvfilterGraphConfig(l int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(&l)))
}

// AvfilterGraphFree Free a graph, destroy its links, and set *graph to NULL.
func (g *AvFilterGraph) AvfilterGraphFree() {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(g)))
}

// AvfilterGraphParse Add a graph described by a string to a graph.
func (g *AvFilterGraph) AvfilterGraphParse(f string, i, o *AvFilterInput, l int) int {
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), C.CString(f), (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l)))
}

// AvfilterGraphParsePtr Add a graph described by a string to a graph.
func (g *AvFilterGraph) AvfilterGraphParsePtr(f string, i, o **AvFilterInput, l int) int {
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(&l)))
}

// AvfilterGraphParse2 Add a graph described by a string to a graph.
func (g *AvFilterGraph) AvfilterGraphParse2(f string, i, o **AvFilterInput) int {
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o))))
}

// AvfilterGraphSendCommand Send a command to one or more filter instances.
func (g *AvFilterGraph) AvfilterGraphSendCommand(t, cmd, arg, res string, resl, f int) int {
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.CString(res), C.int(resl), C.int(f)))
}

// AvfilterGraphQueueCommand Queue a command for one or more filter instances.
func (g *AvFilterGraph) AvfilterGraphQueueCommand(t, cmd, arg string, f int, ts C.double) int {
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.int(f), ts))
}

// AvfilterGraphDump Dump a graph into a human-readable string representation.
func (g *AvFilterGraph) AvfilterGraphDump(o string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), C.CString(o)))
}

// AvfilterGraphRequestOldestlink Request a frame on the oldest sink
func (g *AvFilterGraph) AvfilterGraphRequestOldestlink() int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}

// AvfilterGraphCreateFilter Create and add a filter instance into an existing graph.
func AvfilterGraphCreateFilter(cx **AvFilterContext, f *AvFilter, n, a string, o int, g *AvFilterGraph) int {
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), C.CString(n), C.CString(a), unsafe.Pointer(&o), (*C.struct_AVFilterGraph)(g)))
}
