package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestMultiRequester(t *testing.T) {
	begin := time.Now()

	// 创建一个 MultiRequester 实例，设置超时时间为5秒
	mr := NewMultiRequester(5 * time.Second)

	// 添加请求到 MultiRequester
	params1 := map[string]interface{}{
		"key1": "中国",
		"key2": "美国",
	}
	req1 := mr.Add("http://localhost:8080/endpoint1", params1)

	params2 := map[string]interface{}{
		"key3": "唐朝",
		"key4": "宋朝",
	}
	req2 := mr.Add("http://localhost:8080/endpoint2", params2)

	// 执行所有请求
	mr.Exec()

	// 获取请求1的响应内容
	content1 := mr.GetContent(req1)
	fmt.Println("Response from endpoint 1:", content1)

	// 获取请求2的响应内容
	content2 := mr.GetContent(req2)
	fmt.Println("Response from endpoint 2:", content2)

	// 清理资源
	mr.Cleanup()

	fmt.Printf("\ncost: %dms\n", time.Since(begin).Milliseconds())
}
