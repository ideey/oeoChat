package openai

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OpenaiRouter struct {
}

// InitOpenaiRouter 初始化 Openai 路由信息
func (s *OpenaiRouter) InitOpenaiRouter(Router *gin.RouterGroup, RouterPublic *gin.RouterGroup) {
	myOpenaiRouter := Router.Group("myOpenai").Use(middleware.OperationRecord())
	myOpenaiRouterWithoutRecord := Router.Group("myOpenai")
	myOpenaiRouterPublicWithoutRecord := RouterPublic.Group("myOpenai")
	var myOpenaiApi = v1.ApiGroupApp.OpenaiApiGroup.OpenaiApi
	{
		myOpenaiRouter.POST("createOpenai", myOpenaiApi.CreateOpenai)             // 新建Openai
		myOpenaiRouter.DELETE("deleteOpenai", myOpenaiApi.DeleteOpenai)           // 删除Openai
		myOpenaiRouter.DELETE("deleteOpenaiByIds", myOpenaiApi.DeleteOpenaiByIds) // 批量删除Openai
		myOpenaiRouter.PUT("updateOpenai", myOpenaiApi.UpdateOpenai)              // 更新Openai
	}
	{
		myOpenaiRouterWithoutRecord.GET("findOpenai", myOpenaiApi.FindOpenai)       // 根据ID获取Openai
		myOpenaiRouterWithoutRecord.GET("getOpenaiList", myOpenaiApi.GetOpenaiList) // 获取Openai列表
	}
	{
		myOpenaiRouterPublicWithoutRecord.POST("completions", myOpenaiApi.Completions) //openAi completions接口 接收请求，发送到openai得到答案，再返回给客户

		myOpenaiRouterPublicWithoutRecord.POST("telegrambot", myOpenaiApi.GetMessageFromTelegram) //Telegram机器人接口
	}
}
