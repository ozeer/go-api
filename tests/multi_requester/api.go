package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 定义一个结构体来表示请求的参数
type RequestBody1 struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

type RequestBody2 struct {
	Key3 string `json:"key3"`
	Key4 string `json:"key4"`
}

// 定义一个结构体来表示响应的结果
type ResponseBody struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	City    string `json:"city"`
}

func main() {
	// 创建一个 HTTP 服务器，监听端口 8080
	http.HandleFunc("/endpoint1", handleEndpoint1)
	http.HandleFunc("/endpoint2", handleEndpoint2)

	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

// 处理第一个 POST 请求的函数
func handleEndpoint1(w http.ResponseWriter, r *http.Request) {
	// 检查请求方法是否为 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体中的 JSON 数据
	var reqBody RequestBody1
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 延迟1秒
	time.Sleep(1 * time.Second)

	// 处理请求参数并生成响应
	message := fmt.Sprintf("Received request with key1=%s and key2=%s", reqBody.Key1, reqBody.Key2)
	response := ResponseBody{Message: message, Name: reqBody.Key1, City: reqBody.Key2}

	// 将响应转换为 JSON 并返回给客户端
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 处理第二个 POST 请求的函数
func handleEndpoint2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody2
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 延迟8秒
	time.Sleep(8 * time.Second)

	// 处理请求参数并生成响应
	message := fmt.Sprintf("Received request with key3=%s and key4=%s at endpoint2", reqBody.Key3, reqBody.Key4)
	response := ResponseBody{Message: message, Name: reqBody.Key3, City: reqBody.Key4}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
