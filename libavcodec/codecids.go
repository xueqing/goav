// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

// AV_CODEC_IDxxx Multiple encoders have the same ID and are able to encode compatible streams.
const (
	AvCodecID012v            = int(C.AV_CODEC_ID_012V)
	AvCodecID4xm             = int(C.AV_CODEC_ID_4XM)
	AvCodecID8bps            = int(C.AV_CODEC_ID_8BPS)
	AvCodecID8svxExp         = int(C.AV_CODEC_ID_8SVX_EXP)
	AvCodecID8svxFib         = int(C.AV_CODEC_ID_8SVX_FIB)
	AvCodecIDA64Multi        = int(C.AV_CODEC_ID_A64_MULTI)
	AvCodecIDA64Multi5       = int(C.AV_CODEC_ID_A64_MULTI5)
	AvCodecIDAac             = int(C.AV_CODEC_ID_AAC)
	AvCodecIDAacLatm         = int(C.AV_CODEC_ID_AAC_LATM)
	AvCodecIDAasc            = int(C.AV_CODEC_ID_AASC)
	AvCodecIDAc3             = int(C.AV_CODEC_ID_AC3)
	AvCodecIDAdpcm4xm        = int(C.AV_CODEC_ID_ADPCM_4XM)
	AvCodecIDAdpcmAdx        = int(C.AV_CODEC_ID_ADPCM_ADX)
	AvCodecIDAdpcmAfc        = int(C.AV_CODEC_ID_ADPCM_AFC)
	AvCodecIDAdpcmCt         = int(C.AV_CODEC_ID_ADPCM_CT)
	AvCodecIDAdpcmDtk        = int(C.AV_CODEC_ID_ADPCM_DTK)
	AvCodecIDAdpcmEa         = int(C.AV_CODEC_ID_ADPCM_EA)
	AvCodecIDAdpcmEaMaxisXa  = int(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	AvCodecIDAdpcmEaR1       = int(C.AV_CODEC_ID_ADPCM_EA_R1)
	AvCodecIDAdpcmEaR2       = int(C.AV_CODEC_ID_ADPCM_EA_R2)
	AvCodecIDAdpcmEaR3       = int(C.AV_CODEC_ID_ADPCM_EA_R3)
	AvCodecIDAdpcmEaXas      = int(C.AV_CODEC_ID_ADPCM_EA_XAS)
	AvCodecIDAdpcmG722       = int(C.AV_CODEC_ID_ADPCM_G722)
	AvCodecIDAdpcmG726       = int(C.AV_CODEC_ID_ADPCM_G726)
	AvCodecIDAdpcmG726le     = int(C.AV_CODEC_ID_ADPCM_G726LE)
	AvCodecIDAdpcmImaAmv     = int(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	AvCodecIDAdpcmImaApc     = int(C.AV_CODEC_ID_ADPCM_IMA_APC)
	AvCodecIDAdpcmImaDk3     = int(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	AvCodecIDAdpcmImaDk4     = int(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	AvCodecIDAdpcmImaEaEacs  = int(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	AvCodecIDAdpcmImaEaSead  = int(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	AvCodecIDAdpcmImaIss     = int(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	AvCodecIDAdpcmImaOki     = int(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	AvCodecIDAdpcmImaQt      = int(C.AV_CODEC_ID_ADPCM_IMA_QT)
	AvCodecIDAdpcmImaRad     = int(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	AvCodecIDAdpcmImaSmjpeg  = int(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	AvCodecIDAdpcmImaWav     = int(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	AvCodecIDAdpcmImaWs      = int(C.AV_CODEC_ID_ADPCM_IMA_WS)
	AvCodecIDAdpcmMs         = int(C.AV_CODEC_ID_ADPCM_MS)
	AvCodecIDAdpcmSbpro2     = int(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	AvCodecIDAdpcmSbpro3     = int(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	AvCodecIDAdpcmSbpro4     = int(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	AvCodecIDAdpcmSwf        = int(C.AV_CODEC_ID_ADPCM_SWF)
	AvCodecIDAdpcmThp        = int(C.AV_CODEC_ID_ADPCM_THP)
	AvCodecIDAdpcmVima       = int(C.AV_CODEC_ID_ADPCM_VIMA)
	AvCodecIDAdpcmXa         = int(C.AV_CODEC_ID_ADPCM_XA)
	AvCodecIDAdpcmYamaha     = int(C.AV_CODEC_ID_ADPCM_YAMAHA)
	AvCodecIDAic             = int(C.AV_CODEC_ID_AIC)
	AvCodecIDAlac            = int(C.AV_CODEC_ID_ALAC)
	AvCodecIDAliasPix        = int(C.AV_CODEC_ID_ALIAS_PIX)
	AvCodecIDAmrNb           = int(C.AV_CODEC_ID_AMR_NB)
	AvCodecIDAmrWb           = int(C.AV_CODEC_ID_AMR_WB)
	AvCodecIDAmv             = int(C.AV_CODEC_ID_AMV)
	AvCodecIDAnm             = int(C.AV_CODEC_ID_ANM)
	AvCodecIDAnsi            = int(C.AV_CODEC_ID_ANSI)
	AvCodecIDApe             = int(C.AV_CODEC_ID_APE)
	AvCodecIDAss             = int(C.AV_CODEC_ID_ASS)
	AvCodecIDAsv1            = int(C.AV_CODEC_ID_ASV1)
	AvCodecIDAsv2            = int(C.AV_CODEC_ID_ASV2)
	AvCodecIDAtrac1          = int(C.AV_CODEC_ID_ATRAC1)
	AvCodecIDAtrac3          = int(C.AV_CODEC_ID_ATRAC3)
	AvCodecIDAtrac3p         = int(C.AV_CODEC_ID_ATRAC3P)
	AvCodecIDAura            = int(C.AV_CODEC_ID_AURA)
	AvCodecIDAura2           = int(C.AV_CODEC_ID_AURA2)
	AvCodecIDAvrn            = int(C.AV_CODEC_ID_AVRN)
	AvCodecIDAvrp            = int(C.AV_CODEC_ID_AVRP)
	AvCodecIDAvs             = int(C.AV_CODEC_ID_AVS)
	AvCodecIDAvui            = int(C.AV_CODEC_ID_AVUI)
	AvCodecIDAyuv            = int(C.AV_CODEC_ID_AYUV)
	AvCodecIDBethsoftvid     = int(C.AV_CODEC_ID_BETHSOFTVID)
	AvCodecIDBfi             = int(C.AV_CODEC_ID_BFI)
	AvCodecIDBinkaudioDct    = int(C.AV_CODEC_ID_BINKAUDIO_DCT)
	AvCodecIDBinkaudioRdft   = int(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	AvCodecIDBinkvideo       = int(C.AV_CODEC_ID_BINKVIDEO)
	AvCodecIDBintext         = int(C.AV_CODEC_ID_BINTEXT)
	AvCodecIDBinData         = int(C.AV_CODEC_ID_BIN_DATA)
	AvCodecIDBmp             = int(C.AV_CODEC_ID_BMP)
	AvCodecIDBmvAudio        = int(C.AV_CODEC_ID_BMV_AUDIO)
	AvCodecIDBmvVideo        = int(C.AV_CODEC_ID_BMV_VIDEO)
	AvCodecIDBrenderPix      = int(C.AV_CODEC_ID_BRENDER_PIX)
	AvCodecIDC93             = int(C.AV_CODEC_ID_C93)
	AvCodecIDCavs            = int(C.AV_CODEC_ID_CAVS)
	AvCodecIDCdgraphics      = int(C.AV_CODEC_ID_CDGRAPHICS)
	AvCodecIDCdxl            = int(C.AV_CODEC_ID_CDXL)
	AvCodecIDCelt            = int(C.AV_CODEC_ID_CELT)
	AvCodecIDCinepak         = int(C.AV_CODEC_ID_CINEPAK)
	AvCodecIDCljr            = int(C.AV_CODEC_ID_CLJR)
	AvCodecIDCllc            = int(C.AV_CODEC_ID_CLLC)
	AvCodecIDCmv             = int(C.AV_CODEC_ID_CMV)
	AvCodecIDComfortNoise    = int(C.AV_CODEC_ID_COMFORT_NOISE)
	AvCodecIDCook            = int(C.AV_CODEC_ID_COOK)
	AvCodecIDCpia            = int(C.AV_CODEC_ID_CPIA)
	AvCodecIDCscd            = int(C.AV_CODEC_ID_CSCD)
	AvCodecIDCyuv            = int(C.AV_CODEC_ID_CYUV)
	AvCodecIDDfa             = int(C.AV_CODEC_ID_DFA)
	AvCodecIDDirac           = int(C.AV_CODEC_ID_DIRAC)
	AvCodecIDDnxhd           = int(C.AV_CODEC_ID_DNXHD)
	AvCodecIDDpx             = int(C.AV_CODEC_ID_DPX)
	AvCodecIDDsdLsbf         = int(C.AV_CODEC_ID_DSD_LSBF)
	AvCodecIDDsdLsbfPlanar   = int(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	AvCodecIDDsdMsbf         = int(C.AV_CODEC_ID_DSD_MSBF)
	AvCodecIDDsdMsbfPlanar   = int(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	AvCodecIDDsicinaudio     = int(C.AV_CODEC_ID_DSICINAUDIO)
	AvCodecIDDsicinvideo     = int(C.AV_CODEC_ID_DSICINVIDEO)
	AvCodecIDDts             = int(C.AV_CODEC_ID_DTS)
	AvCodecIDDvaudio         = int(C.AV_CODEC_ID_DVAUDIO)
	AvCodecIDDvbSubtitle     = int(C.AV_CODEC_ID_DVB_SUBTITLE)
	AvCodecIDDvbTeletext     = int(C.AV_CODEC_ID_DVB_TELETEXT)
	AvCodecIDDvdNav          = int(C.AV_CODEC_ID_DVD_NAV)
	AvCodecIDDvdSubtitle     = int(C.AV_CODEC_ID_DVD_SUBTITLE)
	AvCodecIDDvvideo         = int(C.AV_CODEC_ID_DVVIDEO)
	AvCodecIDDxa             = int(C.AV_CODEC_ID_DXA)
	AvCodecIDDxtory          = int(C.AV_CODEC_ID_DXTORY)
	AvCodecIDEac3            = int(C.AV_CODEC_ID_EAC3)
	AvCodecIDEia608          = int(C.AV_CODEC_ID_EIA_608)
	AvCodecIDEscape124       = int(C.AV_CODEC_ID_ESCAPE124)
	AvCodecIDEscape130       = int(C.AV_CODEC_ID_ESCAPE130)
	AvCodecIDEvrc            = int(C.AV_CODEC_ID_EVRC)
	AvCodecIDExr             = int(C.AV_CODEC_ID_EXR)
	AvCodecIDFfmetadata      = int(C.AV_CODEC_ID_FFMETADATA)
	AvCodecIDFfv1            = int(C.AV_CODEC_ID_FFV1)
	AvCodecIDFfvhuff         = int(C.AV_CODEC_ID_FFVHUFF)
	AvCodecIDFfwavesynth     = int(C.AV_CODEC_ID_FFWAVESYNTH)
	AvCodecIDFic             = int(C.AV_CODEC_ID_FIC)
	AvCodecIDFirstAudio      = int(C.AV_CODEC_ID_FIRST_AUDIO)
	AvCodecIDFirstSubtitle   = int(C.AV_CODEC_ID_FIRST_SUBTITLE)
	AvCodecIDFirstUnknown    = int(C.AV_CODEC_ID_FIRST_UNKNOWN)
	AvCodecIDFlac            = int(C.AV_CODEC_ID_FLAC)
	AvCodecIDFlashsv         = int(C.AV_CODEC_ID_FLASHSV)
	AvCodecIDFlashsv2        = int(C.AV_CODEC_ID_FLASHSV2)
	AvCodecIDFlic            = int(C.AV_CODEC_ID_FLIC)
	AvCodecIDFlv1            = int(C.AV_CODEC_ID_FLV1)
	AvCodecIDFraps           = int(C.AV_CODEC_ID_FRAPS)
	AvCodecIDFrwu            = int(C.AV_CODEC_ID_FRWU)
	AvCodecIDG2m             = int(C.AV_CODEC_ID_G2M)
	AvCodecIDG723_1          = int(C.AV_CODEC_ID_G723_1)
	AvCodecIDG729            = int(C.AV_CODEC_ID_G729)
	AvCodecIDGif             = int(C.AV_CODEC_ID_GIF)
	AvCodecIDGsm             = int(C.AV_CODEC_ID_GSM)
	AvCodecIDGsmMs           = int(C.AV_CODEC_ID_GSM_MS)
	AvCodecIDH261            = int(C.AV_CODEC_ID_H261)
	AvCodecIDH263            = int(C.AV_CODEC_ID_H263)
	AvCodecIDH263i           = int(C.AV_CODEC_ID_H263I)
	AvCodecIDH263p           = int(C.AV_CODEC_ID_H263P)
	AvCodecIDH264            = int(C.AV_CODEC_ID_H264)
	AvCodecIDH265            = int(C.AV_CODEC_ID_H265)
	AvCodecIDHdmvPgsSubtitle = int(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	AvCodecIDHevc            = int(C.AV_CODEC_ID_HEVC)
	AvCodecIDHnm4Video       = int(C.AV_CODEC_ID_HNM4_VIDEO)
	AvCodecIDHuffyuv         = int(C.AV_CODEC_ID_HUFFYUV)
	AvCodecIDIac             = int(C.AV_CODEC_ID_IAC)
	AvCodecIDIdcin           = int(C.AV_CODEC_ID_IDCIN)
	AvCodecIDIdf             = int(C.AV_CODEC_ID_IDF)
	AvCodecIDIffByterun1     = int(C.AV_CODEC_ID_IFF_BYTERUN1)
	AvCodecIDIffIlbm         = int(C.AV_CODEC_ID_IFF_ILBM)
	AvCodecIDIlbc            = int(C.AV_CODEC_ID_ILBC)
	AvCodecIDImc             = int(C.AV_CODEC_ID_IMC)
	AvCodecIDIndeo2          = int(C.AV_CODEC_ID_INDEO2)
	AvCodecIDIndeo3          = int(C.AV_CODEC_ID_INDEO3)
	AvCodecIDIndeo4          = int(C.AV_CODEC_ID_INDEO4)
	AvCodecIDIndeo5          = int(C.AV_CODEC_ID_INDEO5)
	AvCodecIDInterplayDpcm   = int(C.AV_CODEC_ID_INTERPLAY_DPCM)
	AvCodecIDInterplayVideo  = int(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	AvCodecIDJacosub         = int(C.AV_CODEC_ID_JACOSUB)
	AvCodecIDJpeg2000        = int(C.AV_CODEC_ID_JPEG2000)
	AvCodecIDJpegls          = int(C.AV_CODEC_ID_JPEGLS)
	AvCodecIDJv              = int(C.AV_CODEC_ID_JV)
	AvCodecIDKgv1            = int(C.AV_CODEC_ID_KGV1)
	AvCodecIDKmvc            = int(C.AV_CODEC_ID_KMVC)
	AvCodecIDLagarith        = int(C.AV_CODEC_ID_LAGARITH)
	AvCodecIDLjpeg           = int(C.AV_CODEC_ID_LJPEG)
	AvCodecIDLoco            = int(C.AV_CODEC_ID_LOCO)
	AvCodecIDMace3           = int(C.AV_CODEC_ID_MACE3)
	AvCodecIDMace6           = int(C.AV_CODEC_ID_MACE6)
	AvCodecIDMad             = int(C.AV_CODEC_ID_MAD)
	AvCodecIDMdec            = int(C.AV_CODEC_ID_MDEC)
	AvCodecIDMetasound       = int(C.AV_CODEC_ID_METASOUND)
	AvCodecIDMicrodvd        = int(C.AV_CODEC_ID_MICRODVD)
	AvCodecIDMimic           = int(C.AV_CODEC_ID_MIMIC)
	AvCodecIDMjpeg           = int(C.AV_CODEC_ID_MJPEG)
	AvCodecIDMjpegb          = int(C.AV_CODEC_ID_MJPEGB)
	AvCodecIDMlp             = int(C.AV_CODEC_ID_MLP)
	AvCodecIDMmvideo         = int(C.AV_CODEC_ID_MMVIDEO)
	AvCodecIDMotionpixels    = int(C.AV_CODEC_ID_MOTIONPIXELS)
	AvCodecIDMovText         = int(C.AV_CODEC_ID_MOV_TEXT)
	AvCodecIDMp1             = int(C.AV_CODEC_ID_MP1)
	AvCodecIDMp2             = int(C.AV_CODEC_ID_MP2)
	AvCodecIDMp3             = int(C.AV_CODEC_ID_MP3)
	AvCodecIDMp3adu          = int(C.AV_CODEC_ID_MP3ADU)
	AvCodecIDMp3on4          = int(C.AV_CODEC_ID_MP3ON4)
	AvCodecIDMp4als          = int(C.AV_CODEC_ID_MP4ALS)
	AvCodecIDMpeg1video      = int(C.AV_CODEC_ID_MPEG1VIDEO)
	AvCodecIDMpeg2ts         = int(C.AV_CODEC_ID_MPEG2TS)
	AvCodecIDMpeg2video      = int(C.AV_CODEC_ID_MPEG2VIDEO)
	// AV_CODEC_ID_MPEG2VIDEO_XVMC   = int(C.AV_CODEC_ID_MPEG2VIDEO_XVMC)
	AvCodecIDMpeg4           = int(C.AV_CODEC_ID_MPEG4)
	AvCodecIDMpeg4systems    = int(C.AV_CODEC_ID_MPEG4SYSTEMS)
	AvCodecIDMpl2            = int(C.AV_CODEC_ID_MPL2)
	AvCodecIDMsa1            = int(C.AV_CODEC_ID_MSA1)
	AvCodecIDMsmpeg4v1       = int(C.AV_CODEC_ID_MSMPEG4V1)
	AvCodecIDMsmpeg4v2       = int(C.AV_CODEC_ID_MSMPEG4V2)
	AvCodecIDMsmpeg4v3       = int(C.AV_CODEC_ID_MSMPEG4V3)
	AvCodecIDMsrle           = int(C.AV_CODEC_ID_MSRLE)
	AvCodecIDMss1            = int(C.AV_CODEC_ID_MSS1)
	AvCodecIDMss2            = int(C.AV_CODEC_ID_MSS2)
	AvCodecIDMsvideo1        = int(C.AV_CODEC_ID_MSVIDEO1)
	AvCodecIDMszh            = int(C.AV_CODEC_ID_MSZH)
	AvCodecIDMts2            = int(C.AV_CODEC_ID_MTS2)
	AvCodecIDMusepack7       = int(C.AV_CODEC_ID_MUSEPACK7)
	AvCodecIDMusepack8       = int(C.AV_CODEC_ID_MUSEPACK8)
	AvCodecIDMvc1            = int(C.AV_CODEC_ID_MVC1)
	AvCodecIDMvc2            = int(C.AV_CODEC_ID_MVC2)
	AvCodecIDMxpeg           = int(C.AV_CODEC_ID_MXPEG)
	AvCodecIDNellymoser      = int(C.AV_CODEC_ID_NELLYMOSER)
	AvCodecIDNone            = int(C.AV_CODEC_ID_NONE)
	AvCodecIDNuv             = int(C.AV_CODEC_ID_NUV)
	AvCodecIDOn2avc          = int(C.AV_CODEC_ID_ON2AVC)
	AvCodecIDOpus            = int(C.AV_CODEC_ID_OPUS)
	AvCodecIDOtf             = int(C.AV_CODEC_ID_OTF)
	AvCodecIDPafAudio        = int(C.AV_CODEC_ID_PAF_AUDIO)
	AvCodecIDPafVideo        = int(C.AV_CODEC_ID_PAF_VIDEO)
	AvCodecIDPam             = int(C.AV_CODEC_ID_PAM)
	AvCodecIDPbm             = int(C.AV_CODEC_ID_PBM)
	AvCodecIDPcmAlaw         = int(C.AV_CODEC_ID_PCM_ALAW)
	AvCodecIDPcmBluray       = int(C.AV_CODEC_ID_PCM_BLURAY)
	AvCodecIDPcmDvd          = int(C.AV_CODEC_ID_PCM_DVD)
	AvCodecIDPcmF32be        = int(C.AV_CODEC_ID_PCM_F32BE)
	AvCodecIDPcmF32le        = int(C.AV_CODEC_ID_PCM_F32LE)
	AvCodecIDPcmF64be        = int(C.AV_CODEC_ID_PCM_F64BE)
	AvCodecIDPcmF64le        = int(C.AV_CODEC_ID_PCM_F64LE)
	AvCodecIDPcmLxf          = int(C.AV_CODEC_ID_PCM_LXF)
	AvCodecIDPcmMulaw        = int(C.AV_CODEC_ID_PCM_MULAW)
	AvCodecIDPcmS16be        = int(C.AV_CODEC_ID_PCM_S16BE)
	AvCodecIDPcmS16bePlanar  = int(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	AvCodecIDPcmS16le        = int(C.AV_CODEC_ID_PCM_S16LE)
	AvCodecIDPcmS16lePlanar  = int(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	AvCodecIDPcmS24be        = int(C.AV_CODEC_ID_PCM_S24BE)
	AvCodecIDPcmS24daud      = int(C.AV_CODEC_ID_PCM_S24DAUD)
	AvCodecIDPcmS24le        = int(C.AV_CODEC_ID_PCM_S24LE)
	AvCodecIDPcmS24lePlanaR  = int(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	AvCodecIDPcmS32be        = int(C.AV_CODEC_ID_PCM_S32BE)
	AvCodecIDPcmS32le        = int(C.AV_CODEC_ID_PCM_S32LE)
	AvCodecIDPcmS32lePlanaR  = int(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	AvCodecIDPcmS8           = int(C.AV_CODEC_ID_PCM_S8)
	AvCodecIDPcmS8Planar     = int(C.AV_CODEC_ID_PCM_S8_PLANAR)
	AvCodecIDPcmU16be        = int(C.AV_CODEC_ID_PCM_U16BE)
	AvCodecIDPcmU16le        = int(C.AV_CODEC_ID_PCM_U16LE)
	AvCodecIDPcmU24be        = int(C.AV_CODEC_ID_PCM_U24BE)
	AvCodecIDPcmU24le        = int(C.AV_CODEC_ID_PCM_U24LE)
	AvCodecIDPcmU32be        = int(C.AV_CODEC_ID_PCM_U32BE)
	AvCodecIDPcmU32le        = int(C.AV_CODEC_ID_PCM_U32LE)
	AvCodecIDPcmU8           = int(C.AV_CODEC_ID_PCM_U8)
	AvCodecIDPcmZork         = int(C.AV_CODEC_ID_PCM_ZORK)
	AvCodecIDPcx             = int(C.AV_CODEC_ID_PCX)
	AvCodecIDPgm             = int(C.AV_CODEC_ID_PGM)
	AvCodecIDPgmyuv          = int(C.AV_CODEC_ID_PGMYUV)
	AvCodecIDPictor          = int(C.AV_CODEC_ID_PICTOR)
	AvCodecIDPjs             = int(C.AV_CODEC_ID_PJS)
	AvCodecIDPng             = int(C.AV_CODEC_ID_PNG)
	AvCodecIDPpm             = int(C.AV_CODEC_ID_PPM)
	AvCodecIDProbe           = int(C.AV_CODEC_ID_PROBE)
	AvCodecIDProres          = int(C.AV_CODEC_ID_PRORES)
	AvCodecIDPtx             = int(C.AV_CODEC_ID_PTX)
	AvCodecIDQcelp           = int(C.AV_CODEC_ID_QCELP)
	AvCodecIDQdm2            = int(C.AV_CODEC_ID_QDM2)
	AvCodecIDQdmc            = int(C.AV_CODEC_ID_QDMC)
	AvCodecIDQdraw           = int(C.AV_CODEC_ID_QDRAW)
	AvCodecIDQpeg            = int(C.AV_CODEC_ID_QPEG)
	AvCodecIDQtrle           = int(C.AV_CODEC_ID_QTRLE)
	AvCodecIDR10k            = int(C.AV_CODEC_ID_R10K)
	AvCodecIDR210            = int(C.AV_CODEC_ID_R210)
	AvCodecIDRalf            = int(C.AV_CODEC_ID_RALF)
	AvCodecIDRawvideo        = int(C.AV_CODEC_ID_RAWVIDEO)
	AvCodecIDRa144           = int(C.AV_CODEC_ID_RA_144)
	AvCodecIDRa288           = int(C.AV_CODEC_ID_RA_288)
	AvCodecIDRealtext        = int(C.AV_CODEC_ID_REALTEXT)
	AvCodecIDRl2             = int(C.AV_CODEC_ID_RL2)
	AvCodecIDRoq             = int(C.AV_CODEC_ID_ROQ)
	AvCodecIDRoqDpcm         = int(C.AV_CODEC_ID_ROQ_DPCM)
	AvCodecIDRpza            = int(C.AV_CODEC_ID_RPZA)
	AvCodecIDRv10            = int(C.AV_CODEC_ID_RV10)
	AvCodecIDRv20            = int(C.AV_CODEC_ID_RV20)
	AvCodecIDRv30            = int(C.AV_CODEC_ID_RV30)
	AvCodecIDRv40            = int(C.AV_CODEC_ID_RV40)
	AvCodecIDS302m           = int(C.AV_CODEC_ID_S302M)
	AvCodecIDSami            = int(C.AV_CODEC_ID_SAMI)
	AvCodecIDSanm            = int(C.AV_CODEC_ID_SANM)
	AvCodecIDSgi             = int(C.AV_CODEC_ID_SGI)
	AvCodecIDSgirle          = int(C.AV_CODEC_ID_SGIRLE)
	AvCodecIDShorten         = int(C.AV_CODEC_ID_SHORTEN)
	AvCodecIDSipr            = int(C.AV_CODEC_ID_SIPR)
	AvCodecIDSmackaudio      = int(C.AV_CODEC_ID_SMACKAUDIO)
	AvCodecIDSmackvideo      = int(C.AV_CODEC_ID_SMACKVIDEO)
	AvCodecIDSmc             = int(C.AV_CODEC_ID_SMC)
	AvCodecIDSmpteKlv        = int(C.AV_CODEC_ID_SMPTE_KLV)
	AvCodecIDSmv             = int(C.AV_CODEC_ID_SMV)
	AvCodecIDSmvjpeg         = int(C.AV_CODEC_ID_SMVJPEG)
	AvCodecIDSnow            = int(C.AV_CODEC_ID_SNOW)
	AvCodecIDSolDpcm         = int(C.AV_CODEC_ID_SOL_DPCM)
	AvCodecIDSonic           = int(C.AV_CODEC_ID_SONIC)
	AvCodecIDSonicLs         = int(C.AV_CODEC_ID_SONIC_LS)
	AvCodecIDSp5x            = int(C.AV_CODEC_ID_SP5X)
	AvCodecIDSpeex           = int(C.AV_CODEC_ID_SPEEX)
	AvCodecIDSrt             = int(C.AV_CODEC_ID_SRT)
	AvCodecIDSsa             = int(C.AV_CODEC_ID_SSA)
	AvCodecIDSubrip          = int(C.AV_CODEC_ID_SUBRIP)
	AvCodecIDSubviewer       = int(C.AV_CODEC_ID_SUBVIEWER)
	AvCodecIDSubviewer1      = int(C.AV_CODEC_ID_SUBVIEWER1)
	AvCodecIDSunrast         = int(C.AV_CODEC_ID_SUNRAST)
	AvCodecIDSvq1            = int(C.AV_CODEC_ID_SVQ1)
	AvCodecIDSvq3            = int(C.AV_CODEC_ID_SVQ3)
	AvCodecIDTak             = int(C.AV_CODEC_ID_TAK)
	AvCodecIDTarga           = int(C.AV_CODEC_ID_TARGA)
	AvCodecIDTargaY216       = int(C.AV_CODEC_ID_TARGA_Y216)
	AvCodecIDText            = int(C.AV_CODEC_ID_TEXT)
	AvCodecIDTgq             = int(C.AV_CODEC_ID_TGQ)
	AvCodecIDTgv             = int(C.AV_CODEC_ID_TGV)
	AvCodecIDTheora          = int(C.AV_CODEC_ID_THEORA)
	AvCodecIDThp             = int(C.AV_CODEC_ID_THP)
	AvCodecIDTiertexseqvideo = int(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	AvCodecIDTiff            = int(C.AV_CODEC_ID_TIFF)
	AvCodecIDTimedID3        = int(C.AV_CODEC_ID_TIMED_ID3)
	AvCodecIDTmv             = int(C.AV_CODEC_ID_TMV)
	AvCodecIDTqi             = int(C.AV_CODEC_ID_TQI)
	AvCodecIDTruehd          = int(C.AV_CODEC_ID_TRUEHD)
	AvCodecIDTruemotion1     = int(C.AV_CODEC_ID_TRUEMOTION1)
	AvCodecIDTruemotion2     = int(C.AV_CODEC_ID_TRUEMOTION2)
	AvCodecIDTruespeech      = int(C.AV_CODEC_ID_TRUESPEECH)
	AvCodecIDTscc            = int(C.AV_CODEC_ID_TSCC)
	AvCodecIDTscc2           = int(C.AV_CODEC_ID_TSCC2)
	AvCodecIDTta             = int(C.AV_CODEC_ID_TTA)
	AvCodecIDTtf             = int(C.AV_CODEC_ID_TTF)
	AvCodecIDTwinvq          = int(C.AV_CODEC_ID_TWINVQ)
	AvCodecIDTxd             = int(C.AV_CODEC_ID_TXD)
	AvCodecIDUlti            = int(C.AV_CODEC_ID_ULTI)
	AvCodecIDUtvideo         = int(C.AV_CODEC_ID_UTVIDEO)
	AvCodecIDV210            = int(C.AV_CODEC_ID_V210)
	AvCodecIDV210x           = int(C.AV_CODEC_ID_V210X)
	AvCodecIDV308            = int(C.AV_CODEC_ID_V308)
	AvCodecIDV408            = int(C.AV_CODEC_ID_V408)
	AvCodecIDV410            = int(C.AV_CODEC_ID_V410)
	AvCodecIDVb              = int(C.AV_CODEC_ID_VB)
	AvCodecIDVble            = int(C.AV_CODEC_ID_VBLE)
	AvCodecIDVc1             = int(C.AV_CODEC_ID_VC1)
	AvCodecIDVc1image        = int(C.AV_CODEC_ID_VC1IMAGE)
	AvCodecIDVcr1            = int(C.AV_CODEC_ID_VCR1)
	// AV_CODEC_ID_VIMA             = int(C.AV_CODEC_ID_VIMA)
	AvCodecIDVixl     = int(C.AV_CODEC_ID_VIXL)
	AvCodecIDVmdaudio = int(C.AV_CODEC_ID_VMDAUDIO)
	AvCodecIDVmdvideo = int(C.AV_CODEC_ID_VMDVIDEO)
	AvCodecIDVmnc     = int(C.AV_CODEC_ID_VMNC)
	AvCodecIDVorbis   = int(C.AV_CODEC_ID_VORBIS)
	// AV_CODEC_ID_VOXWARE          = int(C.AV_CODEC_ID_VOXWARE)
	AvCodecIDVp3          = int(C.AV_CODEC_ID_VP3)
	AvCodecIDVp5          = int(C.AV_CODEC_ID_VP5)
	AvCodecIDVp6          = int(C.AV_CODEC_ID_VP6)
	AvCodecIDVp6a         = int(C.AV_CODEC_ID_VP6A)
	AvCodecIDVp6f         = int(C.AV_CODEC_ID_VP6F)
	AvCodecIDVp7          = int(C.AV_CODEC_ID_VP7)
	AvCodecIDVp8          = int(C.AV_CODEC_ID_VP8)
	AvCodecIDVp9          = int(C.AV_CODEC_ID_VP9)
	AvCodecIDVplayer      = int(C.AV_CODEC_ID_VPLAYER)
	AvCodecIDWavpack      = int(C.AV_CODEC_ID_WAVPACK)
	AvCodecIDWebp         = int(C.AV_CODEC_ID_WEBP)
	AvCodecIDWebvtt       = int(C.AV_CODEC_ID_WEBVTT)
	AvCodecIDWestwoodSnd1 = int(C.AV_CODEC_ID_WESTWOOD_SND1)
	AvCodecIDWmalossless  = int(C.AV_CODEC_ID_WMALOSSLESS)
	AvCodecIDWmapro       = int(C.AV_CODEC_ID_WMAPRO)
	AvCodecIDWmav1        = int(C.AV_CODEC_ID_WMAV1)
	AvCodecIDWmav2        = int(C.AV_CODEC_ID_WMAV2)
	AvCodecIDWmavoice     = int(C.AV_CODEC_ID_WMAVOICE)
	AvCodecIDWmv1         = int(C.AV_CODEC_ID_WMV1)
	AvCodecIDWmv2         = int(C.AV_CODEC_ID_WMV2)
	AvCodecIDWmv3         = int(C.AV_CODEC_ID_WMV3)
	AvCodecIDWmv3image    = int(C.AV_CODEC_ID_WMV3IMAGE)
	AvCodecIDWnv1         = int(C.AV_CODEC_ID_WNV1)
	AvCodecIDWsVqa        = int(C.AV_CODEC_ID_WS_VQA)
	AvCodecIDXanDpcm      = int(C.AV_CODEC_ID_XAN_DPCM)
	AvCodecIDXanWc3       = int(C.AV_CODEC_ID_XAN_WC3)
	AvCodecIDXanWc4       = int(C.AV_CODEC_ID_XAN_WC4)
	AvCodecIDXbin         = int(C.AV_CODEC_ID_XBIN)
	AvCodecIDXbm          = int(C.AV_CODEC_ID_XBM)
	AvCodecIDXface        = int(C.AV_CODEC_ID_XFACE)
	AvCodecIDXsub         = int(C.AV_CODEC_ID_XSUB)
	AvCodecIDXwd          = int(C.AV_CODEC_ID_XWD)
	AvCodecIDY41p         = int(C.AV_CODEC_ID_Y41P)
	AvCodecIDYop          = int(C.AV_CODEC_ID_YOP)
	AvCodecIDYuv4         = int(C.AV_CODEC_ID_YUV4)
	AvCodecIDZerocodec    = int(C.AV_CODEC_ID_ZEROCODEC)
	AvCodecIDZlib         = int(C.AV_CODEC_ID_ZLIB)
	AvCodecIDZmbv         = int(C.AV_CODEC_ID_ZMBV)
)
