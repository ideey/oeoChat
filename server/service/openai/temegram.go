package openai

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/openai"
	"github.com/go-resty/resty/v2"
)

type TelegramService struct {
}

//发送信息
func (telegramService *TelegramService) SendMessage(message openai.SendMessageToUserStruct) {

	url := "https://api.telegram.org/bot5736458056:AAGaEhg33OwjCRFQXzyIQjfvN4dtkNmr8B4/sendMessage"
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(message).
		Post(url)
	if err != nil {
		fmt.Printf("SendMessage出错了:err: %v\n", err)
	}
	s := string(resp.Body())
	fmt.Printf("发信息后返回的数据:s: %v\n", s)
}
