package client_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func processResponse(resp *http.Response) {
	// 处理响应
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status code: %d", resp.StatusCode)
	} else {
		fmt.Printf("status code: %d\n", resp.StatusCode)
		// 这里必须要关闭响应体，否则会导致内存泄漏
		defer resp.Body.Close()
		// 打印响应体
		fmt.Println("响应体：")
		all, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ReadAll failed: %v", err)
		}
		fmt.Println(string(all))
	}
}

func Get(url string) {
	fmt.Println("Get:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get failed: %v", err)
	}
	processResponse(resp)

}

func TestClient(t *testing.T) {
	url := "http://127.0.0.1:8080/home"
	Get(url)

}
