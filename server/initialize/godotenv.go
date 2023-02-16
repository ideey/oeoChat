package initialize

import (
	"fmt"

	"github.com/joho/godotenv"
)

//加载.env配置文件
func GodotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("载入.env 配置文件出错")
	}
}
