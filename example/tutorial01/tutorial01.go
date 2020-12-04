package main

// tutorial01.c
// Code based on a tutorial at http://dranger.com/ffmpeg/tutorial01.html

// A small sample program that shows how to use libavformat and libavcodec to
// read video from a file.
//
// Use
//
// gcc -o tutorial01 tutorial01.c -lavformat -lavcodec -lswscale -lz
//
// to build (assuming libavformat and libavcodec are correctly installed
// your system).
//
// Run using
//
// tutorial01 myvideofile.mpg
//
// to write the first five frames from "myvideofile.mpg" to disk in PPM
// format.
import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/xueqing/goav/libavcodec"
	"github.com/xueqing/goav/libavformat"
	"github.com/xueqing/goav/libavutil"
	"github.com/xueqing/goav/libswscale"
)

// SaveFrame writes a single frame to disk as a PPM file
func SaveFrame(frame *libavutil.AvFrame, width, height, frameNumber int) {
	// Open file
	fileName := fmt.Sprintf("frame%d.ppm", frameNumber)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error Reading")
	}
	defer file.Close()

	// Write header
	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
	file.Write([]byte(header))

	// Write pixel data
	for y := 0; y < height; y++ {
		buf := make([]byte, width*3)
		startPos := uintptr(unsafe.Pointer(*frame.Data())) + uintptr(y)*uintptr(*frame.Linesize())
		for i := 0; i < width*3; i++ {
			element := *(*uint8)(unsafe.Pointer(startPos + uintptr(i)))
			buf[i] = element
		}
		file.Write(buf)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie file")
		os.Exit(1)
	}

	// Open video file
	pFormatContext := libavformat.AvformatAllocContext()
	if libavformat.AvformatOpenInput(&pFormatContext, os.Args[1], nil, nil) != 0 {
		fmt.Printf("Unable to open file %s\n", os.Args[1])
		os.Exit(1)
	}

	// Retrieve stream information
	if pFormatContext.AvformatFindStreamInfo(nil) < 0 {
		fmt.Println("Couldn't find stream information")
		os.Exit(1)
	}

	// Dump information about file onto standard error
	pFormatContext.AvDumpFormat(0, os.Args[1], 0)

	// Find the first video stream
	for i := 0; i < int(pFormatContext.NbStreams()); i++ {
		codecType := libavutil.AvMediaType(pFormatContext.Streams()[i].CodecParameters().CodecType())
		switch codecType {
		case libavutil.AvmediaTypeVideo:

			// Get a pointer to the codec context for the video stream
			pCodecCtxOrig := (*libavcodec.AvCodecContext)(unsafe.Pointer(pFormatContext.Streams()[i].Codec()))
			// Find the decoder for the video stream
			pCodec := libavcodec.AvcodecFindDecoder(libavcodec.AvCodecID(pCodecCtxOrig.CodecID()))
			if pCodec == nil {
				fmt.Println("Unsupported codec!")
				os.Exit(1)
			}
			// Copy context
			pCodecCtx := pCodec.AvcodecAllocContext3()
			if pCodecCtx.AvcodecParametersToContext(pFormatContext.Streams()[i].CodecParameters()) != 0 {
				fmt.Println("Couldn't copy codec context")
				os.Exit(1)
			}

			// Open codec
			if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
				fmt.Println("Could not open codec")
				os.Exit(1)
			}

			// Allocate video frame
			pFrame := libavutil.AvFrameAlloc()

			// Allocate an AvFrame structure
			pFrameRGB := libavutil.AvFrameAlloc()
			if pFrameRGB == nil {
				fmt.Println("Unable to allocate RGB Frame")
				os.Exit(1)
			}

			// Determine required buffer size and allocate buffer
			numBytes := uintptr(libavutil.AvImageGetBufferSize(libavutil.AvPixelFormat(libavcodec.AvPixFmtRgb24), pCodecCtx.Width(),
				pCodecCtx.Height(), 1))
			buffer := libavutil.AvMalloc(numBytes)

			// Assign appropriate parts of buffer to image planes in pFrameRGB
			if ret := libavutil.AvImageFillArrays(pFrameRGB.Data(), pFrameRGB.Linesize(), (*uint8)(buffer),
				libavutil.AvPixelFormat(libavcodec.AvPixFmtRgb24), pCodecCtx.Width(), pCodecCtx.Height(), 1); ret < 0 {
				fmt.Printf("Error while filling an image: %s\n", libavutil.ErrorFromCode(ret))
			}

			// initialize SWS context for software scaling
			swsCtx := libswscale.SwsGetcontext(
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				(libswscale.AvPixelFormat)(pCodecCtx.PixFmt()),
				pCodecCtx.Width(),
				pCodecCtx.Height(),
				libavcodec.AvPixFmtRgb24,
				libavcodec.SwsBilinear,
				nil,
				nil,
				nil,
			)

			// Read frames and save first five frames to disk
			frameNumber := 1
			packet := libavcodec.AvPacketAlloc()
			for pFormatContext.AvReadFrame(packet) >= 0 {
				// Is this a packet from the video stream?
				if packet.StreamIndex() == i {
					// Decode video frame
					response := pCodecCtx.AvcodecSendPacket(packet)
					if response < 0 {
						fmt.Printf("Error while sending a packet to the decoder: %s\n", libavutil.ErrorFromCode(response))
					}
					for response >= 0 {
						response = pCodecCtx.AvcodecReceiveFrame((*libavcodec.AvFrame)(unsafe.Pointer(pFrame)))
						if response == libavutil.AvErrorEAGAIN || response == libavutil.AvErrorEOF {
							break
						} else if response < 0 {
							fmt.Printf("Error while receiving a frame from the decoder: %s\n", libavutil.ErrorFromCode(response))
							return
						}

						if frameNumber <= 5 {
							// Convert the image from its native format to RGB
							libswscale.SwsScale2(swsCtx, pFrame.Data(),
								pFrame.Linesize(), 0, pCodecCtx.Height(),
								pFrameRGB.Data(), pFrameRGB.Linesize())

							// Save the frame to disk
							fmt.Printf("Writing frame %d\n", frameNumber)
							SaveFrame(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(), frameNumber)
						} else {
							return
						}
						frameNumber++
					}
				}

				// Free the packet that was allocated by av_read_frame
				packet.AvPacketUnref()
			}

			// Free the RGB image
			libavutil.AvFree(buffer)
			libavutil.AvFrameFree(pFrameRGB)

			// Free the YUV frame
			libavutil.AvFrameFree(pFrame)

			// Close the codecs
			pCodecCtx.AvcodecClose()
			(*libavcodec.AvCodecContext)(unsafe.Pointer(pCodecCtxOrig)).AvcodecClose()

			// Close the video file
			pFormatContext.AvformatCloseInput()

			// Stop after saving frames of first video straem
			break

		default:
			fmt.Println("Didn't find a video stream")
			os.Exit(1)
		}
	}
}
