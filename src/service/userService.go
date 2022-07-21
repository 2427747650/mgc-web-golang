package service

import "net/http"
import "tool/http_server"

//这是一个SERVICE方法
func HelloGolang(w http.ResponseWriter, r *http.Request){
	http_server.Cross(w,r);
	//操作数据库可直接调用go_interface包里面通过MGC自动生成的接口
	
}