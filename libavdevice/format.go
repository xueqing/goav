// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavdevice

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"

// AvInputAudioDeviceNext Audio input devices iterator.
func (d *AvInputFormat) AvInputAudioDeviceNext() *AvInputFormat {
	return (*AvInputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

// AvInputVideoDeviceNext Video input devices iterator.
func (d *AvInputFormat) AvInputVideoDeviceNext() *AvInputFormat {
	return (*AvInputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

// AvOutputAudioDeviceNext Audio output devices iterator.
func (d *AvOutputFormat) AvOutputAudioDeviceNext() *AvOutputFormat {
	return (*AvOutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

// AvOutputVideoDeviceNext Video output devices iterator.
func (d *AvOutputFormat) AvOutputVideoDeviceNext() *AvOutputFormat {
	return (*AvOutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}
