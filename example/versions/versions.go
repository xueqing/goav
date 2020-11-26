package main

import (
	"log"

	"github.com/xueqing/goav/libavcodec"
	"github.com/xueqing/goav/libavdevice"
	"github.com/xueqing/goav/libavfilter"
	"github.com/xueqing/goav/libavformat"
	"github.com/xueqing/goav/libavutil"
	"github.com/xueqing/goav/libswresample"
	"github.com/xueqing/goav/libswscale"
)

func main() {
	log.Printf("AvCodec  Version:\t%v", libavcodec.AvcodecVersion())
	log.Printf("AvDevice Version:\t%v", libavdevice.AvdeviceVersion())
	log.Printf("AvFilter Version:\t%v", libavfilter.AvfilterVersion())
	log.Printf("AvFormat Version:\t%v", libavformat.AvformatVersion())
	log.Printf("AvUtil   Version:\t%v", libavutil.AvutilVersion())
	log.Printf("Resample Version:\t%v", libswresample.SwresampleLicense())
	log.Printf("SwScale  Version:\t%v", libswscale.SwscaleVersion())
}
