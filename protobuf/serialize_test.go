package protobuf

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"testing"
)

func BenchmarkSerializeToJson(b *testing.B) {
	data := &TestData{
		Message: "Lorem ipsum dolor sit amet",
		Data:    []string{"consectetur", "adipiscing", "elit", "Pellentesque"},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		json.Marshal(data)
	}
}

func BenchmarkSerializeToJsonStream(b *testing.B) {
	data := &TestData{
		Message: "Lorem ipsum dolor sit amet",
		Data:    []string{"consectetur", "adipiscing", "elit", "Pellentesque"},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		json.NewEncoder(ioutil.Discard).Encode(data)
	}
}

func BenchmarkSerializeToProtobuf(b *testing.B) {
	data := &TestData{
		Message: "Lorem ipsum dolor sit amet",
		Data:    []string{"consectetur", "adipiscing", "elit", "Pellentesque"},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		proto.Marshal(data)
	}
}
