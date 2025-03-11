package main

import (
	"errors"
	"fmt"
	"image"
	"math"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/watsonserve/images/cr2"
	"github.com/watsonserve/images/heif"
	"github.com/watsonserve/images/livp"
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

func getArgs() (int64, string, string) {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "usage: images quality(0-100) dst.jpg src.cr2")
		return 0, "", ""
	}
	strQuality := os.Args[1]
	dst := os.Args[2]
	src := os.Args[3]

	quality, err := strconv.ParseInt(strQuality, 10, 32)
	if nil != err {
		panic(err)
	}
	if quality < 1 || 100 < quality {
		panic(errors.New("quality must between 0 and 100"))
	}
	return quality, dst, src
}

func IMRead(src string) (gocv.Mat, error) {
	var imgMat gocv.Mat
	var err error
	switch strings.ToLower(path.Ext(src)) {
	case ".cr2":
		imgMat, err = cr2.IMReadThumb(src)
	case ".livp":
		imgMat, err = livp.IMReadPrimary(src)
	case ".heic":
		imgMat, err = heif.IMReadPrimary(src)
	default:
		imgMat = gocv.IMRead(src, gocv.IMReadUnchanged)
		if imgMat.Empty() {
			err = errors.New("load Image failed")
		}
	}
	return imgMat, err
}

func main() {
	quality, dst, src := getArgs()
	imgMat, err := IMRead(src)
	if nil != err {
		panic(err)
	}
	reSizeByLongSide(&imgMat, 684)
	gocv.IMWriteWithParams(dst, imgMat, []int{int(gocv.IMWriteWebpQuality), int(quality)})
}
