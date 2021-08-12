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

func (s *Stream) AttachedPic() Packet {
	return Packet(s.attached_pic)
}

func (s *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(s.side_data))
}

func (s *Stream) ProbeData() AvProbeData {
	return AvProbeData(s.probe_data)
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

func (s *Stream) InjectGlobalSideData() int {
	return int(s.inject_global_side_data)
}

func (s *Stream) LastIpDuration() int {
	return int(s.last_IP_duration)
}

func (s *Stream) NbDecodedFrames() int {
	return int(s.nb_decoded_frames)
}

func (s *Stream) NbIndexEntries() int {
	return int(s.nb_index_entries)
}

func (s *Stream) NbSideData() int {
	return int(s.nb_side_data)
}

func (s *Stream) ProbePackets() int {
	return int(s.probe_packets)
}

func (s *Stream) PtsWrapBehavior() int {
	return int(s.pts_wrap_behavior)
}

func (s *Stream) RequestProbe() int {
	return int(s.request_probe)
}

func (s *Stream) SkipSamples() int {
	return int(s.skip_samples)
}

func (s *Stream) SkipToKeyframe() int {
	return int(s.skip_to_keyframe)
}

func (s *Stream) StreamIdentifier() int {
	return int(s.stream_identifier)
}

func (s *Stream) UpdateInitialDurationsDone() int {
	return int(s.update_initial_durations_done)
}

func (s *Stream) CurDts() int64 {
	return int64(s.cur_dts)
}

func (s *Stream) Duration() int64 {
	return int64(s.duration)
}

// func (avs *Stream) FirstDiscardSample() int64 {
// 	return int64(avs.first_discard_sample)
// }

func (s *Stream) FirstDts() int64 {
	return int64(s.first_dts)
}

func (s *Stream) InterleaverChunkDuration() int64 {
	return int64(s.interleaver_chunk_duration)
}

func (s *Stream) InterleaverChunkSize() int64 {
	return int64(s.interleaver_chunk_size)
}

// func (avs *Stream) LastDiscardSample() int64 {
// 	return int64(avs.last_discard_sample)
// }

func (s *Stream) LastDtsForOrderCheck() int64 {
	return int64(s.last_dts_for_order_check)
}

func (s *Stream) LastIpPts() int64 {
	return int64(s.last_IP_pts)
}

func (s *Stream) MuxTsOffset() int64 {
	return int64(s.mux_ts_offset)
}

func (s *Stream) NbFrames() int64 {
	return int64(s.nb_frames)
}

func (s *Stream) PtsBuffer() int64 {
	return int64(s.pts_buffer[0])
}

func (s *Stream) PtsReorderError() int64 {
	return int64(s.pts_reorder_error[0])
}

func (s *Stream) PtsWrapReference() int64 {
	return int64(s.pts_wrap_reference)
}

// func (avs *Stream) StartSkipSamples() int64 {
// 	return int64(avs.start_skip_samples)
// }

func (s *Stream) StartTime() int64 {
	return int64(s.start_time)
}

func (s *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(s.parser))
}

func (s *Stream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(s.last_in_packet_buffer))
}

func (s *Stream) DtsMisordered() uint8 {
	return uint8(s.dts_misordered)
}

func (s *Stream) DtsOrdered() uint8 {
	return uint8(s.dts_ordered)
}

func (s *Stream) PtsReorderErrorCount() uint8 {
	return uint8(s.pts_reorder_error_count[0])
}

func (s *Stream) IndexEntriesAllocatedSize() uint {
	return uint(s.index_entries_allocated_size)
}
