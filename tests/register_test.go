package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ozeer/go-api/model/common/response"
	"github.com/ozeer/go-api/utils"
)

func TestRegister(t *testing.T) {
	for i := 0; i < 200000; i++ {
		// 创建一个包含JSON数据的map
		data := map[string]interface{}{
			// "username": "堂吉柯德",
			"phone":    gofakeit.Phone(),
			"password": "12345",
			"email":    gofakeit.Email(),
			"sex":      rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3),
			"birthday": utils.GenerateRandomDate(1949, 2023),
		}

		// 将map转换为JSON格式
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("JSON编码错误:", err)
			return
		}
		// var jsonData = []byte(`{
		// 	"username": "堂吉柯德",
		// 	"phone": "18701588472",
		// 	"password": "12345",
		// 	"email": "zhouyang2021@163.com",
		// 	"sex": 1,
		// 	"birthday": "2000-09-18"
		// }`)
		resp, err := http.Post("http://localhost:8888/user/register", "application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			log.Println("注册失败: ", err.Error())
		}

		defer resp.Body.Close()

		// 检查HTTP响应状态码
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("请求失败，状态码：%d\n", resp.StatusCode)
			return
		}

		// 读取响应体
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取响应体发生错误:", err)
			return
		}

		var decodeData response.Response
		if err := json.Unmarshal(respBody, &decodeData); err != nil {
			fmt.Println("解析JSON发生错误:", err)
			return
		}

		log.Printf("#%d#注册返回数据: %s", i, decodeData.Msg)
	}
}
