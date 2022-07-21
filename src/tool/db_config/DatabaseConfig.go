package db_config

import(
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/xormplus/xorm"
)

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

//entity目录
func GetEntityDirectory() string{
	return "D:\\project\\newProject\\2022\\java\\housePukeGame\\goServer\\GOPATH\\src\\entity"
}

//interface目录
func GetInterfaceDirectory() string{
	return "D:\\project\\newProject\\2022\\java\\housePukeGame\\goServer\\GOPATH\\src\\go_interface"
}


//连接数据库
func SqlConnetion() *xorm.Engine{
	var engine *xorm.Engine
	var err error
	server:=Address()
	database:=DataBase()
	uid:=Uid()
	pwd:=Pwd()
	engine, err = xorm.NewEngine("mysql", ""+uid+":"+pwd+"@tcp("+server+":3306)/"+database+"?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败！")
	}
	return engine
}