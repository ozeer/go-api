package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

type DataStruct struct {
	Id uint64 `json:"id"`
}

type Response struct {
	Code int        `json:"code"`
	Data DataStruct `json:"data"`
	Msg  string     `json:"msg"`
}

func GenUid() (uid uint64) {
	resp, err := http.Get("http://localhost:8888/system/genId")

	if err != nil {
		log.Println("snowflake server error: ", err.Error())
		return
	}

	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码：%d\n", resp.StatusCode)
		return
	}

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体发生错误:", err)
		return
	}

	// 解析JSON响应数据
	// var snowflake map[string]interface{}
	var snowflake Response
	if err := json.Unmarshal(responseBody, &snowflake); err != nil {
		fmt.Println("解析JSON发生错误:", err)
		return
	}

	return snowflake.Data.Id
}

func TestGenUid(t *testing.T) {
	log.Println(GenUid())
}
