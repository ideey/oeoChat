package initialize

import (
	"context"
	"fmt"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/qiniu/qmgo"
)

func Mongodb() {
	// const uri = "mongodb://t.deey.top:57890/?maxPoolSize=20&w=majority"
	ctx := context.Background()
	var err error
	global.QmgoClient, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: os.Getenv("MONGODB_URI")})
	if err != nil {
		panic(err)
	}
	global.QmgoDatabase = global.QmgoClient.Database("oeochat")
	global.QmgoCollCompletion = global.QmgoDatabase.Collection("completion")
	fmt.Println("初始化mongodb成功")
}
