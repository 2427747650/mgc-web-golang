package main

import(
	"tool/db_config"
	"net/http"
	"controller"
)

func main(){
	db_config.RunMsg()
	
	//注册用户控制器
	controller.UserController()

	http.ListenAndServe(":8568", nil)

	//通过MGC创建数据库实体和对应接口
	//mgc.MGC_CREATE_TASK()

}