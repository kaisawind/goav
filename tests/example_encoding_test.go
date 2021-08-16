package tests

import (
	"fmt"
	"github.com/asticode/goav/avcodec"
	"github.com/asticode/goav/avformat"
	"github.com/asticode/goav/avutil"
	"unsafe"
)

func Example() {
	outputfilename := "testdata/sample-encoding1.mpg"
	dstWidth, dstHeight := 640, 480

	ccId := avcodec.CodecId(avcodec.CodecIdMpeg1VIDEO)
	cc := avcodec.FindEncoder(ccId)
	if cc == nil {
		fmt.Printf("Could not find encoder for '%s'\n", avcodec.GetName(ccId))
		return
	}
	ctx := cc.AllocContext3()
	defer ctx.Free()
	ctx.SetCodecId(ccId)
	var ofmtCtx *avformat.Context
	var ofmt *avformat.OutputFormat
	ret := avformat.AllocOutputContext2(&ofmtCtx, ofmt, "", outputfilename)
	if ret < 0 {
		fmt.Println("AllocOutputContext2 error", avutil.AvError(ret))
		return
	}
	if ofmtCtx == nil {
		return
	}
	defer ofmtCtx.Free()
	avutil.AvOptSetDefaults(unsafe.Pointer(ofmtCtx))

	ctx.SetBitRate(400000).
		SetWidth(dstWidth).
		SetHeight(dstHeight).
		SetTimeBase(avutil.NewRational(1, 25)).
		SetPixFmt(avutil.AvPixFmtYuv420P).
		SetProfile(avcodec.FfProfileMpeg4Simple).
		SetMbDecision(avcodec.FfMbDecisionRd)

	if ofmtCtx != nil && ofmtCtx.Oformat() != nil && ofmtCtx.Oformat().Flags()&avformat.AvfmtGlobalHeader != 0 {
		ctx.SetFlags(ctx.Flags() | avcodec.CodecFlagGlobalHeader)
		fmt.Printf("ctx.Flags(): %02X \n", ctx.Flags())
	}
	stream := ofmtCtx.NewStream()
	if stream == nil {
		fmt.Println("new stream error")
		return
	}
	if ctx.IsOpen() == 0 {
		ret = ctx.Open2(cc, nil)
		if ret < 0 {
			fmt.Println("Open2 error", avutil.AvError(ret))
			return
		}
	}
	ret = avcodec.ParametersFromContext(stream.CodecParameters(), ctx)
	if ret != 0 {
		fmt.Println("ParametersFromContext error", avutil.AvError(ret))
		return
	}
	ofmtCtx.SetStartTime(0)
	if ofmtCtx.Oformat().Flags()&avformat.AvfmtNoFile == 0 && ofmtCtx.Pb() == nil {
		pb := ofmtCtx.Pb()
		ret = avformat.AvIOOpen(&pb, ofmtCtx.URL(), avformat.AvioFlagWrite)
		if ret < 0 {
			fmt.Println("AvIOOpen error", avutil.AvError(ret))
			return
		}
	}
	ret = ofmtCtx.WriteHeader(nil)
	if ret < 0 {
		fmt.Println("WriteHeader error", avutil.AvError(ret))
		return
	}
	ofmtCtx.AvDumpFormat(0, ofmtCtx.URL(), 1)
For:
	for i := 0; i < 25; i++ {
		f := avutil.AvFrameAlloc().
			SetWidth(ctx.Width()).
			SetHeight(ctx.Height()).
			SetFormat(ctx.PixFmt()).
			SetPts(int64(i))

		ret = f.AvImageAlloc()
		if ret < 0 {
			fmt.Println("AvImageAlloc error", avutil.AvError(ret))
			return
		}
		for y := 0; y < f.Height(); y++ {
			for x := 0; x < f.Width(); x++ {
				f.SetData(0, y*f.LineSize()+x, byte(x+y+i*3))
			}
		}

		// Cb and Cr
		for y := 0; y < f.Height()/2; y++ {
			for x := 0; x < f.Width()/2; x++ {
				f.SetData(1, y*f.LineSizeItem(1)+x, byte(128+y+i*2))
				f.SetData(2, y*f.LineSizeItem(2)+x, byte(64+y+i*2))
			}
		}

		ret = avcodec.SendFrame(ctx, f)
		if ret < 0 {
			fmt.Println("SendFrame error", avutil.AvError(ret))
			return
		}
		p := avcodec.AvPacketAlloc()
		for {
			ret = avcodec.ReceivePacket(ctx, p)
			if ret == avutil.AvErrorEagain {
				fmt.Println("ReceivePacket error ", avutil.AvError(ret))
				p.Free()
				continue For
			}
			if ret >= 0 {
				break
			} else {
				fmt.Println("ReceivePacket error ", avutil.AvError(ret))
				return
			}
		}
		if p.Pts() != avutil.AvNoPtsValue {
			p.SetPts(avutil.AvRescaleQ(p.Pts(), ctx.TimeBase(), stream.TimeBase()))
		}

		if p.Dts() != avutil.AvNoPtsValue {
			p.SetDts(avutil.AvRescaleQ(p.Pts(), ctx.TimeBase(), stream.TimeBase()))
		}
		ret = ofmtCtx.AvInterleavedWriteFrame(p)
		if ret < 0 {
			fmt.Println("AvInterleavedWriteFrame error", avutil.AvError(ret))
			return
		}
		p.Free()
		f.Free()
	}
	// Output: frames written to examples/sample-encoding1.mpg
}
