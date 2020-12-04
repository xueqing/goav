package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// codec capabilities AV_CODEC_CAP_xxx
const (

	/**
	 * Decoder can use draw_horiz_band callback.
	 */
	AvCodecCapDrawHorizBand = int(C.AV_CODEC_CAP_DRAW_HORIZ_BAND)
	/**
	 * Codec uses get_buffer() for allocating buffers and supports custom allocators.
	 * If not set, it might not use get_buffer() at all or use operations that
	 * assume the buffer was allocated by avcodec_default_get_buffer.
	 */
	AvCodecCapDr1       = int(C.AV_CODEC_CAP_DR1)
	AvCodecCapTruncated = int(C.AV_CODEC_CAP_TRUNCATED)
	/**
	 * Encoder or decoder requires flushing with NULL input at the end in order to
	 * give the complete and correct output.
	 *
	 * NOTE: If this flag is not set, the codec is guaranteed to never be fed with
	 *       with NULL data. The user can still send NULL data to the public encode
	 *       or decode function, but libavcodec will not pass it along to the codec
	 *       unless this flag is set.
	 *
	 * Decoders:
	 * The decoder has a non-zero delay and needs to be fed with avpkt->data=NULL,
	 * avpkt->size=0 at the end to get the delayed data until the decoder no longer
	 * returns frames.
	 *
	 * Encoders:
	 * The encoder needs to be fed with NULL data at the end of encoding until the
	 * encoder no longer returns data.
	 *
	 * NOTE: For encoders implementing the AVCodec.encode2() function, setting this
	 *       flag also means that the encoder must set the pts and duration for
	 *       each output packet. If this flag is not set, the pts and duration will
	 *       be determined by libavcodec from the input frame.
	 */
	AvCodecCapDelay = int(C.AV_CODEC_CAP_DELAY)
	/**
	 * Codec can be fed a final frame with a smaller size.
	 * This can be used to prevent truncation of the last audio samples.
	 */
	AvCodecCapSmallLastFrame = int(C.AV_CODEC_CAP_SMALL_LAST_FRAME)

	/**
	 * Codec can output multiple frames per AVPacket
	 * Normally demuxers return one frame at a time, demuxers which do not do
	 * are connected to a parser to split what they return into proper frames.
	 * This flag is reserved to the very rare category of codecs which have a
	 * bitstream that cannot be split into frames without timeconsuming
	 * operations like full decoding. Demuxers carrying such bitstreams thus
	 * may return multiple frames in a packet. This has many disadvantages like
	 * prohibiting stream copy in many cases thus it should only be considered
	 * as a last resort.
	 */
	AvCodecCapSubframes = int(C.AV_CODEC_CAP_SUBFRAMES)
	/**
	 * Codec is experimental and is thus avoided in favor of non experimental
	 * encoders
	 */
	AvCodecCapExperimental = int(C.AV_CODEC_CAP_EXPERIMENTAL)
	/**
	 * Codec should fill in channel configuration and samplerate instead of container
	 */
	AvCodecCapChannelConf = int(C.AV_CODEC_CAP_CHANNEL_CONF)
	/**
	 * Codec supports frame-level multithreading.
	 */
	AvCodecCapFrameThreads = int(C.AV_CODEC_CAP_FRAME_THREADS)
	/**
	 * Codec supports slice-based (or partition-based) multithreading.
	 */
	AvCodecCapSliceThreads = int(C.AV_CODEC_CAP_SLICE_THREADS)
	/**
	 * Codec supports changed parameters at any point.
	 */
	AvCodecCapParamChange = int(C.AV_CODEC_CAP_PARAM_CHANGE)
	/**
	 * Codec supports avctx->thread_count == 0 (auto).
	 */
	AvCodecCapAutoThreads = int(C.AV_CODEC_CAP_AUTO_THREADS)
	/**
	 * Audio encoder supports receiving a different number of samples in each call.
	 */
	AvCodecCapVariableFrameSize = int(C.AV_CODEC_CAP_VARIABLE_FRAME_SIZE)
	/**
	 * Decoder is not a preferred choice for probing.
	 * This indicates that the decoder is not a good choice for probing.
	 * It could for example be an expensive to spin up hardware decoder,
	 * or it could simply not provide a lot of useful information about
	 * the stream.
	 * A decoder marked with this flag should only be used as last resort
	 * choice for probing.
	 */
	AvCodecCapAvoidProbing = int(C.AV_CODEC_CAP_AVOID_PROBING)
	/**
	 * Codec is intra only.
	 */
	AvCodecCapIntraOnly = int(C.AV_CODEC_CAP_INTRA_ONLY)
	/**
	 * Codec is lossless.
	 */
	AvCodecCapLossless = int(C.AV_CODEC_CAP_LOSSLESS)

	/**
	 * Codec is backed by a hardware implementation. Typically used to
	 * identify a non-hwaccel hardware decoder. For information about hwaccels, use
	 * avcodec_get_hw_config() instead.
	 */
	AvCodecCapHardware = int(C.AV_CODEC_CAP_HARDWARE)

	/**
	 * Codec is potentially backed by a hardware implementation, but not
	 * necessarily. This is used instead of AV_CODEC_CAP_HARDWARE, if the
	 * implementation provides some sort of internal fallback.
	 */
	AvCodecCapHybrid = int(C.AV_CODEC_CAP_HYBRID)
)
