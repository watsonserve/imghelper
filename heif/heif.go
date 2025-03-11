package heif

/*
#include "heif.h"
*/
import "C"
import "gocv.io/x/gocv"

func ReadPrimary(src []byte) ([]byte, error) {
	C.test(C.CString(string(src)))
	return nil, nil
}

func IMReadPrimary(src string) (gocv.Mat, error) {
	return gocv.Mat{}, nil
}
