package livp

import (
	"archive/zip"
	"bytes"
	"io"
	"path"
	"strings"

	"github.com/watsonserve/images/heif"

	"gocv.io/x/gocv"
)

func exportZFile(zf *zip.File) ([]byte, error) {
	fileSize := zf.FileInfo().Size()
	srcFp, err := zf.Open()
	if nil != err {
		return nil, err
	}
	defer srcFp.Close()
	dstFp := bytes.NewBuffer(make([]byte, fileSize))
	if nil != err {
		return nil, err
	}
	_, err = io.Copy(dstFp, srcFp)
	return dstFp.Bytes(), err
}

func ReadPrimary(src string) (img []byte, isHeic bool, err error) {
	img = nil
	isHeic = false
	reader, err := zip.OpenReader(src)
	if nil != err {
		return nil, false, err
	}
	defer reader.Close()

	for _, item := range reader.File {
		extName := strings.ToLower(path.Ext(item.Name))
		switch extName {
		case ".heic":
			isHeic = true
			fallthrough
		case ".jpg":
			fallthrough
		case ".jpeg":
			content, _err := exportZFile(item)
			err = _err
			if nil == err && (isHeic || nil == img) {
				img = content
			}
		}
	}
	return img, isHeic, err
}

func IMReadPrimary(src string) (gocv.Mat, error) {
	buf, isHeic, err := ReadPrimary(src)
	if isHeic && nil == err {
		buf, err = heif.ReadPrimary(buf)
	}
	if nil != err {
		return gocv.Mat{}, err
	}
	return gocv.IMDecode(buf, gocv.IMReadUnchanged)
}
