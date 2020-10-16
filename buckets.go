package gobed

import (
	"github.com/gohouse/gobed/filetype"
	"github.com/gohouse/gobed/telegraph"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	lock   sync.RWMutex
	Buckets = map[string]filetype.IUpload{}

	BucketTelegraph = telegraph.BucketName
)

func init()  {
	rand.Seed(time.Now().UnixNano())
	Register(BucketTelegraph, telegraph.NewTelegraph())
}

func Register(name string, iu filetype.IUpload) {
	lock.Lock()
	defer lock.Unlock()
	Buckets[name] = iu
}
func Getter(name string) filetype.IUpload {
	log.Printf("select bucket: %s\n", name)
	log.Println("Buckets:",Buckets)
	return Buckets[name]
}
func RandBucket() filetype.IUpload {
	if len(Buckets) > 0 {
		idx := rand.Intn(len(Buckets))
		var cur int
		for k, v := range Buckets {
			if idx == cur {
				log.Printf("select bucket: %s\n", k)
				return v
			}
			cur++
		}
	}
	return nil
}
