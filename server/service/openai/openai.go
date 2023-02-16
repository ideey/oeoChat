package openai

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/openai"
	openaiReq "github.com/flipped-aurora/gin-vue-admin/server/model/openai/request"
	gogpt "github.com/sashabaranov/go-gpt3"
	"gorm.io/gorm"
)

type OpenaiService struct {
}

// CreateOpenai 创建Openai记录
// Author [piexlmax](https://github.com/piexlmax)
func (myOpenaiService *OpenaiService) CreateOpenai(myOpenai openai.Openai) (err error) {
	err = global.GVA_DB.Create(&myOpenai).Error
	return err
}

// DeleteOpenai 删除Openai记录
// Author [piexlmax](https://github.com/piexlmax)
func (myOpenaiService *OpenaiService) DeleteOpenai(myOpenai openai.Openai) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openai.Openai{}).Where("id = ?", myOpenai.ID).Update("deleted_by", myOpenai.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&myOpenai).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteOpenaiByIds 批量删除Openai记录
// Author [piexlmax](https://github.com/piexlmax)
func (myOpenaiService *OpenaiService) DeleteOpenaiByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openai.Openai{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&openai.Openai{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateOpenai 更新Openai记录
// Author [piexlmax](https://github.com/piexlmax)
func (myOpenaiService *OpenaiService) UpdateOpenai(myOpenai openai.Openai) (err error) {
	err = global.GVA_DB.Save(&myOpenai).Error
	return err
}

// GetOpenai 根据id获取Openai记录
// Author [piexlmax](https://github.com/piexlmax)
func (myOpenaiService *OpenaiService) GetOpenai(id uint) (myOpenai openai.Openai, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&myOpenai).Error
	return
}

// GetOpenaiInfoList 分页获取Openai记录
// Author [piexlmax](https://github.com/piexlmax)
func (myOpenaiService *OpenaiService) GetOpenaiInfoList(info openaiReq.OpenaiSearch) (list []openai.Openai, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&openai.Openai{})
	var myOpenais []openai.Openai
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&myOpenais).Error
	return myOpenais, total, err
}

//=====以下为自定义接口服务============

// OpenaiCompletions openai的Completions接口服务
// Author [piexlmax](https://github.com/piexlmax)
func (myOpenaiService *OpenaiService) OpenaiCompletions(completionsRequest gogpt.CompletionRequest) (resp gogpt.CompletionResponse) {
	gpt3 := gogpt.NewClient(global.GVA_CONFIG.OpenAi.Token)
	ctx := context.Background()
	resp, err := gpt3.CreateCompletion(ctx, completionsRequest)
	if err != nil {
		fmt.Println("OPENAI completion请求出现错误.")
		return
	}
	return resp
}
