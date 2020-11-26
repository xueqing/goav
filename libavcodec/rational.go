// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package libavcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package libavcodec

//#cgo pkg-config: libavformat libavcodec libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import "fmt"

// Num Return num
func (r Rational) Num() int {
	return int(r.num)
}

// Den Return den
func (r Rational) Den() int {
	return int(r.den)
}

// String ...
func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", int(r.num), int(r.den))
}

// Assign ...
func (r *Rational) Assign(o Rational) {
	r.Set(o.Num(), o.Den())
}

// Set ...
func (r *Rational) Set(num, den int) {
	r.num, r.den = C.int(num), C.int(den)
}

// NewRational ...
func NewRational(num, den int) Rational {
	return Rational{
		num: C.int(num),
		den: C.int(den),
	}
}
