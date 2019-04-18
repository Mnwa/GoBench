package pool

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"sync"
	"testing"
)

var writerBufioPool = sync.Pool{
	New: func() interface{} {
		return bufio.NewWriter(ioutil.Discard)
	},
}

func GetBufioWriter() *bufio.Writer {
	return writerBufioPool.Get().(*bufio.Writer)
}

func PutBufioWriter(buf *bufio.Writer) {
	buf.Flush()
	writerBufioPool.Put(buf)
}

func BenchmarkWriteBufioWithPool(b *testing.B)  {
	data := bytes.NewReader([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie."))
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		writer := GetBufioWriter()
		io.Copy(writer, data)
		PutBufioWriter(writer)
	}
}

func BenchmarkWriteBufioWithoutPool(b *testing.B)  {
	data := bytes.NewReader([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie."))
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		writer := bufio.NewWriter(ioutil.Discard)
		io.Copy(writer, data)
	}
}