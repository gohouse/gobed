package base

import (
	"bytes"
	"github.com/gohouse/gobed/filetype"
	"github.com/gohouse/t"
	"io"
	"net/http"
	"os"
)

type UpReaderHand func(r io.Reader) filetype.IUpload
type ResultHand func() *filetype.ResponseData
type BaseUpload struct {
	t.TypeContext
	Err          error
	TargetFile   string
	FileField    string
	UpReaderHand UpReaderHand
	ResultHand ResultHand
}

func NewBaseUpload(urh UpReaderHand, rh ResultHand) *BaseUpload {
	return &BaseUpload{UpReaderHand: urh, ResultHand: rh}
}

func (bu *BaseUpload) SetRemoteFileName(fileName string) filetype.IUpload {
	return nil
}
func (bu *BaseUpload) UploadFromBytes(fileBytes *[]byte) filetype.IUpload {
	return bu.UploadFromReader(bytes.NewBuffer(*fileBytes))
}
func (bu *BaseUpload) UploadFromLocalFile(localFile string) filetype.IUpload {
	open, err := os.Open(localFile)
	if err != nil {
		bu.Err = err
		return nil
	}
	return bu.UploadFromReader(open)
}
func (bu *BaseUpload) UploadFromLocalDirectory(localDirectory string) filetype.IUpload {
	return nil
}
func (bu *BaseUpload) UploadFromUrl(fileUrl string) filetype.IUpload {
	get, err := http.Get(fileUrl)
	if err != nil {
		bu.Err = err
		return bu
	}
	defer get.Body.Close()

	return bu.UploadFromReader(get.Body)
}
func (bu *BaseUpload) History(args ...filetype.HistoryOption) filetype.IUpload {
	return nil
}
func (bu *BaseUpload) Delete(delParam string) filetype.IUpload {
	return nil
}
func (bu *BaseUpload) Error() error {
	return bu.Err
}

func (bu *BaseUpload) Result() *filetype.ResponseData {
	return bu.ResultHand()
}

func (bu *BaseUpload) ResultSimpleMulti() (res []string) {
	return bu.Result().Files
}

func (bu *BaseUpload) ResultSimple() string {
	res := bu.ResultSimpleMulti()
	if len(res) > 0 {
		return res[0]
	}
	return ""
}

func (bu *BaseUpload) Response() t.Type {
	return bu.TypeContext
}

func (bu *BaseUpload) UploadFromReader(r io.Reader) filetype.IUpload {
	return bu.UpReaderHand(r)
}
