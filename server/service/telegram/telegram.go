package telegram

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/openai"
	"github.com/go-resty/resty/v2"
)

type TelegramService struct {
}

//发送信息
func (telegramService *TelegramService) SendMessage(message openai.SendMessageToUserStruct) {

	// url := "https://api.telegram.org/bot6250735595:AAF8GxRVi0wSY734CLoJd-3GbT2Yof-gdOA/sendMessage"
	url := "https://api.telegram.org/bot" + global.TELEGRAM_TOKEN + "/sendMessage"
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(message).
		Post(url)
	if err != nil {
		fmt.Printf("SendMessage出错了:err: %v\n", err)
	}
	// s := string(resp.Body())
	// fmt.Printf("发信息后返回的数据:s: %v\n", s)
}

//接受telegram服务器传过来的命令命令,处理命令
func (telegramService *TelegramService) CommonDo(command string, telegramBotUpdate openai.TelegramBotUpdate, args ...string) {
	var sendMessageInfo openai.SendMessageToUserStruct
	if command == "/ask" {
		fmt.Println(telegramBotUpdate.Message.Text)
		sendMessageInfo.ChatId = telegramBotUpdate.Message.Chat.Id
		sendMessageInfo.Text = "回复了哟"
		sendMessageInfo.ParseMode = "HTML"
		telegramService.SendMessage(sendMessageInfo)
	}

	//帮助信息
	if command == "/help" {
		// var count int64 = 100
		// db := global.GVA_DB.Model(&wallpaper)
		// db.Count(&count)
		tempStr := `欢迎使用神医的oeoChat机器人,此机器人对接openAi官方API。
		使用方法:直接发送文字问题,免费额度有限，请珍惜使用.
		按文字生成图片的功能正在开发中
`
		sendMessageInfo.ChatId = telegramBotUpdate.Message.Chat.Id
		// sendMessageInfo.Text = fmt.Sprintf(tempStr)
		sendMessageInfo.Text = tempStr
		sendMessageInfo.ParseMode = "HTML"
		telegramService.SendMessage(sendMessageInfo)
	}
	//帮助信息
	if command == "/start" {
		sendMessageInfo.ChatId = telegramBotUpdate.Message.Chat.Id
		sendMessageInfo.Text = `欢迎使用神医的oeoChat机器人,此机器人对接openAi官方API。
		使用方法: 直接发送文字问题,免费额度有限，请珍惜使用.
		按文字生成图片的功能正在开发中
		`
		sendMessageInfo.ParseMode = "HTML"
		telegramService.SendMessage(sendMessageInfo)
	}

}
