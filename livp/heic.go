package livp

import (
	libheif "github.com/strukturag/libheif-go"
	"github.com/watsonserve/imghelper/cv"
)

func readPrimary(ctx *libheif.Context) (*cv.Mat, error) {
	handler, err := ctx.GetPrimaryImageHandle()
	if nil != err {
		return nil, err
	}
	heicImg, err := handler.DecodeImage(libheif.ColorspaceRGB, libheif.ChromaInterleavedRGBA, nil)
	if nil != err {
		return nil, err
	}
	img, err := heicImg.GetImage()
	if nil != err {
		return nil, err
	}
	imgMat, err := cv.ImageToMatRGBA(img)
	return &imgMat, err
}

func IMReadHeicPrimaryByMem(data []byte) (*cv.Mat, error) {
	ctx, err := libheif.NewContext()
	if nil == err {
		err = ctx.ReadFromMemory(data)
	}
	if nil != err {
		return nil, err
	}
	return readPrimary(ctx)
}

func IMReadHeicPrimaryByFile(src string) (*cv.Mat, error) {
	ctx, err := libheif.NewContext()
	if nil == err {
		err = ctx.ReadFromFile(src)
	}
	if nil != err {
		return nil, err
	}
	return readPrimary(ctx)
}
