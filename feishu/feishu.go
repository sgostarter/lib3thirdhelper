package feishu

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type TextContent struct {
	Text string `json:"text"`
}

type fsTextMessage struct {
	MsgType string       `json:"msg_type"`
	Content *TextContent `json:"content"`
}

func SendTextNotify(token, text string) (result string, err error) {
	msg := &fsTextMessage{
		MsgType: "text",
		Content: &TextContent{
			Text: text,
		},
	}

	d, err := json.Marshal(msg)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", "https://open.feishu.cn/open-apis/bot/v2/hook/"+token,
		bytes.NewReader(d))
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	d, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	result = string(d)

	return
}
