// 自动生成模板Openai
package openai

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Openai 结构体
type Openai struct {
	global.GVA_MODEL
	Tesst     string `json:"tesst" form:"tesst" gorm:"column:tesst;comment:测试;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Openai 表名
func (Openai) TableName() string {
	return "openai"
}

//Completion 一条Test请求与返回的记录表
type Completion struct {
	Model            string             `json:"model" bson:"model"`
	Prompt           string             `json:"prompt,omitempty" bson:"prompt"`
	Suffix           string             `json:"suffix,omitempty" bson:"suffix"`
	MaxTokens        int                `json:"max_tokens,omitempty" bson:"max_tokens"`
	Temperature      float32            `json:"temperature,omitempty" bson:"temperature"`
	TopP             float32            `json:"top_p,omitempty" bson:"top_p"`
	N                int                `json:"n,omitempty" bson:"n"`
	Stream           bool               `json:"stream,omitempty" bson:"stream"`
	LogProbs         int                `json:"logprobs,omitempty" bson:"logprobs"`
	Echo             bool               `json:"echo,omitempty" bson:"echo"`
	Stop             []string           `json:"stop,omitempty" bson:"stop"`
	PresencePenalty  float32            `json:"presence_penalty,omitempty" bson:"presence_penalty"`
	FrequencyPenalty float32            `json:"frequency_penalty,omitempty" bson:"frequency_penalty"`
	BestOf           int                `json:"best_of,omitempty" bson:"best_of"`
	LogitBias        map[string]int     `json:"logit_bias,omitempty" bson:"logit_bias"`
	User             string             `json:"user,omitempty" bson:"user"` //这里是客户的username
	ID               string             `json:"id" bson:"id"`               //这里是返回的ID,与mongodb的id不同,要区别
	Object           string             `json:"object" bson:"object"`
	Created          int64              `json:"created" bson:"created"`
	Choices          []CompletionChoice `json:"choices" bson:"choices"`
	Usage            Usage              `json:"usage" bson:"usage"`
	CreateTimeAt     time.Time          `json:"createTimeAt" bson:"createTimeAt,omitempty"`
	UpdateTimeAt     time.Time          `json:"updateTimeAt" bson:"updateTimeAt,omitempty"`
}

type CompletionRequest struct {
	Model            string         `json:"model"`
	Prompt           string         `json:"prompt,omitempty"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	LogProbs         int            `json:"logprobs,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int64              `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
	Usage   Usage              `json:"usage"`
}

type CompletionChoice struct {
	Text         string        `json:"text" bson:"text"`
	Index        int           `json:"index" bson:"index"`
	FinishReason string        `json:"finish_reason" bson:"finish_reason"`
	LogProbs     LogprobResult `json:"logprobs" bson:"logprobs"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens" bson:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens" bson:"completion_tokens"`
	TotalTokens      int `json:"total_tokens" bson:"total_tokens"`
}

type LogprobResult struct {
	Tokens        []string             `json:"tokens" bson:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs" bson:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs" bson:"top_logprobs"`
	TextOffset    []int                `json:"text_offset" bson:"text_offset"`
}
