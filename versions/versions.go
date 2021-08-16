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
	log.Printf("AvFilter Version:\t%s", avutil.AvVersion(avfilter.Version()))
	log.Printf("AvDevice Version:\t%s", avutil.AvVersion(avdevice.Version()))
	log.Printf("SWScale Version:\t%s", avutil.AvVersion(swscale.Version()))
	log.Printf("AvUtil Version:\t%s", avutil.AvVersion(avutil.Version()))
	log.Printf("AvCodec Version:\t%s", avutil.AvVersion(avcodec.Version()))
	log.Printf("Resample Version:\t%s", avutil.AvVersion(swresample.Version()))
}
