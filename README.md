# mgc-web-golang
基于golang的web开发框架，内置mysql数据库实体类和interface接口自动生成的工具，开发人员只需要开发服务即可

#1.下载程序，将程序放入GOPATH目录，添加此目录环境变量

#2添加xorm
go get github.com/xormplus/xorm

#3添加mysql驱动
go get github.com/go-sql-driver/mysql

#4配置数据库信息 tool>db_config>DatabaseConfig.go
//数据库地址
func Address() string{
	return "localhost"
}

//数据库名称
func DataBase() string{
	return "database"
}

//数据库用户名
func Uid() string{
	return "root"
}

//数据库密码
func Pwd() string{
	return "password"
}

//entity目录(使用MGC生成实体类目录)
func GetEntityDirectory() string{
	return "D:\\project\\newProject\\2022\\java\\housePukeGame\\goServer\\GOPATH\\src\\entity"
}

//interface目录（使用MGC生成接口文件目录）
func GetInterfaceDirectory() string{
	return "D:\\project\\newProject\\2022\\java\\housePukeGame\\goServer\\GOPATH\\src\\go_interface"
}

#5生成实体类和接口文件 main.go
func main(){
//特别注意，数据库表有变更再执行一下调用，生成完毕请注释此调用
	mgc.MGC_CREATE_TASK()//调用此方法即在配置的entity和interface目录生成对应的实体类struct和接口（接口包含新增、删除、编辑、查询、ID查询五个基础接口，其他接口需要开发者自行编写）
}

#6创建控制器（请在controller中创建路由控制器，如创建user用户控制器就在controller创建userController.go，代码如下）
package controller

import "tool/router"
import "service"

func UserController(){
	//定义服务
	router.Router("/helloGolang",service.HelloGolang)
}

#7创建Service实现（请在service中创建对应的service，如userService.go，代码如下）
package service

import "net/http"
import "tool/http_server"

//这是一个SERVICE方法
func HelloGolang(w http.ResponseWriter, r *http.Request){
	http_server.Cross(w,r);
	//操作数据库可直接调用go_interface包里面通过MGC自动生成的接口
	
}

#8 main.go中注册控制器和启动监听端口，代码如下
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

#9 MGC封装orm调用如下
引用包 import "tool/db_config"

（1）通过sql查询  db_config.Select(sqlString string,param map[string]interface{})[]map[string]xorm.Value
（2）ORM新增 db_config.Insert(mod interface{}) int64
（3）SQL语句新增 db_config.InsertSql(sqlString string,param map[string]interface{})
（4）ORM编辑 db_config.Update(id int64,mod interface{}) int64
（5）SQL语句编辑 db_config.UpdateSql(sqlString string,param map[string]interface{})
（6）ORM删除 db_config.Delete(id int64,mod interface{}) int64
（7）SQL语句删除 db_config.DeleteSql(sqlString string,param map[string]interface{})


欢迎各位开发者积极提供反馈和BUG，我一个人能力有限，会在收到各位同僚的反馈情况下，尽可能更新修改，谢谢大家支持！
