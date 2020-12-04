package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

// AVFMT_xxx
const (
	/// Demuxer will use avio_open, no opened file should be provided by the caller.
	AvfmtNofile       = C.AVFMT_NOFILE
	AvfmtNeednumber   = C.AVFMT_NEEDNUMBER    /**< Needs '%d' in filename. */
	AvfmtShowIds      = C.AVFMT_SHOW_IDS      /**< Show format stream IDs numbers. */
	AvfmtGlobalheader = C.AVFMT_GLOBALHEADER  /**< Format wants global header. */
	AvfmtNotimestamps = C.AVFMT_NOTIMESTAMPS  /**< Format does not need / have any timestamps. */
	AvfmtGenericIndex = C.AVFMT_GENERIC_INDEX /**< Use generic index building code. */
	AvfmtTsDiscont    = C.AVFMT_TS_DISCONT    /**< Format allows timestamp discontinuities. Note, muxers always require valid (monotone) timestamps */
	AvfmtVariableFps  = C.AVFMT_VARIABLE_FPS  /**< Format allows variable fps. */
	AvfmtNodimensions = C.AVFMT_NODIMENSIONS  /**< Format does not need width/height */
	AvfmtNostreams    = C.AVFMT_NOSTREAMS     /**< Format does not require any streams */
	AvfmtNobinsearch  = C.AVFMT_NOBINSEARCH   /**< Format does not allow to fall back on binary search via read_timestamp */
	AvfmtNogensearch  = C.AVFMT_NOGENSEARCH   /**< Format does not allow to fall back on generic search */
	AvfmtNoByteSeek   = C.AVFMT_NO_BYTE_SEEK  /**< Format does not allow seeking by bytes */
	AvfmtAllowFlush   = C.AVFMT_ALLOW_FLUSH   /**< Format allows flushing. If not set, the muxer will not receive a NULL packet in the write_packet function. */
	AvfmtTsNonstrict  = C.AVFMT_TS_NONSTRICT  /**< Format does not require strictly increasing timestamps, but they must still be monotonic */
	AvfmtTsNegative   = C.AVFMT_TS_NEGATIVE   /**< Format allows muxing negative timestamps. If not set the timestamp will be shifted in av_write_frame
	and av_interleaved_write_frame so they start from 0. The user or muxer can override this through AVFormatContext.avoid_negative_ts */
	AvfmtSeekToPts = C.AVFMT_SEEK_TO_PTS /**< Seeking is based on PTS */
)
