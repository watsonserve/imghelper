package cr2

/*
#include "cr2_thumbnail.h"
*/
import "C"

import (
	"os"
	"unsafe"

	"gocv.io/x/gocv"
)

func cr2Thumb(src string) ([]byte, error) {
	fpSrc, err := os.Open(src)
	if nil != err {
		return nil, err
	}
	defer fpSrc.Close()
	fdptr := fpSrc.Fd()
	fd := *(*C.int)(unsafe.Pointer(&fdptr))

	var off C.int64_t = 0
	var length C.uint64_t = 0
	offP := (*C.int64_t)(unsafe.Pointer(&off))
	lenP := (*C.uint64_t)(unsafe.Pointer(&length))
	C.cr2_thumbnail(fd, offP, lenP)
	buf := make([]byte, length)
	fpSrc.ReadAt(buf, int64(off))
	return buf, nil
}

func IMReadThumb(src string) (gocv.Mat, error) {
	buf, err := cr2Thumb(src)
	if nil != err {
		return gocv.Mat{}, err
	}
	return gocv.IMDecode(buf, gocv.IMReadUnchanged)
}
