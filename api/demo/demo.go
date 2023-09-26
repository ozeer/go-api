package demo

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/model/common/response"
	"github.com/ozeer/go-api/utils"
)

type DemoApi struct{}

type RespMultiRequest struct {
	Data1 any   `json:"data1"`
	Data2 any   `json:"data2"`
	Cost  int64 `json:"cost"`
}

func (d *DemoApi) MultiRequest(c *gin.Context) {
	begin := time.Now()

	// 创建一个 MultiRequester 实例，设置超时时间为2秒
	mr := utils.NewMultiRequester(2 * time.Second)

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
	// fmt.Println("Response from endpoint 1:", content1)

	// 获取请求2的响应内容
	content2 := mr.GetContent(req2)
	// fmt.Println("Response from endpoint 2:", content2)

	// 清理资源
	mr.Cleanup()

	response.OkWithDetailed(RespMultiRequest{
		Data1: content1,
		Data2: content2,
		Cost:  time.Since(begin).Milliseconds(),
	}, "返回成功", c)
}
