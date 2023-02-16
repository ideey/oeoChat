package initialize

import (
	"fmt"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/joho/godotenv"
)

//加载.env配置文件 保存在全局变量中
func GodotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("载入.env 配置文件出错")
	} else {
		global.OPENAI_TOKEN = os.Getenv("OPENAI_TOKEN")
		global.MONGODB_URI = os.Getenv("MONGODB_URI")
		global.TELEGRAM_TOKEN = os.Getenv("TELEGRAM_TOKEN")
	}
}
