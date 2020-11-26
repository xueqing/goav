// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/avutil.h>
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// AvMalloc Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU).
func AvMalloc(size uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.av_malloc(C.size_t(size)))
}

// AvMallocArray Allocate a memory block for an array with av_malloc().
func AvMallocArray(nmemb, size uintptr) unsafe.Pointer {
	return C.av_malloc_array(C.size_t(nmemb), C.size_t(size))
}

// AvRealloc Allocate or reallocate a block of memory.
func AvRealloc(ptr *int, size uintptr) unsafe.Pointer {
	return C.av_realloc(unsafe.Pointer(&ptr), C.size_t(size))
}

// AvReallocF Allocate or reallocate a block of memory.
func AvReallocF(ptr int, nelem, elsize uintptr) unsafe.Pointer {
	return C.av_realloc_f(unsafe.Pointer(&ptr), C.size_t(nelem), C.size_t(elsize))
}

// AvReallocp Allocate or reallocate a block of memory.
func AvReallocp(ptr int, size uintptr) int {
	return int(C.av_reallocp(unsafe.Pointer(&ptr), C.size_t(size)))
}

// AvReallocArray Allocate, reallocate, or free an array.
func AvReallocArray(ptr int, nmemb, size uintptr) unsafe.Pointer {
	return C.av_realloc_array(unsafe.Pointer(&ptr), C.size_t(nmemb), C.size_t(size))
}

// AvReallocpArray Allocate, reallocate, or free an array through a pointer to a pointer.
func AvReallocpArray(ptr int, nmemb, size uintptr) int {
	return int(C.av_reallocp_array(unsafe.Pointer(&ptr), C.size_t(nmemb), C.size_t(size)))
}

// AvFree Free a memory block which has been allocated with av_malloc(z)() or av_realloc().
func AvFree(ptr unsafe.Pointer) {
	C.av_free(ptr)
}

// AvMallocz Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func AvMallocz(size uintptr) unsafe.Pointer {
	return C.av_mallocz(C.size_t(size))
}

// AvCalloc Allocate a block of nmemb * size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func AvCalloc(nmemb, size uintptr) unsafe.Pointer {
	return C.av_calloc(C.size_t(nmemb), C.size_t(size))
}

// AvMalloczArray Allocate a memory block for an array with av_mallocz().
func AvMalloczArray(nmemb, size uintptr) unsafe.Pointer {
	return C.av_mallocz_array(C.size_t(nmemb), C.size_t(size))
}

// AvStrdup Duplicate the string size.
func AvStrdup(s string) string {
	return C.GoString(C.av_strdup(C.CString(s)))
}

// AvStrndup Duplicate a substring of the string size.
func AvStrndup(s string, length uintptr) string {
	return C.GoString(C.av_strndup(C.CString(s), C.size_t(length)))
}

// AvMemdup Duplicate the buffer ptr.
func AvMemdup(ptr *int, size uintptr) unsafe.Pointer {
	return C.av_memdup(unsafe.Pointer(ptr), C.size_t(size))
}

// AvFreep Free a memory block which has been allocated with av_malloc(z)() or av_realloc() and set the pointer pointing to it to NULL.
func AvFreep(ptr unsafe.Pointer) {
	C.av_freep(ptr)
}

// AvDynarrayAdd Add an element to a dynamic array.
func AvDynarrayAdd(tabPtr unsafe.Pointer, nbPtr *int, elem unsafe.Pointer) {
	C.av_dynarray_add(tabPtr, (*C.int)(unsafe.Pointer(nbPtr)), elem)
}

// AvDynarrayAddNofree Add an element to a dynamic array.
func AvDynarrayAddNofree(tabPtr unsafe.Pointer, nbPtr *int, elem unsafe.Pointer) int {
	return int(C.av_dynarray_add_nofree(tabPtr, (*C.int)(unsafe.Pointer(nbPtr)), elem))
}

// AvDynarray2Add Add an element of size elem_size to a dynamic array.
func AvDynarray2Add(tabPtr *unsafe.Pointer, nbPtr *int, elemSize uintptr, elemData uint8) unsafe.Pointer {
	return C.av_dynarray2_add(tabPtr, (*C.int)(unsafe.Pointer(&nbPtr)),
		(C.size_t)(elemSize), (*C.uint8_t)(unsafe.Pointer(&elemData)))
}

// AvSizeMult Multiply two size_t values checking for overflow.
func AvSizeMult(a, b uintptr, r *uintptr) int {
	return int(C.av_size_mult((C.size_t)(a), (C.size_t)(b), (*C.size_t)(unsafe.Pointer(r))))
}

// AvMaxAlloc Set the maximum size that may me allocated in one block.
func AvMaxAlloc(max uintptr) {
	C.av_max_alloc(C.size_t(max))
}

// AvMemcpyBackptr deliberately overlapping memcpy implementation
func AvMemcpyBackptr(dst *uintptr, back, cnt int) {
	C.av_memcpy_backptr((*C.uint8_t)(unsafe.Pointer(dst)), C.int(back), C.int(cnt))
}

// AvFastRealloc Reallocate the given block if it is not large enough, otherwise do nothing.
func AvFastRealloc(ptr unsafe.Pointer, size *uint, minSize uintptr) unsafe.Pointer {
	return C.av_fast_realloc(ptr, (*C.uint)(unsafe.Pointer(size)), (C.size_t)(minSize))
}

// AvFastMalloc Allocate a buffer, reusing the given one if large enough.
func AvFastMalloc(ptr unsafe.Pointer, size *uint, minSize uintptr) {
	C.av_fast_malloc(ptr, (*C.uint)(unsafe.Pointer(size)), (C.size_t)(minSize))
}
