package pool

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

type TestingData struct {
	Data string `json:"data"`
	Key  string `json:"key"`
}

func BenchmarkReadStreamWithPool(b *testing.B) {
	data := TestingData{
		Data: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.",
		Key:  "Lorem",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf := GetBuffer()

		_ = json.NewEncoder(buf).Encode(&data)
		io.Copy(ioutil.Discard, buf)

		PutBuffer(buf)
	}
}

func BenchmarkReadStreamWithoutPool(b *testing.B) {
	data := TestingData{
		Data: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.",
		Key:  "Lorem",
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(&data)
		io.Copy(ioutil.Discard, buf)
	}
}
