package wxpusher

import (
	"bytes"
	"encoding/json"
	"github.com/goravel/framework/facades"
	"io"
	"net/http"
)

type SendTongzhiParams struct {
	AppToken    string `json:"appToken"`
	Content     string `json:"content"`
	Summary     string `json:"summary"`
	ContentType int    `json:"contentType"`
	TopicIds    []int  `json:"topicIds"`
	Url         string `json:"url"`
	VerifyPay   bool   `json:"verifyPay"`
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
