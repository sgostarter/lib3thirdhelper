package feishu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TextContent struct {
	Text string `json:"text"`
}

type fsTextMessage struct {
	MsgType string       `json:"msg_type"`
	Content *TextContent `json:"content"`
}

type fsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func SendTextNotify(token, text string) (errMessage string, err error) {
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

	var fsResp fsResponse

	err = json.NewDecoder(resp.Body).Decode(&fsResp)
	if err != nil {
		return
	}

	if fsResp.Code != 0 {
		err = fmt.Errorf("code:%d", fsResp.Code)

		errMessage = fsResp.Msg

		return
	}

	return
}
