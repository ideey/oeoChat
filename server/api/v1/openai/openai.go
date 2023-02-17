package openai

import (
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/openai"
	openaiReq "github.com/flipped-aurora/gin-vue-admin/server/model/openai/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"go.uber.org/zap"
)

type OpenaiApi struct {
}

var myOpenaiService = service.ServiceGroupApp.OpenaiServiceGroup.OpenaiService
var telegramService = service.ServiceGroupApp.OpenaiServiceGroup.TelegramService

// CreateOpenai 创建Openai
// @Tags Openai
// @Summary 创建Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openai.Openai true "创建Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myOpenai/createOpenai [post]
func (myOpenaiApi *OpenaiApi) CreateOpenai(c *gin.Context) {
	var myOpenai openai.Openai
	err := c.ShouldBindJSON(&myOpenai)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	myOpenai.CreatedBy = utils.GetUserID(c)
	if err := myOpenaiService.CreateOpenai(myOpenai); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOpenai 删除Openai
// @Tags Openai
// @Summary 删除Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openai.Openai true "删除Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /myOpenai/deleteOpenai [delete]
func (myOpenaiApi *OpenaiApi) DeleteOpenai(c *gin.Context) {
	var myOpenai openai.Openai
	err := c.ShouldBindJSON(&myOpenai)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	myOpenai.DeletedBy = utils.GetUserID(c)
	if err := myOpenaiService.DeleteOpenai(myOpenai); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOpenaiByIds 批量删除Openai
// @Tags Openai
// @Summary 批量删除Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /myOpenai/deleteOpenaiByIds [delete]
func (myOpenaiApi *OpenaiApi) DeleteOpenaiByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := myOpenaiService.DeleteOpenaiByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOpenai 更新Openai
// @Tags Openai
// @Summary 更新Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openai.Openai true "更新Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /myOpenai/updateOpenai [put]
func (myOpenaiApi *OpenaiApi) UpdateOpenai(c *gin.Context) {
	var myOpenai openai.Openai
	err := c.ShouldBindJSON(&myOpenai)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	myOpenai.UpdatedBy = utils.GetUserID(c)
	if err := myOpenaiService.UpdateOpenai(myOpenai); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOpenai 用id查询Openai
// @Tags Openai
// @Summary 用id查询Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openai.Openai true "用id查询Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /myOpenai/findOpenai [get]
func (myOpenaiApi *OpenaiApi) FindOpenai(c *gin.Context) {
	var myOpenai openai.Openai
	err := c.ShouldBindQuery(&myOpenai)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remyOpenai, err := myOpenaiService.GetOpenai(myOpenai.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remyOpenai": remyOpenai}, c)
	}
}

// GetOpenaiList 分页获取Openai列表
// @Tags Openai
// @Summary 分页获取Openai列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openaiReq.OpenaiSearch true "分页获取Openai列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myOpenai/getOpenaiList [get]
func (myOpenaiApi *OpenaiApi) GetOpenaiList(c *gin.Context) {
	var pageInfo openaiReq.OpenaiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := myOpenaiService.GetOpenaiInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

//openAi的Comletions接口,处理接入数据，以及判断收费情况
func (myOpenaiApi *OpenaiApi) Completions(c *gin.Context) {
	var completionsRequest gogpt.CompletionRequest
	err := c.ShouldBindJSON(&completionsRequest)
	if err != nil {
		fmt.Println("发生错误10")
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Printf("请求结构:%+v\n", completionsRequest)
	res, err := myOpenaiService.OpenaiCompletions(completionsRequest)
	if err != nil {
		fmt.Println("Completion请求出错:", err)
		response.Fail(c)
		return
	}
	fmt.Printf("返回的结构:%+v\n", res)
	response.OkWithData(res, c)
}

//接收telegram信息
func (myOpenaiApi *OpenaiApi) GetMessageFromTelegram(c *gin.Context) {
	fmt.Println("收到telegram服务器webHook消息:")
	response.Ok(c)
	var telegramBotUpdate openai.TelegramBotUpdate
	err := c.ShouldBindJSON(&telegramBotUpdate)
	if err != nil {
		fmt.Printf("绑定出错:err: %v\n", err)
		return
	}
	var sendMessageInfo openai.SendMessageToUserStruct
	//处理文字信息
	if telegramBotUpdate.Message.Text != "" {
		fmt.Println("文字信息")
		if telegramBotUpdate.Message.Entities != nil {
			for i, entitie := range telegramBotUpdate.Message.Entities {
				//目前只处理第一个命令   暂不处理别的形式
				if i > 0 {
					return
				}
				if entitie.Type == "bot_command" {
					//命令
					commandText := telegramBotUpdate.Message.Text
					command := commandText[entitie.Offset:(entitie.Offset + entitie.Length)]
					fmt.Printf("得到的命令是: %v\n", command)
					args := strings.Split(commandText, " ")
					telegramService.CommonDo(command, telegramBotUpdate, args[1:]...)
				}
			}
		} else { //问题处理环节 - 非命令
			ask := telegramBotUpdate.Message.Text
			user := telegramBotUpdate.Message.From.UserName
			//构建openai的 Completions  请求体
			var completionRequest gogpt.CompletionRequest
			var myCompletion openai.Completion
			completionRequest.Model = gogpt.GPT3TextDavinci003
			completionRequest.Temperature = 0.8
			completionRequest.Prompt = ask
			completionRequest.N = 1
			completionRequest.User = user
			completionRequest.MaxTokens = 500

			//组建插入mongdb的结构体
			myCompletion.Model = completionRequest.Model
			myCompletion.Temperature = completionRequest.Temperature
			myCompletion.Prompt = completionRequest.Prompt
			myCompletion.N = completionRequest.N
			myCompletion.User = completionRequest.User
			myCompletion.MaxTokens = completionRequest.MaxTokens
			myCompletion.Status = -1
			myCompletion.CreateTimeAt = time.Now()
			resp, err := myOpenaiService.OpenaiCompletions(completionRequest) //发送Completions 请求
			if err != nil {
				fmt.Println("Completion请求出错:", err)
				sendMessageInfo.ChatId = telegramBotUpdate.Message.Chat.Id
				sendMessageInfo.Text = "请求出错,请稍后再试"
				telegramService.SendMessage(sendMessageInfo)
				myCompletion.Status = 0 //失败记录

			} else {
				sendMessageInfo.ChatId = telegramBotUpdate.Message.Chat.Id
				sendMessageInfo.Text = resp.Choices[0].Text
				telegramService.SendMessage(sendMessageInfo)
				myCompletion.Status = 1 //成功记录
				myCompletion.ID = resp.ID
				myCompletion.Object = resp.Object
				myCompletion.Created = resp.Created
				myCompletion.Model = resp.Model
				if len(resp.Choices) >= 1 {
					for i, v := range resp.Choices {
						myCompletion.Choices[i].Text = v.Text
						myCompletion.Choices[i].Index = v.Index
						myCompletion.Choices[i].FinishReason = v.FinishReason
					}
				}
				myCompletion.Usage.PromptTokens = resp.Usage.PromptTokens
				myCompletion.Usage.CompletionTokens = resp.Usage.CompletionTokens
				myCompletion.Usage.TotalTokens = resp.Usage.TotalTokens
			}
			myCompletion.UpdateTimeAt = time.Now()
			myOpenaiService.SaveCompletins(myCompletion) //保存记录到mongodb
		}
	}
	//处理直接发图片的信息
	if telegramBotUpdate.Message.Photos != nil {
		fmt.Println("这里哟")
	}
	//处理以文件形式发送的图片(排除非图片):jpg jpeg png 5M以下
	if telegramBotUpdate.Message.Document != nil {
		fmt.Println("这里哟1")
	}
}
