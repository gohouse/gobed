package util

import (
	"bytes"
	"fmt"
	"github.com/gohouse/t"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

func UploadMultipart(apiurl, targetFile string, values map[string]io.Reader) (resp *http.Response, err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		//if x, ok := r.(*os.File); ok {
		//	if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
		//		return
		//	}
		//} else {
		//	// Add other fields
		//	if fw, err = w.CreateFormField(key); err != nil {
		//		return
		//	}
		//}
		if fw, err = w.CreateFormFile(key, t.New(time.Now().UnixNano()).String()); err != nil {
			return
		}
		if _, err = io.Copy(fw, r); err != nil {
			return
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", apiurl, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	var client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", resp.Status)
	}

	return
}
