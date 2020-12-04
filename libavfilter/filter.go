// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"

// AvfilterGetByName Get a filter definition matching the given name.
func AvfilterGetByName(name string) *AvFilter {
	return (*AvFilter)(C.avfilter_get_by_name(C.CString(name)))
}
