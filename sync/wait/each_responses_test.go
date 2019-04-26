package wait

import (
	"encoding/json"
	"net/http"
	"runtime"
	"sync"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(4)
}

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"string"`
}

func BenchmarkEachResponses(b *testing.B) {
	var posts []Post
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var requestedPosts []Post
		res, err := http.Get("https://my-json-server.typicode.com/typicode/demo/posts")
		if err != nil {
			b.Error(err)
		} else {
			err = json.NewDecoder(res.Body).Decode(&requestedPosts)
			posts = append(posts, requestedPosts...)
			if err != nil {
				b.Error(err)
			}
		}
	}
}

func BenchmarkParallelEachResponses(b *testing.B) {
	var posts []Post
	wg := new(sync.WaitGroup)
	mx := new(sync.Mutex)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		wg.Add(1)
		go func(n int) {
			var requestedPosts []Post
			res, err := http.Get("https://my-json-server.typicode.com/typicode/demo/posts")
			if err != nil {
				b.Error(err)
			} else {
				err = json.NewDecoder(res.Body).Decode(&requestedPosts)
				mx.Lock()
				posts = append(posts, requestedPosts...)
				mx.Unlock()
				if err != nil {
					b.Error(err)
				}
			}
			wg.Done()
		}(n)
	}
	wg.Wait()
}
