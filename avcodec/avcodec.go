// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/frame.h>
import "C"
import (
	"unsafe"

	"github.com/asticode/goav/avutil"
)

type (
	Codec                         C.struct_AVCodec
	Context                       C.struct_AVCodecContext
	CodecParameters               C.struct_AVCodecParameters
	Descriptor                    C.struct_AVCodecDescriptor
	Parser                        C.struct_AVCodecParser
	ParserContext                 C.struct_AVCodecParserContext
	Frame                         C.struct_AVFrame
	MediaType                     C.enum_AVMediaType
	Packet                        C.struct_AVPacket
	BitStreamFilter               C.struct_AVBitStreamFilter
	BitStreamFilterContext        C.struct_AVBitStreamFilterContext
	Rational                      C.struct_AVRational
	Class                         C.struct_AVClass
	AvHWAccel                     C.struct_AVHWAccel
	AvPacketSideData              C.struct_AVPacketSideData
	AvPanScan                     C.struct_AVPanScan
	Picture                       C.struct_AVPicture
	AvProfile                     C.struct_AVProfile
	AvSubtitle                    C.struct_AVSubtitle
	AvSubtitleRect                C.struct_AVSubtitleRect
	RcOverride                    C.struct_RcOverride
	AvBufferRef                   C.struct_AVBufferRef
	AvAudioServiceType            C.enum_AVAudioServiceType
	AvChromaLocation              C.enum_AVChromaLocation
	CodecId                       C.enum_AVCodecID
	AvColorPrimaries              C.enum_AVColorPrimaries
	AvColorRange                  C.enum_AVColorRange
	AvColorSpace                  C.enum_AVColorSpace
	AvColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	AvDiscard                     C.enum_AVDiscard
	AvFieldOrder                  C.enum_AVFieldOrder
	AvPacketSideDataType          C.enum_AVPacketSideDataType
	AvSampleFormat                C.enum_AVSampleFormat
)

const (
	AvPktFlagKey     = C.AV_PKT_FLAG_KEY
	AvPktFlagCorrupt = C.AV_PKT_FLAG_CORRUPT
	AvPktFlagDiscard = C.AV_PKT_FLAG_DISCARD
)

const (
	AvDiscardNone     = C.AVDISCARD_NONE     // discard nothing
	AvDiscardDefault  = C.AVDISCARD_DEFAULT  // discard useless packets like 0 size packets in avi
	AvDiscardNonref   = C.AVDISCARD_NONREF   // discard all non reference
	AvDiscardBidir    = C.AVDISCARD_BIDIR    // discard all bidirectional frames
	AvDiscardNonintra = C.AVDISCARD_NONINTRA // discard all non intra frames
	AvDiscardNonkey   = C.AVDISCARD_NONKEY   // discard all frames except keyframes
	AvDiscardAll      = C.AVDISCARD_ALL      // discard all
)

const (
	AvMediaTypeAttachment = C.AVMEDIA_TYPE_ATTACHMENT
	AvMediaTypeAudio      = C.AVMEDIA_TYPE_AUDIO
	AvMediaTypeData       = C.AVMEDIA_TYPE_DATA
	AvMediaTypeNb         = C.AVMEDIA_TYPE_NB
	AvMediaTypeSubtitle   = C.AVMEDIA_TYPE_SUBTITLE
	AvMediaTypeUnknown    = C.AVMEDIA_TYPE_UNKNOWN
	AvMediaTypeVideo      = C.AVMEDIA_TYPE_VIDEO
)

const (
	AvPktDataPalette                  = C.AV_PKT_DATA_PALETTE
	AvPktDataNewExtradata             = C.AV_PKT_DATA_NEW_EXTRADATA
	AvPktDataParamChange              = C.AV_PKT_DATA_PARAM_CHANGE
	AvPktDataH263MbInfo               = C.AV_PKT_DATA_H263_MB_INFO
	AvPktDataReplaygain               = C.AV_PKT_DATA_REPLAYGAIN
	AvPktDataDisplaymatrix            = C.AV_PKT_DATA_DISPLAYMATRIX
	AvPktDataStereo3D                 = C.AV_PKT_DATA_STEREO3D
	AvPktDataAudioServiceType         = C.AV_PKT_DATA_AUDIO_SERVICE_TYPE
	AvPktDataSkipSamples              = C.AV_PKT_DATA_SKIP_SAMPLES
	AvPktDataJpDualmono               = C.AV_PKT_DATA_JP_DUALMONO
	AvPktDataStringsMetadata          = C.AV_PKT_DATA_STRINGS_METADATA
	AvPktDataSubtitlePosition         = C.AV_PKT_DATA_SUBTITLE_POSITION
	AvPktDataMatroskaBlockadditional  = C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL
	AvPktDataWebvttIdentifier         = C.AV_PKT_DATA_WEBVTT_IDENTIFIER
	AvPktDataWebvttSettings           = C.AV_PKT_DATA_WEBVTT_SETTINGS
	AvPktDataMetadataUpdate           = C.AV_PKT_DATA_METADATA_UPDATE
	AvPktDataMpegtsStreamId           = C.AV_PKT_DATA_MPEGTS_STREAM_ID
	AvPktDataMasteringDisplayMetadata = C.AV_PKT_DATA_MASTERING_DISPLAY_METADATA
	AvPktDataContentLightLevel        = C.AV_PKT_DATA_CONTENT_LIGHT_LEVEL
	AvPktDataSpherical                = C.AV_PKT_DATA_SPHERICAL
	AvPktDataA53Cc                    = C.AV_PKT_DATA_A53_CC
)

// AvGetProfileName Return a name for the specified profile, if available.
func (c *Codec) AvGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

// AllocContext3 Allocate an Context and set its fields to default values.
func (c *Codec) AllocContext3() *Context {
	return (*Context)(unsafe.Pointer(C.avcodec_alloc_context3((*C.struct_AVCodec)(c))))
}

func (c *Codec) IsEncoder() int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c)))
}

func (c *Codec) IsDecoder() int {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(c)))
}

func (c *Codec) SupportedSamplerates() []int {
	r := make([]int, 0)
	if c.supported_samplerates == nil {
		return r
	}
	size := unsafe.Sizeof(*c.supported_samplerates)
	for i := 0; ; i++ {
		p := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(c.supported_samplerates)) + uintptr(i)*size))
		if p == 0 {
			break
		}
		r = append(r, int(p))
	}
	return r
}

func (c *Codec) SampleFmts() []AvSampleFormat {
	r := make([]AvSampleFormat, 0)
	if c.sample_fmts == nil {
		return r
	}
	size := unsafe.Sizeof(*c.sample_fmts)
	for i := 0; ; i++ {
		p := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(c.sample_fmts)) + uintptr(i)*size))
		if p == C.AV_SAMPLE_FMT_NONE {
			break
		}
		r = append(r, AvSampleFormat(p))
	}
	return r
}

func (c *Codec) ChannelLayouts() []uint64 {
	r := make([]uint64, 0)
	if c.channel_layouts == nil {
		return r
	}
	size := unsafe.Sizeof(*c.channel_layouts)
	for i := 0; ; i++ {
		p := *(*C.uint64_t)(unsafe.Pointer(uintptr(unsafe.Pointer(c.channel_layouts)) + uintptr(i)*size))
		if p == 0 {
			break
		}
		r = append(r, uint64(p))
	}
	return r
}

// AvFastPaddedMalloc Same behaviour av_fast_malloc but the buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
func AvFastPaddedMalloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// Version Return the LIBAvCODEC_VERSION_INT constant.
func Version() uint {
	return uint(C.avcodec_version())
}

// Configuration Return the libavcodec build-time configuration.
func Configuration() string {
	return C.GoString(C.avcodec_configuration())

}

// License Return the libavcodec license.
func License() string {
	return C.GoString(C.avcodec_license())
}

// GetClass Get the Class for Context.
func GetClass() *Class {
	return (*Class)(C.avcodec_get_class())
}

// GetSubtitleRectClass Get the Class for AvSubtitleRect.
func GetSubtitleRectClass() *Class {
	return (*Class)(C.avcodec_get_subtitle_rect_class())
}

// AvSubtitleFree Free all allocated data in the given subtitle struct.
func AvSubtitleFree(s *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// AvPacketPackDictionary Pack a dictionary for use in side_data.
func AvPacketPackDictionary(d *avutil.Dictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

// AvPacketUnpackDictionary Unpack a dictionary from side_data.
func AvPacketUnpackDictionary(d *uint8, s int, dt **avutil.Dictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

// FindDecoder Find a registered decoder with a matching codec ID.
func FindDecoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// FindDecoderByName Find a registered decoder with the specified name.
func FindDecoderByName(n string) *Codec {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Codec)(C.avcodec_find_decoder_by_name(cn))
}

// EnumToChromaPos Converts AvChromaLocation to swscale x/y chroma position.
func EnumToChromaPos(x, y *int, l AvChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

// ChromaPosToEnum Converts swscale x/y chroma position to AvChromaLocation.
func ChromaPosToEnum(x, y int) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

// FindEncoder Find a registered encoder with a matching codec ID.
func FindEncoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

// FindEncoderByName Find a registered encoder with the specified name.
func FindEncoderByName(c string) *Codec {
	cc := C.CString(c)
	defer C.free(unsafe.Pointer(cc))
	return (*Codec)(C.avcodec_find_encoder_by_name(cc))
}

func String(b string, bs int, ctxt *Context, e int) {
	cb := C.CString(b)
	defer C.free(unsafe.Pointer(cb))
	C.avcodec_string(cb, C.int(bs), (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), C.int(e))
}

// FillAudioFrame Fill Frame audio data and linesize pointers.
func FillAudioFrame(f *Frame, c int, s AvSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

// AvGetBitsPerSample Return codec bits per sample.
func AvGetBitsPerSample(c CodecId) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

// AvGetPcmCodec Return the PCM codec associated with a sample format.
func AvGetPcmCodec(f AvSampleFormat, b int) CodecId {
	return (CodecId)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

// AvGetExactBitsPerSample Return codec bits per sample.
func AvGetExactBitsPerSample(c CodecId) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

// AvFastPaddedMallocz Same behaviour av_fast_padded_malloc except that buffer will always be 0-initialized after call.
func AvFastPaddedMallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// AvXiphlacing Encode extradata length to a buffer.
func AvXiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

// GetType Get the type of the given codec.
func GetType(c CodecId) MediaType {
	return (MediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

// GetName Get the name of a codec.
func GetName(d CodecId) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

// DescriptorGet const Descriptor *avcodec_descriptor_get (enum CodecId id)
func DescriptorGet(id CodecId) *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

func (d *Descriptor) Name() string {
	return C.GoString(d.name)
}

// AvcodecDescriptorNext Iterate over all codec descriptors known to libavcodec.
func (d *Descriptor) AvcodecDescriptorNext() *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

func DescriptorGetByName(n string) *Descriptor {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Descriptor)(C.avcodec_descriptor_get_by_name(cn))
}

func ReceivePacket(avctx *Context, avpkt *Packet) int {
	return int(C.avcodec_receive_packet((*C.struct_AVCodecContext)(avctx), (*C.struct_AVPacket)(avpkt)))
}

// SendPacket AvcodecSendPacket send packet data as input to a decoder.
// See the ffmpeg documentation: https://www.ffmpeg.org/doxygen/trunk/group__lavc__encdec.html
func SendPacket(avctx *Context, avpkt *Packet) int {
	return int(C.avcodec_send_packet((*C.struct_AVCodecContext)(avctx), (*C.struct_AVPacket)(avpkt)))
}

// ReceiveFrame AvcodecReceiveFrame receives a frame from the codec
// See the ffmpeg documentation: https://www.ffmpeg.org/doxygen/trunk/group__lavc__encdec.html
func ReceiveFrame(avctx *Context, frame *avutil.Frame) int {
	return int(C.avcodec_receive_frame((*C.struct_AVCodecContext)(avctx), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

func SendFrame(avctx *Context, frame *avutil.Frame) int {
	return int(C.avcodec_send_frame((*C.struct_AVCodecContext)(avctx), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}
