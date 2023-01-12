package main

import (
	"LAST_PROJECT/config"
	"LAST_PROJECT/log"
	"LAST_PROJECT/model"
)

func main() {
	if err := config.Init("", "LAST_PROJECT"); err != nil {
		panic(err)
	}

	model.DB.Init()
	defer model.DB.Close()

	defer log.SyncLogger()

	// 开始执行脚本
	//handler.StartScript()
}
