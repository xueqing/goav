package main

import (
	"log"

	"github.com/xueqing/goav/avcodec"
	"github.com/xueqing/goav/avdevice"
	"github.com/xueqing/goav/avfilter"
	"github.com/xueqing/goav/avutil"
	"github.com/xueqing/goav/swresample"
	"github.com/xueqing/goav/swscale"
)

func main() {
	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())
}
