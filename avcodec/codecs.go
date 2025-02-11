// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

// Multiple encoders have the same ID and are able to encode compatible streams.
const (
	CodecId012V                = int(C.AV_CODEC_ID_012V)
	CodecId4XM                 = int(C.AV_CODEC_ID_4XM)
	CodecId8BPS                = int(C.AV_CODEC_ID_8BPS)
	CodecId8SVXExp             = int(C.AV_CODEC_ID_8SVX_EXP)
	CodecId8SVXFib             = int(C.AV_CODEC_ID_8SVX_FIB)
	CodecIdA64Multi            = int(C.AV_CODEC_ID_A64_MULTI)
	CodecIdA64Multi5           = int(C.AV_CODEC_ID_A64_MULTI5)
	CodecIdAac                 = int(C.AV_CODEC_ID_AAC)
	CodecIdAacLatm             = int(C.AV_CODEC_ID_AAC_LATM)
	CodecIdAasc                = int(C.AV_CODEC_ID_AASC)
	CodecIdAc3                 = int(C.AV_CODEC_ID_AC3)
	CodecIdAdpcm4XM            = int(C.AV_CODEC_ID_ADPCM_4XM)
	CodecIdAdpcmAdx            = int(C.AV_CODEC_ID_ADPCM_ADX)
	CodecIdAdpcmAfc            = int(C.AV_CODEC_ID_ADPCM_AFC)
	CodecIdAdpcmCt             = int(C.AV_CODEC_ID_ADPCM_CT)
	CodecIdAdpcmDtk            = int(C.AV_CODEC_ID_ADPCM_DTK)
	CodecIdAdpcmEa             = int(C.AV_CODEC_ID_ADPCM_EA)
	CodecIdAdpcmEaMaxisXa      = int(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	CodecIdAdpcmEaR1           = int(C.AV_CODEC_ID_ADPCM_EA_R1)
	CodecIdAdpcmEaR2           = int(C.AV_CODEC_ID_ADPCM_EA_R2)
	CodecIdAdpcmEaR3           = int(C.AV_CODEC_ID_ADPCM_EA_R3)
	CodecIdAdpcmEaXas          = int(C.AV_CODEC_ID_ADPCM_EA_XAS)
	CodecIdAdpcmG722           = int(C.AV_CODEC_ID_ADPCM_G722)
	CodecIdAdpcmG726           = int(C.AV_CODEC_ID_ADPCM_G726)
	CodecIdAdpcmG726LE         = int(C.AV_CODEC_ID_ADPCM_G726LE)
	CodecIdAdpcmImaAmv         = int(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	CodecIdAdpcmImaApc         = int(C.AV_CODEC_ID_ADPCM_IMA_APC)
	CodecIdAdpcmImaDk3         = int(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	CodecIdAdpcmImaDk4         = int(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	CodecIdAdpcmImaEaEacs      = int(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	CodecIdAdpcmImaEaSead      = int(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	CodecIdAdpcmImaIss         = int(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	CodecIdAdpcmImaOki         = int(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	CodecIdAdpcmImaQt          = int(C.AV_CODEC_ID_ADPCM_IMA_QT)
	CodecIdAdpcmImaRad         = int(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	CodecIdAdpcmImaSmjpeg      = int(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	CodecIdAdpcmImaWav         = int(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	CodecIdAdpcmImaWs          = int(C.AV_CODEC_ID_ADPCM_IMA_WS)
	CodecIdAdpcmMs             = int(C.AV_CODEC_ID_ADPCM_MS)
	CodecIdAdpcmSbpro2         = int(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	CodecIdAdpcmSbpro3         = int(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	CodecIdAdpcmSbpro4         = int(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	CodecIdAdpcmSwf            = int(C.AV_CODEC_ID_ADPCM_SWF)
	CodecIdAdpcmThp            = int(C.AV_CODEC_ID_ADPCM_THP)
	CodecIdAdpcmVima           = int(C.AV_CODEC_ID_ADPCM_VIMA) // Deprecated
	CodecIdAdpcmXa             = int(C.AV_CODEC_ID_ADPCM_XA)
	CodecIdAdpcmYamaha         = int(C.AV_CODEC_ID_ADPCM_YAMAHA)
	CodecIdAic                 = int(C.AV_CODEC_ID_AIC)
	CodecIdAlac                = int(C.AV_CODEC_ID_ALAC)
	CodecIdAliasPix            = int(C.AV_CODEC_ID_ALIAS_PIX)
	CodecIdAmrNb               = int(C.AV_CODEC_ID_AMR_NB)
	CodecIdAmrWb               = int(C.AV_CODEC_ID_AMR_WB)
	CodecIdAmv                 = int(C.AV_CODEC_ID_AMV)
	CodecIdAnm                 = int(C.AV_CODEC_ID_ANM)
	CodecIdAnsi                = int(C.AV_CODEC_ID_ANSI)
	CodecIdApe                 = int(C.AV_CODEC_ID_APE)
	CodecIdAss                 = int(C.AV_CODEC_ID_ASS)
	CodecIdAsv1                = int(C.AV_CODEC_ID_ASV1)
	CodecIdAsv2                = int(C.AV_CODEC_ID_ASV2)
	CodecIdAtrac1              = int(C.AV_CODEC_ID_ATRAC1)
	CodecIdAtrac3              = int(C.AV_CODEC_ID_ATRAC3)
	CodecIdAtrac3P             = int(C.AV_CODEC_ID_ATRAC3P)
	CodecIdAura                = int(C.AV_CODEC_ID_AURA)
	CodecIdAura2               = int(C.AV_CODEC_ID_AURA2)
	CodecIdAvrn                = int(C.AV_CODEC_ID_AVRN)
	CodecIdAvrp                = int(C.AV_CODEC_ID_AVRP)
	CodecIdAvs                 = int(C.AV_CODEC_ID_AVS)
	CodecIdAvui                = int(C.AV_CODEC_ID_AVUI)
	CodecIdAyuv                = int(C.AV_CODEC_ID_AYUV)
	CodecIdBethsoftvid         = int(C.AV_CODEC_ID_BETHSOFTVID)
	CodecIdBfi                 = int(C.AV_CODEC_ID_BFI)
	CodecIdBinkaudioDct        = int(C.AV_CODEC_ID_BINKAUDIO_DCT)
	CodecIdBinkaudioRdft       = int(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	CodecIdBinkvideo           = int(C.AV_CODEC_ID_BINKVIDEO)
	CodecIdBintext             = int(C.AV_CODEC_ID_BINTEXT)
	CodecIdBinData             = int(C.AV_CODEC_ID_BIN_DATA)
	CodecIdBmp                 = int(C.AV_CODEC_ID_BMP)
	CodecIdBmvAudio            = int(C.AV_CODEC_ID_BMV_AUDIO)
	CodecIdBmvVideo            = int(C.AV_CODEC_ID_BMV_VIDEO)
	CodecIdBrenderPix          = int(C.AV_CODEC_ID_BRENDER_PIX) // Deprecated
	CodecIdC93                 = int(C.AV_CODEC_ID_C93)
	CodecIdCavs                = int(C.AV_CODEC_ID_CAVS)
	CodecIdCdgraphics          = int(C.AV_CODEC_ID_CDGRAPHICS)
	CodecIdCdxl                = int(C.AV_CODEC_ID_CDXL)
	CodecIdCelt                = int(C.AV_CODEC_ID_CELT)
	CodecIdCinepak             = int(C.AV_CODEC_ID_CINEPAK)
	CodecIdCljr                = int(C.AV_CODEC_ID_CLJR)
	CodecIdCllc                = int(C.AV_CODEC_ID_CLLC)
	CodecIdCmv                 = int(C.AV_CODEC_ID_CMV)
	CodecIdComfortNoise        = int(C.AV_CODEC_ID_COMFORT_NOISE)
	CodecIdCook                = int(C.AV_CODEC_ID_COOK)
	CodecIdCpia                = int(C.AV_CODEC_ID_CPIA)
	CodecIdCscd                = int(C.AV_CODEC_ID_CSCD)
	CodecIdCyuv                = int(C.AV_CODEC_ID_CYUV)
	CodecIdDfa                 = int(C.AV_CODEC_ID_DFA)
	CodecIdDirac               = int(C.AV_CODEC_ID_DIRAC)
	CodecIdDnxhd               = int(C.AV_CODEC_ID_DNXHD)
	CodecIdDpx                 = int(C.AV_CODEC_ID_DPX)
	CodecIdDsdLsbf             = int(C.AV_CODEC_ID_DSD_LSBF)
	CodecIdDsdLsbfPlanar       = int(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	CodecIdDsdMsbf             = int(C.AV_CODEC_ID_DSD_MSBF)
	CodecIdDsdMsbfPlanar       = int(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	CodecIdDsicinaudio         = int(C.AV_CODEC_ID_DSICINAUDIO)
	CodecIdDsicinvideo         = int(C.AV_CODEC_ID_DSICINVIDEO)
	CodecIdDts                 = int(C.AV_CODEC_ID_DTS)
	CodecIdDvaudio             = int(C.AV_CODEC_ID_DVAUDIO)
	CodecIdDvbSubtitle         = int(C.AV_CODEC_ID_DVB_SUBTITLE)
	CodecIdDvbTeletext         = int(C.AV_CODEC_ID_DVB_TELETEXT)
	CodecIdDvdNav              = int(C.AV_CODEC_ID_DVD_NAV)
	CodecIdDvdSubtitle         = int(C.AV_CODEC_ID_DVD_SUBTITLE)
	CodecIdDvvideo             = int(C.AV_CODEC_ID_DVVIDEO)
	CodecIdDxa                 = int(C.AV_CODEC_ID_DXA)
	CodecIdDxtory              = int(C.AV_CODEC_ID_DXTORY)
	CodecIdEac3                = int(C.AV_CODEC_ID_EAC3)
	CodecIdEia608              = int(C.AV_CODEC_ID_EIA_608)
	CodecIdEscape124           = int(C.AV_CODEC_ID_ESCAPE124)
	CodecIdEscape130           = int(C.AV_CODEC_ID_ESCAPE130)
	CodecIdEscape130Deprecated = int(C.AV_CODEC_ID_ESCAPE130)
	CodecIdEvrc                = int(C.AV_CODEC_ID_EVRC)
	CodecIdExr                 = int(C.AV_CODEC_ID_EXR)
	CodecIdExrDeprecated       = int(C.AV_CODEC_ID_EXR)
	CodecIdFfmetadata          = int(C.AV_CODEC_ID_FFMETADATA)
	CodecIdFfv1                = int(C.AV_CODEC_ID_FFV1)
	CodecIdFfvhuff             = int(C.AV_CODEC_ID_FFVHUFF)
	CodecIdFfwavesynth         = int(C.AV_CODEC_ID_FFWAVESYNTH)
	CodecIdFic                 = int(C.AV_CODEC_ID_FIC)
	CodecIdFirstAudio          = int(C.AV_CODEC_ID_FIRST_AUDIO)
	CodecIdFirstSubtitle       = int(C.AV_CODEC_ID_FIRST_SUBTITLE)
	CodecIdFirstUnknown        = int(C.AV_CODEC_ID_FIRST_UNKNOWN)
	CodecIdFlac                = int(C.AV_CODEC_ID_FLAC)
	CodecIdFlashsv             = int(C.AV_CODEC_ID_FLASHSV)
	CodecIdFlashsv2            = int(C.AV_CODEC_ID_FLASHSV2)
	CodecIdFlic                = int(C.AV_CODEC_ID_FLIC)
	CodecIdFlv1                = int(C.AV_CODEC_ID_FLV1)
	CodecIdFraps               = int(C.AV_CODEC_ID_FRAPS)
	CodecIdFrwu                = int(C.AV_CODEC_ID_FRWU)
	CodecIdG2M                 = int(C.AV_CODEC_ID_G2M) // Deprecated
	CodecIdG7231               = int(C.AV_CODEC_ID_G723_1)
	CodecIdG729                = int(C.AV_CODEC_ID_G729)
	CodecIdGif                 = int(C.AV_CODEC_ID_GIF)
	CodecIdGsm                 = int(C.AV_CODEC_ID_GSM)
	CodecIdGsmMs               = int(C.AV_CODEC_ID_GSM_MS)
	CodecIdH261                = int(C.AV_CODEC_ID_H261)
	CodecIdH263                = int(C.AV_CODEC_ID_H263)
	CodecIdH263I               = int(C.AV_CODEC_ID_H263I)
	CodecIdH263P               = int(C.AV_CODEC_ID_H263P)
	CodecIdH264                = int(C.AV_CODEC_ID_H264)
	CodecIdHdmvPgsSubtitle     = int(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	CodecIdHevc                = int(C.AV_CODEC_ID_HEVC)
	CodecIdHevcDeprecated      = int(C.AV_CODEC_ID_HEVC)
	CodecIdHnm4Video           = int(C.AV_CODEC_ID_HNM4_VIDEO)
	CodecIdHuffyuv             = int(C.AV_CODEC_ID_HUFFYUV)
	CodecIdIac                 = int(C.AV_CODEC_ID_IAC)
	CodecIdIdcin               = int(C.AV_CODEC_ID_IDCIN)
	CodecIdIdf                 = int(C.AV_CODEC_ID_IDF)
	CodecIdIffByterun1         = int(C.AV_CODEC_ID_IFF_BYTERUN1)
	CodecIdIffIlbm             = int(C.AV_CODEC_ID_IFF_ILBM)
	CodecIdIlbc                = int(C.AV_CODEC_ID_ILBC)
	CodecIdImc                 = int(C.AV_CODEC_ID_IMC)
	CodecIdIndeo2              = int(C.AV_CODEC_ID_INDEO2)
	CodecIdIndeo3              = int(C.AV_CODEC_ID_INDEO3)
	CodecIdIndeo4              = int(C.AV_CODEC_ID_INDEO4)
	CodecIdIndeo5              = int(C.AV_CODEC_ID_INDEO5)
	CodecIdInterplayDpcm       = int(C.AV_CODEC_ID_INTERPLAY_DPCM)
	CodecIdInterplayVideo      = int(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	CodecIdJacosub             = int(C.AV_CODEC_ID_JACOSUB)
	CodecIdJpeg2000            = int(C.AV_CODEC_ID_JPEG2000)
	CodecIdJpegls              = int(C.AV_CODEC_ID_JPEGLS)
	CodecIdJv                  = int(C.AV_CODEC_ID_JV)
	CodecIdKgv1                = int(C.AV_CODEC_ID_KGV1)
	CodecIdKmvc                = int(C.AV_CODEC_ID_KMVC)
	CodecIdLagarith            = int(C.AV_CODEC_ID_LAGARITH)
	CodecIdLjpeg               = int(C.AV_CODEC_ID_LJPEG)
	CodecIdLoco                = int(C.AV_CODEC_ID_LOCO)
	CodecIdMace3               = int(C.AV_CODEC_ID_MACE3)
	CodecIdMace6               = int(C.AV_CODEC_ID_MACE6)
	CodecIdMad                 = int(C.AV_CODEC_ID_MAD)
	CodecIdMdec                = int(C.AV_CODEC_ID_MDEC)
	CodecIdMetasound           = int(C.AV_CODEC_ID_METASOUND)
	CodecIdMicrodvd            = int(C.AV_CODEC_ID_MICRODVD)
	CodecIdMimic               = int(C.AV_CODEC_ID_MIMIC)
	CodecIdMjpeg               = int(C.AV_CODEC_ID_MJPEG)
	CodecIdMjpegb              = int(C.AV_CODEC_ID_MJPEGB)
	CodecIdMlp                 = int(C.AV_CODEC_ID_MLP)
	CodecIdMmvideo             = int(C.AV_CODEC_ID_MMVIDEO)
	CodecIdMotionpixels        = int(C.AV_CODEC_ID_MOTIONPIXELS)
	CodecIdMovText             = int(C.AV_CODEC_ID_MOV_TEXT)
	CodecIdMp1                 = int(C.AV_CODEC_ID_MP1)
	CodecIdMp2                 = int(C.AV_CODEC_ID_MP2)
	CodecIdMp3                 = int(C.AV_CODEC_ID_MP3)
	CodecIdMp3ADU              = int(C.AV_CODEC_ID_MP3ADU)
	CodecIdMp3ON4              = int(C.AV_CODEC_ID_MP3ON4)
	CodecIdMp4ALS              = int(C.AV_CODEC_ID_MP4ALS)
	CodecIdMpeg1VIDEO          = int(C.AV_CODEC_ID_MPEG1VIDEO)
	CodecIdMpeg2TS             = int(C.AV_CODEC_ID_MPEG2TS)
	CodecIdMpeg2VIDEO          = int(C.AV_CODEC_ID_MPEG2VIDEO)
	CodecIdMpeg4               = int(C.AV_CODEC_ID_MPEG4)
	CodecIdMpeg4SYSTEMS        = int(C.AV_CODEC_ID_MPEG4SYSTEMS)
	CodecIdMpl2                = int(C.AV_CODEC_ID_MPL2)
	CodecIdMsa1                = int(C.AV_CODEC_ID_MSA1)
	CodecIdMsmpeg4V1           = int(C.AV_CODEC_ID_MSMPEG4V1)
	CodecIdMsmpeg4V2           = int(C.AV_CODEC_ID_MSMPEG4V2)
	CodecIdMsmpeg4V3           = int(C.AV_CODEC_ID_MSMPEG4V3)
	CodecIdMsrle               = int(C.AV_CODEC_ID_MSRLE)
	CodecIdMss1                = int(C.AV_CODEC_ID_MSS1)
	CodecIdMss2                = int(C.AV_CODEC_ID_MSS2)
	CodecIdMsvideo1            = int(C.AV_CODEC_ID_MSVIDEO1)
	CodecIdMszh                = int(C.AV_CODEC_ID_MSZH)
	CodecIdMts2                = int(C.AV_CODEC_ID_MTS2)
	CodecIdMusepack7           = int(C.AV_CODEC_ID_MUSEPACK7)
	CodecIdMusepack8           = int(C.AV_CODEC_ID_MUSEPACK8)
	CodecIdMvc1                = int(C.AV_CODEC_ID_MVC1) // Deprecated
	CodecIdMvc2                = int(C.AV_CODEC_ID_MVC2) // Deprecated
	CodecIdMxpeg               = int(C.AV_CODEC_ID_MXPEG)
	CodecIdNellymoser          = int(C.AV_CODEC_ID_NELLYMOSER)
	CodecIdNone                = int(C.AV_CODEC_ID_NONE)
	CodecIdNuv                 = int(C.AV_CODEC_ID_NUV)
	CodecIdOn2AVC              = int(C.AV_CODEC_ID_ON2AVC)
	CodecIdOpus                = int(C.AV_CODEC_ID_OPUS) // Deprecated
	CodecIdOtf                 = int(C.AV_CODEC_ID_OTF)
	CodecIdPafAudio            = int(C.AV_CODEC_ID_PAF_AUDIO) // Deprecated
	CodecIdPafVideo            = int(C.AV_CODEC_ID_PAF_VIDEO) // Deprecated
	CodecIdPam                 = int(C.AV_CODEC_ID_PAM)
	CodecIdPbm                 = int(C.AV_CODEC_ID_PBM)
	CodecIdPcmAlaw             = int(C.AV_CODEC_ID_PCM_ALAW)
	CodecIdPcmBluray           = int(C.AV_CODEC_ID_PCM_BLURAY)
	CodecIdPcmDvd              = int(C.AV_CODEC_ID_PCM_DVD)
	CodecIdPcmF32BE            = int(C.AV_CODEC_ID_PCM_F32BE)
	CodecIdPcmF32LE            = int(C.AV_CODEC_ID_PCM_F32LE)
	CodecIdPcmF64BE            = int(C.AV_CODEC_ID_PCM_F64BE)
	CodecIdPcmF64LE            = int(C.AV_CODEC_ID_PCM_F64LE)
	CodecIdPcmLxf              = int(C.AV_CODEC_ID_PCM_LXF)
	CodecIdPcmMulaw            = int(C.AV_CODEC_ID_PCM_MULAW)
	CodecIdPcmS16BE            = int(C.AV_CODEC_ID_PCM_S16BE)
	CodecIdPcmS16BEPlanar      = int(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	CodecIdPcmS16LE            = int(C.AV_CODEC_ID_PCM_S16LE)
	CodecIdPcmS16LEPlanar      = int(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	CodecIdPcmS24BE            = int(C.AV_CODEC_ID_PCM_S24BE)
	CodecIdPcmS24DAUD          = int(C.AV_CODEC_ID_PCM_S24DAUD)
	CodecIdPcmS24LE            = int(C.AV_CODEC_ID_PCM_S24LE)
	CodecIdPcmS24LEPlanar      = int(C.AV_CODEC_ID_PCM_S24LE_PLANAR) // Deprecated
	CodecIdPcmS32BE            = int(C.AV_CODEC_ID_PCM_S32BE)
	CodecIdPcmS32LE            = int(C.AV_CODEC_ID_PCM_S32LE)
	CodecIdPcmS32LEPlanar      = int(C.AV_CODEC_ID_PCM_S32LE_PLANAR) // Deprecated
	CodecIdPcmS8               = int(C.AV_CODEC_ID_PCM_S8)
	CodecIdPcmS8Planar         = int(C.AV_CODEC_ID_PCM_S8_PLANAR)
	CodecIdPcmU16BE            = int(C.AV_CODEC_ID_PCM_U16BE)
	CodecIdPcmU16LE            = int(C.AV_CODEC_ID_PCM_U16LE)
	CodecIdPcmU24BE            = int(C.AV_CODEC_ID_PCM_U24BE)
	CodecIdPcmU24LE            = int(C.AV_CODEC_ID_PCM_U24LE)
	CodecIdPcmU32BE            = int(C.AV_CODEC_ID_PCM_U32BE)
	CodecIdPcmU32LE            = int(C.AV_CODEC_ID_PCM_U32LE)
	CodecIdPcmU8               = int(C.AV_CODEC_ID_PCM_U8)
	CodecIdPcmZork             = int(C.AV_CODEC_ID_PCM_ZORK)
	CodecIdPcx                 = int(C.AV_CODEC_ID_PCX)
	CodecIdPgm                 = int(C.AV_CODEC_ID_PGM)
	CodecIdPgmyuv              = int(C.AV_CODEC_ID_PGMYUV)
	CodecIdPictor              = int(C.AV_CODEC_ID_PICTOR)
	CodecIdPjs                 = int(C.AV_CODEC_ID_PJS)
	CodecIdPng                 = int(C.AV_CODEC_ID_PNG)
	CodecIdPpm                 = int(C.AV_CODEC_ID_PPM)
	CodecIdProbe               = int(C.AV_CODEC_ID_PROBE)
	CodecIdProres              = int(C.AV_CODEC_ID_PRORES)
	CodecIdPtx                 = int(C.AV_CODEC_ID_PTX)
	CodecIdQcelp               = int(C.AV_CODEC_ID_QCELP)
	CodecIdQdm2                = int(C.AV_CODEC_ID_QDM2)
	CodecIdQdmc                = int(C.AV_CODEC_ID_QDMC)
	CodecIdQdraw               = int(C.AV_CODEC_ID_QDRAW)
	CodecIdQpeg                = int(C.AV_CODEC_ID_QPEG)
	CodecIdQtrle               = int(C.AV_CODEC_ID_QTRLE)
	CodecIdR10K                = int(C.AV_CODEC_ID_R10K)
	CodecIdR210                = int(C.AV_CODEC_ID_R210)
	CodecIdRalf                = int(C.AV_CODEC_ID_RALF)
	CodecIdRawvideo            = int(C.AV_CODEC_ID_RAWVIDEO)
	CodecIdRa144               = int(C.AV_CODEC_ID_RA_144)
	CodecIdRa288               = int(C.AV_CODEC_ID_RA_288)
	CodecIdRealtext            = int(C.AV_CODEC_ID_REALTEXT)
	CodecIdRl2                 = int(C.AV_CODEC_ID_RL2)
	CodecIdRoq                 = int(C.AV_CODEC_ID_ROQ)
	CodecIdRoqDpcm             = int(C.AV_CODEC_ID_ROQ_DPCM)
	CodecIdRpza                = int(C.AV_CODEC_ID_RPZA)
	CodecIdRv10                = int(C.AV_CODEC_ID_RV10)
	CodecIdRv20                = int(C.AV_CODEC_ID_RV20)
	CodecIdRv30                = int(C.AV_CODEC_ID_RV30)
	CodecIdRv40                = int(C.AV_CODEC_ID_RV40)
	CodecIdS302M               = int(C.AV_CODEC_ID_S302M)
	CodecIdSami                = int(C.AV_CODEC_ID_SAMI)
	CodecIdSanm                = int(C.AV_CODEC_ID_SANM)
	CodecIdSanmDeprecated      = int(C.AV_CODEC_ID_SANM)
	CodecIdSgi                 = int(C.AV_CODEC_ID_SGI)
	CodecIdSgirle              = int(C.AV_CODEC_ID_SGIRLE)
	CodecIdSgirleDeprecated    = int(C.AV_CODEC_ID_SGIRLE)
	CodecIdShorten             = int(C.AV_CODEC_ID_SHORTEN)
	CodecIdSipr                = int(C.AV_CODEC_ID_SIPR)
	CodecIdSmackaudio          = int(C.AV_CODEC_ID_SMACKAUDIO)
	CodecIdSmackvideo          = int(C.AV_CODEC_ID_SMACKVIDEO)
	CodecIdSmc                 = int(C.AV_CODEC_ID_SMC)
	CodecIdSmpteKlv            = int(C.AV_CODEC_ID_SMPTE_KLV)
	CodecIdSmv                 = int(C.AV_CODEC_ID_SMV)
	CodecIdSmvjpeg             = int(C.AV_CODEC_ID_SMVJPEG)
	CodecIdSnow                = int(C.AV_CODEC_ID_SNOW)
	CodecIdSolDpcm             = int(C.AV_CODEC_ID_SOL_DPCM)
	CodecIdSonic               = int(C.AV_CODEC_ID_SONIC)
	CodecIdSonicLs             = int(C.AV_CODEC_ID_SONIC_LS)
	CodecIdSp5X                = int(C.AV_CODEC_ID_SP5X)
	CodecIdSpeex               = int(C.AV_CODEC_ID_SPEEX)
	CodecIdSrt                 = int(C.AV_CODEC_ID_SRT)
	CodecIdSsa                 = int(C.AV_CODEC_ID_SSA)
	CodecIdSubrip              = int(C.AV_CODEC_ID_SUBRIP)
	CodecIdSubviewer           = int(C.AV_CODEC_ID_SUBVIEWER)
	CodecIdSubviewer1          = int(C.AV_CODEC_ID_SUBVIEWER1)
	CodecIdSunrast             = int(C.AV_CODEC_ID_SUNRAST)
	CodecIdSvq1                = int(C.AV_CODEC_ID_SVQ1)
	CodecIdSvq3                = int(C.AV_CODEC_ID_SVQ3)
	CodecIdTak                 = int(C.AV_CODEC_ID_TAK)
	CodecIdTakDeprecated       = int(C.AV_CODEC_ID_TAK)
	CodecIdTarga               = int(C.AV_CODEC_ID_TARGA)
	CodecIdTargaY216           = int(C.AV_CODEC_ID_TARGA_Y216)
	CodecIdText                = int(C.AV_CODEC_ID_TEXT)
	CodecIdTgq                 = int(C.AV_CODEC_ID_TGQ)
	CodecIdTgv                 = int(C.AV_CODEC_ID_TGV)
	CodecIdTheora              = int(C.AV_CODEC_ID_THEORA)
	CodecIdThp                 = int(C.AV_CODEC_ID_THP)
	CodecIdTiertexseqvideo     = int(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	CodecIdTiff                = int(C.AV_CODEC_ID_TIFF)
	CodecIdTimedId3            = int(C.AV_CODEC_ID_TIMED_ID3)
	CodecIdTmv                 = int(C.AV_CODEC_ID_TMV)
	CodecIdTqi                 = int(C.AV_CODEC_ID_TQI)
	CodecIdTruehd              = int(C.AV_CODEC_ID_TRUEHD)
	CodecIdTruemotion1         = int(C.AV_CODEC_ID_TRUEMOTION1)
	CodecIdTruemotion2         = int(C.AV_CODEC_ID_TRUEMOTION2)
	CodecIdTruespeech          = int(C.AV_CODEC_ID_TRUESPEECH)
	CodecIdTscc                = int(C.AV_CODEC_ID_TSCC)
	CodecIdTscc2               = int(C.AV_CODEC_ID_TSCC2)
	CodecIdTta                 = int(C.AV_CODEC_ID_TTA)
	CodecIdTtf                 = int(C.AV_CODEC_ID_TTF)
	CodecIdTwinvq              = int(C.AV_CODEC_ID_TWINVQ)
	CodecIdTxd                 = int(C.AV_CODEC_ID_TXD)
	CodecIdUlti                = int(C.AV_CODEC_ID_ULTI)
	CodecIdUtvideo             = int(C.AV_CODEC_ID_UTVIDEO)
	CodecIdV210                = int(C.AV_CODEC_ID_V210)
	CodecIdV210X               = int(C.AV_CODEC_ID_V210X)
	CodecIdV308                = int(C.AV_CODEC_ID_V308)
	CodecIdV408                = int(C.AV_CODEC_ID_V408)
	CodecIdV410                = int(C.AV_CODEC_ID_V410)
	CodecIdVb                  = int(C.AV_CODEC_ID_VB)
	CodecIdVble                = int(C.AV_CODEC_ID_VBLE)
	CodecIdVc1                 = int(C.AV_CODEC_ID_VC1)
	CodecIdVc1IMAGE            = int(C.AV_CODEC_ID_VC1IMAGE)
	CodecIdVcr1                = int(C.AV_CODEC_ID_VCR1)
	CodecIdVixl                = int(C.AV_CODEC_ID_VIXL)
	CodecIdVmdaudio            = int(C.AV_CODEC_ID_VMDAUDIO)
	CodecIdVmdvideo            = int(C.AV_CODEC_ID_VMDVIDEO)
	CodecIdVmnc                = int(C.AV_CODEC_ID_VMNC)
	CodecIdVorbis              = int(C.AV_CODEC_ID_VORBIS)
	CodecIdVp3                 = int(C.AV_CODEC_ID_VP3)
	CodecIdVp5                 = int(C.AV_CODEC_ID_VP5)
	CodecIdVp6                 = int(C.AV_CODEC_ID_VP6)
	CodecIdVp6A                = int(C.AV_CODEC_ID_VP6A)
	CodecIdVp6F                = int(C.AV_CODEC_ID_VP6F)
	CodecIdVp7                 = int(C.AV_CODEC_ID_VP7) // Deprecated
	CodecIdVp8                 = int(C.AV_CODEC_ID_VP8)
	CodecIdVp9                 = int(C.AV_CODEC_ID_VP9)
	CodecIdVplayer             = int(C.AV_CODEC_ID_VPLAYER)
	CodecIdWavpack             = int(C.AV_CODEC_ID_WAVPACK)
	CodecIdWebp                = int(C.AV_CODEC_ID_WEBP) // Deprecated
	CodecIdWebvtt              = int(C.AV_CODEC_ID_WEBVTT)
	CodecIdWestwoodSnd1        = int(C.AV_CODEC_ID_WESTWOOD_SND1)
	CodecIdWmalossless         = int(C.AV_CODEC_ID_WMALOSSLESS)
	CodecIdWmapro              = int(C.AV_CODEC_ID_WMAPRO)
	CodecIdWmav1               = int(C.AV_CODEC_ID_WMAV1)
	CodecIdWmav2               = int(C.AV_CODEC_ID_WMAV2)
	CodecIdWmavoice            = int(C.AV_CODEC_ID_WMAVOICE)
	CodecIdWmv1                = int(C.AV_CODEC_ID_WMV1)
	CodecIdWmv2                = int(C.AV_CODEC_ID_WMV2)
	CodecIdWmv3                = int(C.AV_CODEC_ID_WMV3)
	CodecIdWmv3IMAGE           = int(C.AV_CODEC_ID_WMV3IMAGE)
	CodecIdWnv1                = int(C.AV_CODEC_ID_WNV1)
	CodecIdWsVqa               = int(C.AV_CODEC_ID_WS_VQA)
	CodecIdXanDpcm             = int(C.AV_CODEC_ID_XAN_DPCM)
	CodecIdXanWc3              = int(C.AV_CODEC_ID_XAN_WC3)
	CodecIdXanWc4              = int(C.AV_CODEC_ID_XAN_WC4)
	CodecIdXbin                = int(C.AV_CODEC_ID_XBIN)
	CodecIdXbm                 = int(C.AV_CODEC_ID_XBM)
	CodecIdXface               = int(C.AV_CODEC_ID_XFACE)
	CodecIdXsub                = int(C.AV_CODEC_ID_XSUB)
	CodecIdXwd                 = int(C.AV_CODEC_ID_XWD)
	CodecIdY41P                = int(C.AV_CODEC_ID_Y41P)
	CodecIdYop                 = int(C.AV_CODEC_ID_YOP)
	CodecIdYuv4                = int(C.AV_CODEC_ID_YUV4)
	CodecIdZerocodec           = int(C.AV_CODEC_ID_ZEROCODEC)
	CodecIdZlib                = int(C.AV_CODEC_ID_ZLIB)
	CodecIdZmbv                = int(C.AV_CODEC_ID_ZMBV)
)

// Flags
const (
	// CodecFlagGlobalHeader global header flag
	CodecFlagGlobalHeader = 1 << 22
)
