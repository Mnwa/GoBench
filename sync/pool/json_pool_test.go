package pool

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

type TestingResponseData struct {
	Data string `json:"data"`
	Key  string `json:"key"`
}

var responsePool = sync.Pool{
	New: func() interface{} {
		return new(TestingResponseData)
	},
}

func GetResponse() *TestingResponseData {
	return responsePool.Get().(*TestingResponseData)
}

func PutResponse(buf *TestingResponseData) {
	responsePool.Put(buf)
}

func BenchmarkJsonDecodeWithPool(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := bytes.NewReader([]byte("{\"data\":\"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.\",\"key\":\"Lorem\"}"))
		response := GetResponse()

		_ = json.NewDecoder(data).Decode(response)

		PutResponse(response)
	}
}

func BenchmarkJsonDecodeWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := bytes.NewReader([]byte("{\"data\":\"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie.\",\"key\":\"Lorem\"}"))
		response := new(TestingResponseData)

		_ = json.NewDecoder(data).Decode(response)
	}
}
