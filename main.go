package main

import (
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
	"DataCertPlatform/blockchain"
	"fmt"
)

func main() {
	//1创世区块
	bc := blockchain.NewBlockChain() //封装
	fmt.Printf("创世区块的哈希值:%x\n", bc.LastHash)
	//bc.SaveData([]byte("用户要保存到区块中的数据"))
	return

	//连接数据库
	db_mysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run() //阻塞
	//http.ListenAndServe(":8080")
}
