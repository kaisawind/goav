package main

import (
	"log"

	"github.com/asticode/goav/avcodec"
	"github.com/asticode/goav/avdevice"
	"github.com/asticode/goav/avfilter"
	"github.com/asticode/goav/avutil"
	"github.com/asticode/goav/swresample"
	"github.com/asticode/goav/swscale"
)

func main() {
	log.Printf("AvFilter Version:\t%v", avfilter.Version())
	log.Printf("AvDevice Version:\t%v", avdevice.Version())
	log.Printf("SWScale Version:\t%v", swscale.Version())
	log.Printf("AvUtil Version:\t%v", avutil.Version())
	log.Printf("AvCodec Version:\t%v", avcodec.Version())
	log.Printf("Resample Version:\t%v", swresample.Version())
}
