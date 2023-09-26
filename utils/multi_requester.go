package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/ozeer/go-api/global"
	"go.uber.org/zap"
)

type MultiRequester struct {
	client             *http.Client
	requestPool        []*http.Request
	requestBox         map[*http.Request]map[string]interface{}
	requestID          string
	defaultInnerDomain string
	mu                 sync.Mutex
	responseCache      map[*http.Request]map[string]interface{}
}

func NewMultiRequester(timeout time.Duration) *MultiRequester {
	mr := &MultiRequester{
		client:             &http.Client{},
		requestPool:        make([]*http.Request, 0),
		requestBox:         make(map[*http.Request]map[string]interface{}),
		requestID:          "0",
		defaultInnerDomain: "",
		responseCache:      make(map[*http.Request]map[string]interface{}), // 初始化响应缓存
	}

	if timeout.Seconds() > 0 {
		mr.client.Timeout = timeout
	}

	return mr
}

func (mr *MultiRequester) getDefaultDomain() string {
	if mr.defaultInnerDomain == "" {
		// Implement the logic to get the default inner domain here
		// You can use a configuration file or a constant value
		// Example:
		// mr.defaultInnerDomain = "https://example.com"
	}
	return mr.defaultInnerDomain
}

func (mr *MultiRequester) Add(url string, params map[string]interface{}) *http.Request {
	commonParams := mr.getCommonParams()
	for k, v := range params {
		commonParams[k] = v
	}

	jsonParams, _ := json.Marshal(commonParams)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonParams))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "your_user_agent")

	mr.requestPool = append(mr.requestPool, req)
	mr.requestBox[req] = map[string]interface{}{"url": url, "params": params}

	return req
}

func (mr *MultiRequester) Exec() {
	var wg sync.WaitGroup

	for _, req := range mr.requestPool {
		wg.Add(1)
		go func(req *http.Request) {
			defer wg.Done()
			resp, err := mr.doRequest(req)

			if err != nil {
				// 处理请求错误
				global.LOG.Error("Client executing request error", zap.Error(err))
				return
			}

			// 将响应缓存到 responseCache 中
			mr.mu.Lock()
			mr.responseCache[req] = resp
			mr.mu.Unlock()
		}(req)
	}

	wg.Wait()
}

func (mr *MultiRequester) doRequest(req *http.Request) (map[string]interface{}, error) {
	resp, err := mr.client.Do(req)
	if err != nil {
		// Handle the error
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle the error
		global.LOG.Error("Client read resp error", zap.Error(err))

		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// 处理非 200 响应
		return nil, errors.New("Non-200 HTTP status code")
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		global.LOG.Error("Client unmarshal error", zap.Error(err))

		return nil, err
	}

	return responseMap, nil
}

func (mr *MultiRequester) getCommonParams() map[string]interface{} {
	commonParams := map[string]interface{}{
		"inner_request": 1,
		"request_id":    mr.requestID,
	}
	return commonParams
}

func (mr *MultiRequester) GetContent(req *http.Request) map[string]interface{} {
	// Implement logic to extract content from the response
	// You can use a similar approach as in the "doRequest" method
	// Return the content as a map[string]interface{}
	mr.mu.Lock()
	defer mr.mu.Unlock()

	// 从 responseCache 中获取响应
	if resp, ok := mr.responseCache[req]; ok {
		return resp
	}

	// 如果请求不存在于 responseCache 中，返回 nil
	return nil
}

func (mr *MultiRequester) Cleanup() {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	for req := range mr.requestBox {
		delete(mr.requestBox, req)
	}
}

// You can add additional methods and functionality as needed
