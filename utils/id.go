package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SnowFlake struct {
	Id        uint64 `json:"id"`
	MachineId uint16 `json:"machine-id"`
	Msb       uint8  `json:"msb"`
	Sequence  uint32 `json:"sequence"`
	Time      uint64 `json:"time"`
}

func GenUid() (id uint64) {
	response, err := http.Get("http://localhost:8080")

	if err != nil {
		log.Println("snowflake server error: ", err.Error())
		return
	}

	defer response.Body.Close()

	// 检查HTTP响应状态码
	if response.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码：%d\n", response.StatusCode)
		return
	}

	// 读取响应体
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应体发生错误:", err)
		return
	}

	// 解析JSON响应数据
	// var snowflake map[string]interface{}
	var snowflake SnowFlake
	if err := json.Unmarshal(responseBody, &snowflake); err != nil {
		fmt.Println("解析JSON发生错误:", err)
		return
	}

	return snowflake.Id
}
