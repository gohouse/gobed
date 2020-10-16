package main

import (
	"fmt"
	"github.com/gohouse/gobed"
)

func main() {
	//var fileurl = "http://telegra.ph/file/6a5b15e7eb4d7329ca7af.jpg"
	//remoteUrl := gobed.Getter(gobed.BucketTelegraph).UploadFromUrl(fileurl).ResultSimple()
	//gb := gobed.Default().UploadFromUrl(fileurl)
	gb := gobed.Default().UploadFromLocalFile("demo.png")
	if gb.Error() != nil {
		fmt.Printf("err:%s", gb.Error())
		return
	}
	fmt.Printf("err:%s, 上传地址为: %s\n", gb.Error(), gb.ResultSimple())
}