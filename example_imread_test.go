package imghelper_test

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/watsonserve/imghelper"
)

func getArgs() (int, string, string) {
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
	return int(quality), dst, src
}

func ExampleIMWrite() {
	quality, dst, src := getArgs()
	if 0 == quality {
		return
	}
	mat, err := imghelper.IMRead(src)
	if nil == err {
		err = imghelper.IMWrite(mat, dst, quality, 684)
	}
	panic(err)
}
