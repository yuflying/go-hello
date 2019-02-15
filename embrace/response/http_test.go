package response

import (
	"fmt"
	"net/http"
	"testing"
)

// go test -count=1 -v *.go -test.run TestResponseJson
func TestResponseJson(t *testing.T) {
	http.HandleFunc("/t1", func(w http.ResponseWriter, r *http.Request) {
		HTTPJson(w, 0, "success", map[string]interface{}{"name": "flyu", "age": 17})
	})
	http.HandleFunc("/t2", func(w http.ResponseWriter, r *http.Request) {
		HTTPJson(w, 1, "failed", nil)
	})
	fmt.Println("开始监听80端口服务...")
	http.ListenAndServe(":80", nil)
}

// go test -count=1 -v *.go -test.run TestResponseJsonWithPage
func TestResponseJsonWithPage(t *testing.T) {
	http.HandleFunc("/t1", func(w http.ResponseWriter, r *http.Request) {
		HTTPJsonWithPage(w, 0, "success", 1, 10, 1000, map[string]interface{}{"name": "flyu", "age": 17})
	})
	http.HandleFunc("/t2", func(w http.ResponseWriter, r *http.Request) {
		HTTPJsonWithPage(w, 1, "failed", 1, 10, 1000, nil)
	})
	fmt.Println("开始监听80端口服务...")
	http.ListenAndServe(":80", nil)
}
