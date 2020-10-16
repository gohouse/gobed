package telegraph

import (
	"fmt"
	base "github.com/gohouse/gobed/base_upload"
	"github.com/gohouse/gobed/filetype"
	"github.com/gohouse/gobed/util"
	"github.com/gohouse/t"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var BucketName = "telegraph"

type Telegraph struct {
	*base.BaseUpload
	Api        string
	Host       string
	TargetFile string
	FileField  string
}
type TelegraphResult []struct {
	Src string `json:"src"`
}

func NewTelegraph() *Telegraph {
	var tg = &Telegraph{Api: "https://telegra.ph/upload", Host: "https://telegra.ph"}
	tg.BaseUpload =  base.NewBaseUpload(tg.UploadFromReader, tg.Result)
	return tg
}

func (tg *Telegraph) Result() *filetype.ResponseData {
	var res []string
	var trs TelegraphResult
	tg.Bind(&res)
	for _, v := range trs {
		res = append(res, fmt.Sprintf("%s%s", tg.Host, v.Src))
	}
	return &filetype.ResponseData{
		ServerName: BucketName,
		ServerHost: tg.Host,
		Files:      res,
		ExtraInfo:  nil,
	}
}

func (tg *Telegraph) UploadFromReader(r io.Reader) (iu filetype.IUpload) {
	targetFile := tg.TargetFile
	fileField := t.If(tg.FileField == "", "file", tg.FileField).String()
	var resp *http.Response
	resp, tg.Err = util.UploadMultipart(tg.Api, targetFile, map[string]io.Reader{fileField: r})
	if tg.Err != nil {
		return nil
	}
	defer resp.Body.Close()

	var all []byte
	all, tg.Err = ioutil.ReadAll(resp.Body)
	if tg.Err != nil {
		return nil
	}
	log.Printf("上传返回:%s\n", all)

	tg.TypeContext = t.New(all)
	return tg
}
