package imghelper

import (
	"errors"
	"image"
	"math"
	"path"
	"strings"

	"github.com/watsonserve/imghelper/cr2"
	"github.com/watsonserve/imghelper/livp"
	"gocv.io/x/gocv"
)

func reSizeByLongSide(imgMat *gocv.Mat, maxLongSide int) {
	siz := imgMat.Size()
	height := siz[0]
	width := siz[1]
	longSide := width
	if width < height {
		longSide = height
	}
	if longSide <= maxLongSide {
		return
	}
	scale := 1 / math.Ceil(float64(longSide)/float64(maxLongSide))
	sz := image.Point{}
	gocv.Resize(*imgMat, imgMat, sz, scale, scale, gocv.InterpolationArea)
}

func IMLoad(fileName string, content []byte) (*gocv.Mat, error) {
	switch strings.ToLower(path.Ext(fileName)) {
	case ".cr2":
		return cr2.IMReadThumb(fileName)
	case ".livp":
		return livp.IMReadLivpPrimary(fileName)
	case ".heic":
		return livp.IMReadHeicPrimaryByFile(fileName)
	default:
	}
	imgMat, err := gocv.IMDecode(content, gocv.IMReadUnchanged)
	if nil != err {
		return nil, err
	}
	if imgMat.Empty() {
		return nil, errors.New("load Image failed")
	}
	return &imgMat, nil
}

func IMRead(src string) (*gocv.Mat, error) {
	switch strings.ToLower(path.Ext(src)) {
	case ".cr2":
		return cr2.IMReadThumb(src)
	case ".livp":
		return livp.IMReadLivpPrimary(src)
	case ".heic":
		return livp.IMReadHeicPrimaryByFile(src)
	default:
	}
	imgMat := gocv.IMRead(src, gocv.IMReadUnchanged)
	if imgMat.Empty() {
		return nil, errors.New("load Image failed")
	}
	return &imgMat, nil
}

func IMWrite(imgMat *gocv.Mat, dst string, quality, maxLongSide int) error {
	if 0 < maxLongSide {
		reSizeByLongSide(imgMat, maxLongSide)
	}
	params := make([]int, 0)
	if 0 < quality && quality < 100 {
		params = []int{int(gocv.IMWriteWebpQuality), int(quality)}
	}
	if !gocv.IMWriteWithParams(dst, *imgMat, params) {
		return errors.New("save image failed")
	}
	return nil
}
