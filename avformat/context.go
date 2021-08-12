// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/asticode/goav/avcodec"
	"github.com/asticode/goav/avutil"
)

// CloseInput Close an opened input Context.
func CloseInput(ctxt *Context) {
	var ptr *C.struct_AVFormatContext = (*C.struct_AVFormatContext)(unsafe.Pointer(ctxt))
	C.avformat_close_input((**C.struct_AVFormatContext)(&ptr))
}

// InjectGlobalSideData This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func (ctxt *Context) InjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(ctxt))
}

// AvFmtCtxGetDurationEstimationMethod Returns the method used to set ctx->duration.
func (ctxt *Context) AvFmtCtxGetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(ctxt)))
}

// FreeContext Free an Context and all its streams.
func (ctxt *Context) FreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(ctxt))
}

// NewStream Add a new stream to a media file.
func (ctxt *Context) NewStream(c *AvCodec) *Stream {
	return (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(ctxt), (*C.struct_AVCodec)(c)))
}

func (ctxt *Context) AvNewProgram(id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(ctxt), C.int(id)))
}

// FindStreamInfo Read packets of a media file to get stream information.
func (ctxt *Context) FindStreamInfo(d **avutil.Dictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ctxt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// AvFindProgramFromStream Find the programs which belong to a given stream.
func (ctxt *Context) AvFindProgramFromStream(l *AvProgram, su int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(ctxt), (*C.struct_AVProgram)(l), C.int(su)))
}

// AvFindBestStream Find the "best" stream in the file.
func AvFindBestStream(ic *Context, t MediaType, ws, rs int, c **AvCodec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

// AvReadFrame Return the next frame of a stream.
func (ctxt *Context) AvReadFrame(pkt *avcodec.Packet) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(ctxt)), (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

// AvSeekFrame Seek to the keyframe at timestamp.
func (ctxt *Context) AvSeekFrame(st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(ctxt), C.int(st), C.int64_t(t), C.int(f)))
}

// SeekFile Seek to timestamp ts.
func (ctxt *Context) SeekFile(si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(ctxt), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

// Flush Discard all internally buffered data.
func (ctxt *Context) Flush() int {
	return int(C.avformat_flush((*C.struct_AVFormatContext)(ctxt)))
}

// AvReadPlay Start playing a network-based stream (e.g.
func (ctxt *Context) AvReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(ctxt)))
}

// AvReadPause Pause a network-based stream (e.g.
func (ctxt *Context) AvReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(ctxt)))
}

// WriteHeader Allocate the stream private data and write the stream header to an output media file.
func (ctxt *Context) WriteHeader(o **avutil.Dictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(ctxt), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

// AvWriteFrame Write a packet to an output media file.
func (ctxt *Context) AvWriteFrame(pkt *Packet) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(ctxt), (*C.struct_AVPacket)(pkt)))
}

// AvInterleavedWriteFrame Write a packet to an output media file ensuring correct interleaving.
func (ctxt *Context) AvInterleavedWriteFrame(pkt *Packet) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(ctxt), (*C.struct_AVPacket)(pkt)))
}

// AvWriteUncodedFrame Write a uncoded frame to an output media file.
func (ctxt *Context) AvWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(ctxt), C.int(si), (*C.struct_AVFrame)(f)))
}

// AvInterleavedWriteUncodedFrame Write a uncoded frame to an output media file.
func (ctxt *Context) AvInterleavedWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(ctxt), C.int(si), (*C.struct_AVFrame)(f)))
}

// AvWriteUncodedFrameQuery Test whether a muxer supports uncoded frame.
func (ctxt *Context) AvWriteUncodedFrameQuery(si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(ctxt), C.int(si)))
}

// AvWriteTrailer Write the stream trailer to an output media file and free the file private data.
func (ctxt *Context) AvWriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(ctxt)))
}

// AvGetOutputTimestamp Get timing information for the data currently output.
func (ctxt *Context) AvGetOutputTimestamp(st int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(ctxt), C.int(st), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

func (ctxt *Context) AvFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(ctxt)))
}

// AvDumpFormat Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (ctxt *Context) AvDumpFormat(i int, url string, io int) {
	cu := C.CString(url)
	defer C.free(unsafe.Pointer(cu))
	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(ctxt)), C.int(i), cu, C.int(io))
}

// AvGuessSampleAspectRatio Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func (ctxt *Context) AvGuessSampleAspectRatio(st *Stream, fr *Frame) avutil.Rational {
	r := (C.struct_AVRational)(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(ctxt), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))

	return *(*avutil.Rational)(unsafe.Pointer(&r))
}

// AvGuessFrameRate Guess the frame rate, based on both the container and codec information.
func (ctxt *Context) AvGuessFrameRate(st *Stream, fr *Frame) avutil.Rational {
	r := (C.struct_AVRational)(C.av_guess_frame_rate((*C.struct_AVFormatContext)(ctxt), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))

	return *(*avutil.Rational)(unsafe.Pointer(&r))
}

// MatchStreamSpecifier Check if the stream st contained in s is matched by the stream specifier spec.
func (ctxt *Context) MatchStreamSpecifier(st *Stream, spec string) int {
	cs := C.CString(spec)
	defer C.free(unsafe.Pointer(cs))
	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(ctxt), (*C.struct_AVStream)(st), cs))
}

func (ctxt *Context) QueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(ctxt)))
}
