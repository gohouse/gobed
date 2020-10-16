package filetype

import (
	"github.com/gohouse/t"
	"io"
)

type HistoryOption struct {
	Limit int
	Page  int
}
type ResponseData struct {
	ServerName string
	ServerHost string
	Files      []string
	ExtraInfo  interface{}
}
type IUpload interface {
	t.Type
	// SetRemoteFileName set the file name for server
	// ex: /img/2020/e10abc23sdfsadfsfsd23sddfs.jpg
	SetRemoteFileName(fileName string) IUpload
	UploadFromBytes(fileBytes *[]byte) IUpload
	UploadFromLocalFile(localFile string) IUpload
	UploadFromLocalDirectory(localDirectory string) IUpload
	UploadFromUrl(fileUrl string) IUpload
	History(args ...HistoryOption) IUpload
	Delete(delParam string) IUpload
	Error() error
	Result() *ResponseData
	ResultSimpleMulti() (res []string)
	ResultSimple() string
	Response() t.Type
	UploadFromReader(r io.Reader) IUpload
}