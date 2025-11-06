package cv

// #cgo CXXFLAGS: -I /usr/include/opencv4 -O2 -Wall
// #cgo LDFLAGS: -lopencv_core -lopencv_imgproc -lopencv_imgcodecs
// #include "cv.h"
import "C"
import (
	"errors"
	"image"
	"image/color"
	"unsafe"
)

type Mat struct {
	p C.matptr_t
}

type InterpolationFlags int
type IMReadFlag int

const (
	// IMReadUnchanged return the loaded image as is (with alpha channel,
	// otherwise it gets cropped).
	IMReadUnchanged IMReadFlag = -1

	// IMReadGrayScale always convert image to the single channel
	// grayscale image.
	IMReadGrayScale IMReadFlag = 0

	// IMReadColor always converts image to the 3 channel BGR color image.
	IMReadColor IMReadFlag = 1

	// IMReadAnyDepth returns 16-bit/32-bit image when the input has the corresponding
	// depth, otherwise convert it to 8-bit.
	IMReadAnyDepth IMReadFlag = 2

	// IMReadAnyColor the image is read in any possible color format.
	IMReadAnyColor IMReadFlag = 4

	// IMReadLoadGDAL uses the gdal driver for loading the image.
	IMReadLoadGDAL IMReadFlag = 8

	// IMReadReducedGrayscale2 always converts image to the single channel grayscale image
	// and the image size reduced 1/2.
	IMReadReducedGrayscale2 IMReadFlag = 16

	// IMReadReducedColor2 always converts image to the 3 channel BGR color image and the
	// image size reduced 1/2.
	IMReadReducedColor2 IMReadFlag = 17

	// IMReadReducedGrayscale4 always converts image to the single channel grayscale image and
	// the image size reduced 1/4.
	IMReadReducedGrayscale4 IMReadFlag = 32

	// IMReadReducedColor4 always converts image to the 3 channel BGR color image and
	// the image size reduced 1/4.
	IMReadReducedColor4 IMReadFlag = 33

	// IMReadReducedGrayscale8 always convert image to the single channel grayscale image and
	// the image size reduced 1/8.
	IMReadReducedGrayscale8 IMReadFlag = 64

	// IMReadReducedColor8 always convert image to the 3 channel BGR color image and the
	// image size reduced 1/8.
	IMReadReducedColor8 IMReadFlag = 65

	// IMReadIgnoreOrientation do not rotate the image according to EXIF's orientation flag.
	IMReadIgnoreOrientation IMReadFlag = 128

	// InterpolationNearestNeighbor is nearest neighbor. (fast but low quality)
	InterpolationNearestNeighbor InterpolationFlags = 0

	// InterpolationLinear is bilinear interpolation.
	InterpolationLinear InterpolationFlags = 1

	// InterpolationCubic is bicube interpolation.
	InterpolationCubic InterpolationFlags = 2

	// InterpolationArea uses pixel area relation. It is preferred for image
	// decimation as it gives moire-free results.
	InterpolationArea InterpolationFlags = 3

	// InterpolationLanczos4 is Lanczos interpolation over 8x8 neighborhood.
	InterpolationLanczos4 InterpolationFlags = 4

	// InterpolationDefault is an alias for InterpolationLinear.
	InterpolationDefault = InterpolationLinear

	// InterpolationMax indicates use maximum interpolation.
	InterpolationMax InterpolationFlags = 7

	// WarpFillOutliers fills all of the destination image pixels. If some of them correspond to outliers in the source image, they are set to zero.
	WarpFillOutliers = 8

	// WarpInverseMap, inverse transformation.
	WarpInverseMap = 16

	// MatTypeCV8U is a Mat of 8-bit unsigned int
	MatTypeCV8U = 0

	// MatChannels3 is 3 channel Mat.
	MatChannels3 = 16

	// MatChannels4 is 4 channel Mat.
	MatChannels4 = 24

	// MatTypeCV8UC3 is a Mat of 8-bit unsigned int with 3 channels
	MatTypeCV8UC3 = MatTypeCV8U + MatChannels3

	// MatTypeCV8UC4 is a Mat of 8-bit unsigned int with 4 channels
	MatTypeCV8UC4 = MatTypeCV8U + MatChannels4

	ColorBGRAToRGBA = 5
	// IMWriteWebpQuality is the quality from 1 to 100 for WEBP (the higher is
	// the better). By default (without any parameter) and for quality above
	// 100 the lossless compression is used.
	IMWriteWebpQuality = 64
)

func toByteArray(buf []byte) (*C.uchar, int) {
	length := 0
	var data []C.uchar = nil

	if nil != buf {
		length = len(buf)
	}

	if 0 == length {
		return (*C.uchar)(unsafe.Pointer(nil)), 0
	}
	data = make([]C.uchar, length)
	for i, v := range buf {
		data[i] = C.uchar(v)
	}
	return (*C.uchar)(&data[0]), length
}

func toIntArray(buf []int) (*C.int, int) {
	length := 0
	var data []C.int = nil

	if nil != buf {
		length = len(buf)
	}

	if 0 == length {
		return (*C.int)(unsafe.Pointer(nil)), 0
	}
	data = make([]C.int, length)
	for i, v := range buf {
		data[i] = C.int(v)
	}
	return (*C.int)(&data[0]), length
}

func wrap(p C.matptr_t) Mat {
	return Mat{p: p}
}

func newMatFromBytes(rows int, cols int, mt int, data []byte) (Mat, error) {
	cBytes, length := toByteArray(data)
	if 0 == length {
		return wrap(nil), errors.New("no data")
	}
	return wrap(C.cvNewFromBytes(C.int(rows), C.int(cols), C.int(mt), cBytes)), nil
}

func ImageToMatRGBA(img image.Image) (Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()

	var data []uint8
	switch img.ColorModel() {
	case color.RGBAModel:
		m, res := img.(*image.RGBA)
		if !res {
			return wrap(nil), errors.New("Image color format error")
		}
		data = m.Pix

	case color.NRGBAModel:
		m, res := img.(*image.NRGBA)
		if !res {
			return wrap(nil), errors.New("Image color format error")
		}
		data = m.Pix

	default:
		data := make([]byte, 0, x*y*3)
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			for i := bounds.Min.X; i < bounds.Max.X; i++ {
				r, g, b, _ := img.At(i, j).RGBA()
				data = append(data, byte(b>>8), byte(g>>8), byte(r>>8))
			}
		}
		return newMatFromBytes(y, x, MatTypeCV8UC3, data)
	}

	// speed up the conversion process of RGBA format
	cvt, err := newMatFromBytes(y, x, MatTypeCV8UC4, data)
	if err != nil {
		return cvt, err
	}

	defer cvt.Close()

	dst := C.cvNew()
	C.cvtColor(cvt.p, dst, C.int(ColorBGRAToRGBA))
	return wrap(dst), nil
}

func IMRead(filename string, flags IMReadFlag) Mat {
	cStr := C.CString(filename)
	defer C.free(unsafe.Pointer(cStr))

	return wrap(C.cvIMRead(cStr, C.int(flags)))
}

func IMDecode(buf []byte, flags IMReadFlag) (Mat, error) {
	data, length := toByteArray(buf)
	if 0 == length {
		return wrap(nil), errors.New("no data")
	}
	return wrap(C.cvIMDecode(data, C.int(length), C.int(flags))), nil
}

func (mat *Mat) Empty() bool {
	if nil == mat.p {
		return true
	}
	return C.int(0) != C.cvEmpty(mat.p)
}

func (mat *Mat) Close() {
	C.cvClose(mat.p)
	mat.p = nil
}

func (mat *Mat) IMWrite(filename string, params []int) bool {
	if mat.Empty() {
		return false
	}

	cStr := C.CString(filename)
	defer C.free(unsafe.Pointer(cStr))

	cparams, length := toIntArray(params)

	return C.int(1) == C.cvImageIMWriteWithParams(cStr, mat.p, cparams, C.int(length))
}

func (mat *Mat) Resize(sz image.Point, fx, fy float64, interp InterpolationFlags) Mat {
	dst := wrap(C.cvNew())
	if !mat.Empty() {
		C.cvResize(mat.p, dst.p, C.int(sz.X), C.int(sz.Y), C.double(fx), C.double(fy), C.int(interp))
	}
	return dst
}

func (mat *Mat) Size() []int {
	if mat.Empty() {
		return []int{0, 0}
	}
	cdims := (*C.int)(unsafe.Pointer(nil))
	length := int(C.cvSize(mat.p, &cdims))
	defer C.free(unsafe.Pointer(cdims))

	pdims := unsafe.Slice(cdims, length)

	dims := make([]int, length)
	for i, v := range pdims {
		dims[i] = int(v)
	}
	return dims
}
