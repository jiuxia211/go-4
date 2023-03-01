package main

import (
	"fmt"
	"jiuxia/crowdfunding/conf"
	"jiuxia/crowdfunding/model"
	"jiuxia/crowdfunding/routes"
)

func main() {
	fmt.Println("正在加载配置文件")
	conf.Init() //加载配置文件
	fmt.Println("正在连接数据库")

	model.Database() //连接数据库
	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}
