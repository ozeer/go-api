package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ozeer/go-api/model/common/response"
	"github.com/ozeer/go-api/utils"
	"github.com/redis/go-redis/v9"
)

func TestRegister(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	for i := 0; i < 1000000; i++ {
		phone, _ := rdb.SPop(context.Background(), "phone_pool").Result()
		// 创建一个包含JSON数据的map
		data := map[string]interface{}{
			// "username": "堂吉柯德",
			"username": uuid.New().String(),
			// "phone":    gofakeit.Phone(),
			"phone":    phone,
			"password": "123456zy",
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
		url := "http://localhost:8888/user/register"
		contentType := "application/json"
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("请求失败: ", err.Error())
			return
		}
		req.Header.Add("Content-Type", contentType)
		// req.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")
		req.Header.Add("x-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYWExODZhMGEtY2M0OC00NTNkLWI0ZGItNGVmMzUxNzA1MmUxIiwiSUQiOjEsIlVzZXJuYW1lIjoi5aCC5ZCJS0wiLCJOaWNrTmFtZSI6IiIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR08tQVBJIl0sImV4cCI6MTY5NDk3MzQzNCwibmJmIjoxNjk0MzY4NjM0fQ.jJwWSeAHItymbOPYAJij2dt8-4Iu2jyY9eE7Rs8zpqs")

		resp, err := client.Do(req)

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
