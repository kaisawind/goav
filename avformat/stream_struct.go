// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
//#include <libavutil/rational.h>
import "C"
import (
	"unsafe"

	"github.com/asticode/goav/avcodec"
	"github.com/asticode/goav/avutil"
)

func (s *Stream) Codec() *CodecContext {
	return (*CodecContext)(unsafe.Pointer(s.codec))
}

func (s *Stream) CodecParameters() *avcodec.CodecParameters {
	return (*avcodec.CodecParameters)(unsafe.Pointer(s.codecpar))
}

func (s *Stream) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(s.metadata))
}

func (s *Stream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(s.index_entries))
}

func (s *Stream) AttachedPic() avcodec.Packet {
	return *(*avcodec.Packet)(unsafe.Pointer(&s.attached_pic))
}

func (s *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(s.side_data))
}

func (s *Stream) AvgFrameRate() avutil.Rational {
	return *(*avutil.Rational)(unsafe.Pointer(&s.avg_frame_rate))
}

// func (avs *Stream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

func (s *Stream) RFrameRate() avutil.Rational {
	return *(*avutil.Rational)(unsafe.Pointer(&s.r_frame_rate))
}

func (s *Stream) SampleAspectRatio() avutil.Rational {
	return *(*avutil.Rational)(unsafe.Pointer(&s.sample_aspect_ratio))
}

func (s *Stream) TimeBase() avutil.Rational {
	return *(*avutil.Rational)(unsafe.Pointer(&s.time_base))
}

func (s *Stream) SetTimeBase(r avutil.Rational) {
	s.time_base = *((*C.struct_AVRational)(unsafe.Pointer(&r)))
}

func (s *Stream) Discard() AvDiscard {
	return AvDiscard(s.discard)
}

func (s *Stream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(s.need_parsing)
}

func (s *Stream) CodecInfoNbFrames() int {
	return int(s.codec_info_nb_frames)
}

func (s *Stream) Disposition() int {
	return int(s.disposition)
}

func (s *Stream) EventFlags() int {
	return int(s.event_flags)
}

func (s *Stream) Id() int {
	return int(s.id)
}

func (s *Stream) Index() int {
	return int(s.index)
}

func (s *Stream) LastIpDuration() int {
	return int(s.last_IP_duration)
}

func (s *Stream) NbIndexEntries() int {
	return int(s.nb_index_entries)
}

func (s *Stream) ProbePackets() int {
	return int(s.probe_packets)
}

func (s *Stream) StreamIdentifier() int {
	return int(s.stream_identifier)
}

func (s *Stream) CurDts() int64 {
	return int64(s.cur_dts)
}

func (s *Stream) Duration() int64 {
	return int64(s.duration)
}

func (s *Stream) FirstDts() int64 {
	return int64(s.first_dts)
}

func (s *Stream) LastIpPts() int64 {
	return int64(s.last_IP_pts)
}

func (s *Stream) NbFrames() int64 {
	return int64(s.nb_frames)
}

func (s *Stream) StartTime() int64 {
	return int64(s.start_time)
}

func (s *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(s.parser))
}

func (s *Stream) IndexEntriesAllocatedSize() uint {
	return uint(s.index_entries_allocated_size)
}
