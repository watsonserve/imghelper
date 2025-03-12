package livp

import (
	libheif "github.com/strukturag/libheif-go"
	"gocv.io/x/gocv"
)

func readPrimary(ctx *libheif.Context) (gocv.Mat, error) {
	empty := gocv.Mat{}
	handler, err := ctx.GetPrimaryImageHandle()
	if nil != err {
		return empty, err
	}
	heicImg, err := handler.DecodeImage(libheif.ColorspaceRGB, libheif.ChromaInterleavedRGBA, nil)
	if nil != err {
		return empty, err
	}
	img, err := heicImg.GetImage()
	if nil != err {
		return empty, err
	}
	return gocv.ImageToMatRGBA(img)
}

func IMReadHeicPrimaryByMem(data []byte) (gocv.Mat, error) {
	ctx, err := libheif.NewContext()
	if nil == err {
		err = ctx.ReadFromMemory(data)
	}
	if nil != err {
		return gocv.Mat{}, err
	}
	return readPrimary(ctx)
}

func IMReadHeicPrimaryByFile(src string) (gocv.Mat, error) {
	ctx, err := libheif.NewContext()
	if nil == err {
		err = ctx.ReadFromFile(src)
	}
	if nil != err {
		return gocv.Mat{}, err
	}
	return readPrimary(ctx)
}
