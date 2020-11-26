// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package libavutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package libavutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/dict.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type (
	AvDictionary      C.struct_AVDictionary
	AvDictionaryEntry C.struct_AVDictionaryEntry
)

// AV_DICT_xxx
const (
	AvDictMatchCase     = int(C.AV_DICT_MATCH_CASE)
	AvDictIgnoreSuffix  = int(C.AV_DICT_IGNORE_SUFFIX)
	AvDictDontStrdupKey = int(C.AV_DICT_DONT_STRDUP_KEY)
	AvDictDontStrdupVal = int(C.AV_DICT_DONT_STRDUP_VAL)
	AvDictDontOverwrite = int(C.AV_DICT_DONT_OVERWRITE)
	AvDictAppend        = int(C.AV_DICT_APPEND)
	AvDictMultikey      = int(C.AV_DICT_MULTIKEY)
)

// AvDictGet Get a dictionary entry with matching key.
func (d *AvDictionary) AvDictGet(key string, prev *AvDictionaryEntry, flags int) *AvDictionaryEntry {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	return (*AvDictionaryEntry)(C.av_dict_get(
		(*C.struct_AVDictionary)(d),
		Ckey,
		(*C.struct_AVDictionaryEntry)(prev),
		C.int(flags),
	))
}

// AvDictCount Get number of entries in dictionary.
func (d *AvDictionary) AvDictCount() int {
	return int(C.av_dict_count((*C.struct_AVDictionary)(d)))
}

// AvDictSet Set the given entry in *pm, overwriting an existing entry.
func (d *AvDictionary) AvDictSet(key, value string, flags int) int {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	Cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(Cvalue))

	return int(C.av_dict_set(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		Ckey,
		Cvalue,
		C.int(flags),
	))
}

// AvDictSetInt Convenience wrapper for av_dict_set that converts the value to a string and stores it.
func (d *AvDictionary) AvDictSetInt(key string, value int64, flags int) int {
	Ckey := C.CString(key)
	defer C.free(unsafe.Pointer(Ckey))

	return int(C.av_dict_set_int(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		Ckey,
		C.int64_t(value),
		C.int(flags),
	))
}

// AvDictParseString Parse the key/value pairs list and add the parsed entries to a dictionary.
func (d *AvDictionary) AvDictParseString(str, keyValSep, pairsSep string, flags int) int {
	Cstr := C.CString(str)
	defer C.free(unsafe.Pointer(Cstr))

	cKeyValSep := C.CString(keyValSep)
	defer C.free(unsafe.Pointer(cKeyValSep))

	cPairsSep := C.CString(pairsSep)
	defer C.free(unsafe.Pointer(cPairsSep))

	return int(C.av_dict_parse_string(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		Cstr,
		cKeyValSep,
		cPairsSep,
		C.int(flags),
	))
}

// AvDictCopy Copy entries from one AvDictionary struct into another.
func (d *AvDictionary) AvDictCopy(src *AvDictionary, flags int) int {
	return int(C.av_dict_copy(
		(**C.struct_AVDictionary)(unsafe.Pointer(&d)),
		(*C.struct_AVDictionary)(unsafe.Pointer(src)),
		C.int(flags),
	))
}

// AvDictFree Free all the memory allocated for an AvDictionary struct and all keys and values.
func (d *AvDictionary) AvDictFree() {
	C.av_dict_free((**C.struct_AVDictionary)(unsafe.Pointer(&d)))
}

// AvDictGetString Get dictionary entries as a string.
func (d *AvDictionary) AvDictGetString(keyValSep, pairsSep byte) (int, string) {
	var Cbuf *C.char

	ret := int(C.av_dict_get_string(
		(*C.struct_AVDictionary)(d),
		(**C.char)(&Cbuf),
		C.char(keyValSep),
		C.char(pairsSep),
	))

	var buf string
	if ret == 0 {
		buf = C.GoString(Cbuf)
		C.free(unsafe.Pointer(Cbuf))
	}

	return ret, buf
}

// Key Return key
func (d *AvDictionaryEntry) Key() string {
	return C.GoString(d.key)
}

// Value Return value
func (d *AvDictionaryEntry) Value() string {
	return C.GoString(d.value)
}
