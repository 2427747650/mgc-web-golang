package controller

import "tool/router"
import "service"

func UserController(){
	//定义服务
	router.Router("/helloGolang",service.HelloGolang)
}