package protobuf

import (
	"bytes"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"testing"
)

func BenchmarkUnserializeToJson(b *testing.B) {
	data, _ := json.Marshal(&TestData{
		Message: "Lorem ipsum dolor sit amet",
		Data:    []string{"consectetur", "adipiscing", "elit", "Pellentesque"},
	})
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var result TestData
		json.Unmarshal(data, &result)
	}
}

func BenchmarkUnserializeToJsonStream(b *testing.B) {
	data, _ := json.Marshal(&TestData{
		Message: "Lorem ipsum dolor sit amet",
		Data:    []string{"consectetur", "adipiscing", "elit", "Pellentesque"},
	})
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var result TestData
		json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	}
}

func BenchmarkUnserializeToProtobuf(b *testing.B) {
	data, _ := proto.Marshal(&TestData{
		Message: "Lorem ipsum dolor sit amet",
		Data:    []string{"consectetur", "adipiscing", "elit", "Pellentesque"},
	})
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var result TestData
		proto.Unmarshal(data, &result)
	}
}
