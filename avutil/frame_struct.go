package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
//#include <stdlib.h>
/*
static inline uint8_t ** dataItem(uint8_t * data, int idx)
{
	return (uint8_t **)&data[idx];
}
int lineSizeItem(AVFrame *frame, int idx)
{
	return frame->linesize[idx];
}

void set_data(AVFrame *frame, int idx, int l_size, uint8_t data) {
    frame->data[idx][l_size] = data;
}
*/
import "C"
import "unsafe"

func (f *Frame) Data() *uint8 {
	return (*uint8)(unsafe.Pointer((*C.uint8_t)(unsafe.Pointer(&f.data))))
}

func (f *Frame) DataItem(idx int) **uint8 {
	return (**uint8)(unsafe.Pointer(C.dataItem((*C.uint8_t)(unsafe.Pointer(&f.data)), C.int(idx))))
}

func (f *Frame) SetData(idx int, lineSize int, data byte) *Frame {
	C.set_data((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(idx), C.int(lineSize), (C.uint8_t)(data))
	return f
}

func (f *Frame) LineSize() int {
	return int(*(*C.int)(unsafe.Pointer(&f.linesize)))
}

func (f *Frame) LineSizeItem(idx int) int {
	return int(C.lineSizeItem((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(idx)))
}

func (f *Frame) LineSizePtr() *int {
	return (*int)(unsafe.Pointer((*C.int)(unsafe.Pointer(&f.linesize))))
}

func (f *Frame) Width() int {
	return int(f.width)
}

func (f *Frame) SetWidth(w int) *Frame {
	f.width = C.int(w)
	return f
}

func (f *Frame) Height() int {
	return int(f.height)
}

func (f *Frame) SetHeight(h int) *Frame {
	f.height = C.int(h)
	return f
}

func (f *Frame) Format() int {
	return int(f.format)
}

func (f *Frame) SetFormat(fmt PixelFormat) *Frame {
	f.format = C.int(fmt)
	return f
}

func (f *Frame) NbSamples() int {
	return int(f.nb_samples)
}

func (f *Frame) SetNbSamples(n int) *Frame {
	f.nb_samples = C.int(n)
	return f
}

func (f *Frame) SetKeyFrame(k int) *Frame {
	f.key_frame = C.int(k)
	return f
}

func (f *Frame) SetPictType(t AvPictureType) *Frame {
	f.pict_type = C.enum_AVPictureType(t)
	return f
}

func (f *Frame) Pts() int64 {
	return int64(f.pts)
}

func (f *Frame) SetPts(i int64) *Frame {
	f.pts = C.int64_t(i)
	return f
}

func (f *Frame) PktPts() int64 {
	return int64(f.pkt_pts)
}

func (f *Frame) PktDts() int64 {
	return int64(f.pkt_dts)
}

func (f *Frame) SampleRate() int {
	return int(f.sample_rate)
}

func (f *Frame) SetChannelLayout(l uint64) {
	f.channel_layout = C.uint64_t(l)
}

func (f *Frame) SetSampleRate(r int) {
	f.sample_rate = C.int(r)
}

// TODO Create getters and setters
// https://ffmpeg.org/doxygen/4.0/structAVFrame.html
/*
type Frame struct {
	Coded_picture_number   int32
	Data                   [8]*uint8
	Extended_data          **uint8
	Format                 int32
	Height                 int32
	Key_frame              int32
	Linesize               [8]int32
	Nb_samples             int32
	Pict_type              uint32
	Pts                    int64
	Pkt_pts                int64
	Pkt_dts                int64
	Sample_aspect_ratio    Rational
	Width                  int32
	Display_picture_number int32
	Quality                int32
	Pad_cgo_0              [4]byte
	Opaque                 *byte
	Error                  [8]uint64
	Repeat_pict            int32
	Interlaced_frame       int32
	Top_field_first        int32
	Palette_has_changed    int32
	Reordered_opaque       int64
	Sample_rate            int32
	Pad_cgo_1              [4]byte
	Channel_layout         uint64
	Buf                    [8]*AvBufferRef
	Extended_buf           **AvBufferRef
	Nb_extended_buf        int32
	Pad_cgo_2              [4]byte
	Side_data              **AvFrameSideData
	Nb_side_data           int32
	Flags                  int32
	Color_range            uint32
	Color_primaries        uint32
	Color_trc              uint32
	Colorspace             uint32
	Chroma_location        uint32
	Pad_cgo_3              [4]byte
	Best_effort_timestamp  int64
	Pkt_pos                int64
	Pkt_duration           int64
	Metadata               *Dictionary
	Decode_error_flags     int32
	Channels               int32
	Pkt_size               int32
	Pad_cgo_4              [4]byte
	Qscale_table           *int8
	Qstride                int32
	Qscale_type            int32
	Qp_table_buf           *AvBufferRef
	Hw_frames_ctx          *AvBufferRef
}
*/
