package pool

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"sync"
	"testing"
)

var writerGzipPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(ioutil.Discard)
	},
}

func GetGzipWriter() *gzip.Writer {
	return writerGzipPool.Get().(*gzip.Writer)
}

func PutGzipWriter(buf *gzip.Writer) {
	buf.Flush()
	writerGzipPool.Put(buf)
}

func BenchmarkWriteGzipWithPool(b *testing.B)  {
	data := bytes.NewReader([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie."))
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		writer := GetGzipWriter()
		io.Copy(writer, data)
		PutGzipWriter(writer)
	}
}

func BenchmarkWriteGzipWithoutPool(b *testing.B)  {
	data := bytes.NewReader([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie."))
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		writer := gzip.NewWriter(ioutil.Discard)
		io.Copy(writer, data)
	}
}