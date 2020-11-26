// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package libavcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package libavcodec

//#cgo pkg-config: libavformat libavcodec libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libswscale/swscale.h>
import "C"

// AV_PIX_FMT_xxx
const (
	AvPixFmtYuv        = 0
	AvPixFmtYuv420P9   = C.AV_PIX_FMT_YUV420P9
	AvPixFmtYuv422P9   = C.AV_PIX_FMT_YUV422P9
	AvPixFmtYuv444P9   = C.AV_PIX_FMT_YUV444P9
	AvPixFmtYuv420P10  = C.AV_PIX_FMT_YUV420P10
	AvPixFmtYuv422P10  = C.AV_PIX_FMT_YUV422P10
	AvPixFmtYuv440P10  = C.AV_PIX_FMT_YUV440P10
	AvPixFmtYuv444P10  = C.AV_PIX_FMT_YUV444P10
	AvPixFmtYuv420P12  = C.AV_PIX_FMT_YUV420P12
	AvPixFmtYuv422P12  = C.AV_PIX_FMT_YUV422P12
	AvPixFmtYuv440P12  = C.AV_PIX_FMT_YUV440P12
	AvPixFmtYuv444P12  = C.AV_PIX_FMT_YUV444P12
	AvPixFmtYuv420P14  = C.AV_PIX_FMT_YUV420P14
	AvPixFmtYuv422P14  = C.AV_PIX_FMT_YUV422P14
	AvPixFmtYuv444P14  = C.AV_PIX_FMT_YUV444P14
	AvPixFmtYuv420P16  = C.AV_PIX_FMT_YUV420P16
	AvPixFmtYuv422P16  = C.AV_PIX_FMT_YUV422P16
	AvPixFmtYuv444P16  = C.AV_PIX_FMT_YUV444P16
	AvPixFmtYuvA420P9  = C.AV_PIX_FMT_YUVA420P9
	AvPixFmtYuvA422P9  = C.AV_PIX_FMT_YUVA422P9
	AvPixFmtYuvA444P9  = C.AV_PIX_FMT_YUVA444P9
	AvPixFmtYuvA420P10 = C.AV_PIX_FMT_YUVA420P10
	AvPixFmtYuvA422P10 = C.AV_PIX_FMT_YUVA422P10
	AvPixFmtYuvA444P10 = C.AV_PIX_FMT_YUVA444P10
	AvPixFmtYuvA420P16 = C.AV_PIX_FMT_YUVA420P16
	AvPixFmtYuvA422P16 = C.AV_PIX_FMT_YUVA422P16
	AvPixFmtYuvA444P16 = C.AV_PIX_FMT_YUVA444P16
	AvPixFmtRgb24      = C.AV_PIX_FMT_RGB24
	AvPixFmtRgba       = C.AV_PIX_FMT_RGBA
)

// SWS_xxx
const (
	SwsFastBilinear     = C.SWS_FAST_BILINEAR
	SwsBilinear         = C.SWS_BILINEAR
	SwsBicubic          = C.SWS_BICUBIC
	SwsX                = C.SWS_X
	SwsPoint            = C.SWS_POINT
	SwsArea             = C.SWS_AREA
	SwsBicublin         = C.SWS_BICUBLIN
	SwsGauss            = C.SWS_GAUSS
	SwsSinc             = C.SWS_SINC
	SwsLanczos          = C.SWS_LANCZOS
	SwsSpline           = C.SWS_SPLINE
	SwsSrcVChrDropMask  = C.SWS_SRC_V_CHR_DROP_MASK
	SwsSrcVChrDropShift = C.SWS_SRC_V_CHR_DROP_SHIFT
	SwsParamDefault     = C.SWS_PARAM_DEFAULT
	SwsPrintInfo        = C.SWS_PRINT_INFO
	SwsFullChrHInt      = C.SWS_FULL_CHR_H_INT
	SwsFullChrHInp      = C.SWS_FULL_CHR_H_INP
	SwsDirectBgr        = C.SWS_DIRECT_BGR
	SwsAccurateRnd      = C.SWS_ACCURATE_RND
	SwsBitexact         = C.SWS_BITEXACT
	SwsErrorDiffusion   = C.SWS_ERROR_DIFFUSION
	SwsMaxReduceCutoff  = C.SWS_MAX_REDUCE_CUTOFF
	SwsCsItu709         = C.SWS_CS_ITU709
	SwsCsFcc            = C.SWS_CS_FCC
	SwsCsItu601         = C.SWS_CS_ITU601
	SwsCsItu624         = C.SWS_CS_ITU624
	SwsCsSmpte170m      = C.SWS_CS_SMPTE170M
	SwsCsSmpte240m      = C.SWS_CS_SMPTE240M
	SwsCsDefault        = C.SWS_CS_DEFAULT
	SwsCsBt2020         = C.SWS_CS_BT2020
)

// String ...
func (pf PixelFormat) String() string {
	switch int(pf) {
	case AvPixFmtYuv420P9:
		return "YUV420P9"

	case AvPixFmtYuv422P9:
		return "YUV422P9"

	case AvPixFmtYuv444P9:
		return "YUV444P9"

	case AvPixFmtYuv420P10:
		return "YUV420P10"

	case AvPixFmtYuv422P10:
		return "YUV422P10"

	case AvPixFmtYuv440P10:
		return "YUV440P10"

	case AvPixFmtYuv444P10:
		return "YUV444P10"

	case AvPixFmtYuv420P12:
		return "YUV420P12"

	case AvPixFmtYuv422P12:
		return "YUV422P12"

	case AvPixFmtYuv440P12:
		return "YUV440P12"

	case AvPixFmtYuv444P12:
		return "YUV444P12"

	case AvPixFmtYuv420P14:
		return "YUV420P14"

	case AvPixFmtYuv422P14:
		return "YUV422P14"

	case AvPixFmtYuv444P14:
		return "YUV444P14"

	case AvPixFmtYuv420P16:
		return "YUV420P16"

	case AvPixFmtYuv422P16:
		return "YUV422P16"

	case AvPixFmtYuv444P16:
		return "YUV444P16"

	case AvPixFmtYuvA420P9:
		return "YUVA420P9"

	case AvPixFmtYuvA422P9:
		return "YUVA422P9"

	case AvPixFmtYuvA444P9:
		return "YUVA444P9"

	case AvPixFmtYuvA420P10:
		return "YUVA420P10"

	case AvPixFmtYuvA422P10:
		return "YUVA422P10"

	case AvPixFmtYuvA444P10:
		return "YUVA444P10"

	case AvPixFmtYuvA420P16:
		return "YUVA420P16"

	case AvPixFmtYuvA422P16:
		return "YUVA422P16"

	case AvPixFmtYuvA444P16:
		return "YUVA444P16"

	case AvPixFmtRgb24:
		return "RGB24"

	case AvPixFmtRgba:
		return "RGBA"
	}

	return "{UNKNOWN}"
}
