// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package libavutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

// AVMEDIA_TYPE_xxx
const (
	AvmediaTypeUnknown    = C.AVMEDIA_TYPE_UNKNOWN
	AvmediaTypeVideo      = C.AVMEDIA_TYPE_VIDEO
	AvmediaTypeAudio      = C.AVMEDIA_TYPE_AUDIO
	AvmediaTypeData       = C.AVMEDIA_TYPE_DATA
	AvmediaTypeSubtitle   = C.AVMEDIA_TYPE_SUBTITLE
	AvmediaTypeAttachment = C.AVMEDIA_TYPE_ATTACHMENT
	AvmediaTypeNb         = C.AVMEDIA_TYPE_NB
)
