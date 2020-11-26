// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/log.h>
	#include <stdlib.h>
*/
import "C"

// AV_LOG_xxx
const (
	AvLogQuiet   = -8
	AvLogPanic   = 0
	AvLogFatal   = 8
	AvLogError   = 16
	AvLogWarning = 24
	AvLogInfo    = 32
	AvLogVerbose = 40
	AvLogDebug   = 48
	AvLogTrace   = 56
)

// AvLogSetLevel Set the log level
func AvLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}

// AvLogGetLevel Get the current log level
func AvLogGetLevel() int {
	return int(C.av_log_get_level())
}
