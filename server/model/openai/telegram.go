package openai

type TelegramBotUpdate struct {
	UpdateId int64   `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int64       `json:"message_id"`
	From      MessageFrom `json:"from"`
	Chat      Chat        `json:"chat"`
	Date      int64       `json:"date"`
	Text      string      `json:"text,omitempty"`
	Photos    []PhotoOne  `json:"photo,omitempty"`
	Entities  []Entitie   `json:"entities,omitempty"`
	Document  *Document   `json:"document,omitempty"`
}
type MessageFrom struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
}

type PhotoOne struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
}

type Entitie struct {
	Offset int64  `json:"offset"`
	Length int64  `json:"length"`
	Type   string `json:"type"`
}

type GetFileInfo struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileId       string `json:"file_id"`
		FilePath     string `json:"file_path"`
		FileSize     int64  `json:"file_size"`
		FileUniqueId string `json:"file_unique_id"`
	} `json:"result"`
}

type Document struct {
	FileName     string `json:"file_name"`
	MimeType     string `json:"mime_type"`
	Thumb        *Thumb `json:"thumb,omitempty"`
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
}

type Thumb struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
}

type SendMessageToUserStruct struct {
	ChatId    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type SendPhotoS struct {
	ChatId   int64  `json:"chat_id"`
	Photo    string `json:"photo,omitempty"`
	Document string `json:"document,omitempty"`
}
