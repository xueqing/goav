// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// AV_CODEC_FLAG_xxx
const (
	AvCodecFlagUnaligned          = int(C.AV_CODEC_FLAG_UNALIGNED)
	AvCodecFlagQscale             = int(C.AV_CODEC_FLAG_QSCALE)
	AvCodecFlag4mv                = int(C.AV_CODEC_FLAG_4MV)
	AvCodecFlagOutputCorrupt      = int(C.AV_CODEC_FLAG_OUTPUT_CORRUPT)
	AvCodecFlagQpel               = int(C.AV_CODEC_FLAG_QPEL)
	AvCodecFlagPass1              = int(C.AV_CODEC_FLAG_PASS1)
	AvCodecFlagPass2              = int(C.AV_CODEC_FLAG_PASS2)
	AvCodecFlagLoopFilter         = int(C.AV_CODEC_FLAG_LOOP_FILTER)
	AvCodecFlagGray               = int(C.AV_CODEC_FLAG_GRAY)
	AvCodecFlagPsnr               = int(C.AV_CODEC_FLAG_PSNR)
	AvCodecFlagTruncated          = int(C.AV_CODEC_FLAG_TRUNCATED)
	AvCodecFlagInterlacedDct      = int(C.AV_CODEC_FLAG_INTERLACED_DCT)
	AvCodecFlagLowDelay           = int(C.AV_CODEC_FLAG_LOW_DELAY)
	AvCodecFlagGlobalHeader       = int(C.AV_CODEC_FLAG_GLOBAL_HEADER)
	AvCodecFlagBitexact           = int(C.AV_CODEC_FLAG_BITEXACT)
	AvCodecFlagAcPred             = int(C.AV_CODEC_FLAG_AC_PRED)
	AvCodecFlagInterlacedMe       = int(C.AV_CODEC_FLAG_INTERLACED_ME)
	AvCodecFlagClosedGop          = int(C.AV_CODEC_FLAG_CLOSED_GOP)
	AvCodecFlag2Fast              = int(C.AV_CODEC_FLAG2_FAST)
	AvCodecFlag2NoOutput          = int(C.AV_CODEC_FLAG2_NO_OUTPUT)
	AvCodecFlag2LocalHeader       = int(C.AV_CODEC_FLAG2_LOCAL_HEADER)
	AvCodecFlag2DropFrameTimecode = int(C.AV_CODEC_FLAG2_DROP_FRAME_TIMECODE)
	AvCodecFlag2Chunks            = int(C.AV_CODEC_FLAG2_CHUNKS)
	AvCodecFlag2IgnoreCrop        = int(C.AV_CODEC_FLAG2_IGNORE_CROP)
	AvCodecFlag2ShowAll           = int(C.AV_CODEC_FLAG2_SHOW_ALL)
	AvCodecFlag2ExportMvs         = int(C.AV_CODEC_FLAG2_EXPORT_MVS)
	AvCodecFlag2SkipManual        = int(C.AV_CODEC_FLAG2_SKIP_MANUAL)
	AvCodecFlag2RoFlushNoop       = int(C.AV_CODEC_FLAG2_RO_FLUSH_NOOP)
)
