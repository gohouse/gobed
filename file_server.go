package gobed

import (
	"github.com/gohouse/gobed/filetype"
)

func Default() filetype.IUpload {
	fi := Getter(BucketTelegraph)
	return fi
}
