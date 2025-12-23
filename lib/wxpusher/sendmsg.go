package wxpusher

import (
	"bytes"
	"encoding/json"
	"github.com/goravel/framework/facades"
	"io"
	"net/http"
	"time"
)

type SendTongzhiParams struct {
	AppToken    string `json:"appToken"`
	Content     string `json:"content"`
	Summary     string `json:"summary"`
	ContentType int    `json:"contentType"`
	TopicIds    []int  `json:"topicIds"`
	Url         string `json:"url"`
	VerifyPay   int    `json:"verifyPayType"`
}

type TongzhiResp struct {
	Code int64 `json:"code,omitempty"`
}

func SendMsg(p *SendTongzhiParams) bool {
	client := &http.Client{}
	jsonValue, _ := json.Marshal(p)
	requestBody := bytes.NewBuffer(jsonValue)
	req, err := client.Post("https://wxpusher.zjiecode.com/api/send/message", "application/json", requestBody) // 对于JSON
	defer req.Body.Close()
	if err != nil {
		return false
	}

	body, _ := io.ReadAll(req.Body)
	var tonzhiResp TongzhiResp
	err = json.Unmarshal(body, &tonzhiResp)
	if err != nil || tonzhiResp.Code != 1000 {
		facades.Log().Infof("通知结果异常", tonzhiResp.Code, string(body))
		return false
	}

	return true
}

// SendWorkWechat 发送企业微信机器人消息
func SendWorkWechat(content string) bool {
	webhookUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=12227931-a683-49b7-a338-fa3ffb9a1663"
	
	// 构建消息数据 - 使用 markdown 格式
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": content,
		},
	}
	
	bytesData, err := json.Marshal(data)
	if err != nil {
		facades.Log().Error("序列化企业微信消息失败:", err)
		return false
	}
	
	req, err := http.NewRequest("POST", webhookUrl, bytes.NewReader(bytesData))
	if err != nil {
		facades.Log().Error("创建企业微信请求失败:", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		facades.Log().Error("发送企业微信消息失败:", err)
		return false
	}
	defer resp.Body.Close()
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		facades.Log().Error("读取企业微信响应失败:", err)
		return false
	}
	
	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		facades.Log().Error("解析企业微信响应失败:", err)
		return false
	}
	
	// 检查是否发送成功 (errcode=0 表示成功)
	if errcode, ok := result["errcode"].(float64); ok && errcode != 0 {
		errmsg := result["errmsg"]
		facades.Log().Errorf("企业微信发送失败，errcode: %v, errmsg: %v", errcode, errmsg)
		return false
	}
	
	facades.Log().Info("企业微信消息发送成功:", string(bodyBytes))
	return true
}
