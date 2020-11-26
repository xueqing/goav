// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"
import "github.com/xueqing/goav/libavcodec"

func newRational(r C.struct_AVRational) libavcodec.Rational {
	return libavcodec.NewRational(int(r.num), int(r.den))
}
