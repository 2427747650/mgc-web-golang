package mgc

import(
	"fmt"
	"tool/db_config"
	"os"
	"bufio"
	"strings"
)

//创建生成任务
func MGC_CREATE_TASK(){
	CreateEntity()
	CreateService()
}

//创建数据库实体
func CreateEntity(){
	entityDic:=db_config.GetEntityDirectory()
	result:=db_config.SelectNoParam("SELECT table_name tableName FROM  information_schema.tables WHERE TABLE_SCHEMA = '"+db_config.DataBase()+"'");
	for _, v := range result {
		tableName:=string(v["tableName"])
		tableName=FirstUpper(tableName)
		ClearTxt(entityDic+"\\"+tableName+".go")
		WriteTxt(entityDic+"\\"+tableName+".go","package entity\n")
		WriteTxt(entityDic+"\\"+tableName+".go","type "+tableName+" struct{\n")
		tableResult:=db_config.SelectNoParam("select column_name,column_type,COLUMN_COMMENT from information_schema.columns where table_name = '"+tableName+"' and table_schema='"+db_config.DataBase()+"'")
		for _,v1:=range tableResult {
			name:=string(v1["column_name"])
			name=FirstUpper(name)
			columnType:=string(v1["column_type"])
			sType:=""
			if strings.Index(columnType,"int")!=-1 {
				sType="int"
			}else if strings.Index(columnType,"varchar")!=-1 {
				sType="string"
			}else if strings.Index(columnType,"longtext")!=-1 {
				sType="string"
			}else if strings.Index(columnType,"text")!=-1 {
				sType="string"
			}else if strings.Index(columnType,"decimal")!=-1 {
				sType="float64"
			}else{
				sType="string"
			}
			WriteTxt(entityDic+"\\"+tableName+".go","     //"+string(v1["COLUMN_COMMENT"])+"\n")
			WriteTxt(entityDic+"\\"+tableName+".go","     "+name+" "+sType+"\n")
			
		}
		WriteTxt(entityDic+"\\"+tableName+".go","}")
		fmt.Println(tableName+"创建完毕，路径："+entityDic+"\\"+tableName+".go")
	}
	fmt.Println("MGC_CREATE_ENTITY任务执行完毕！")
}

//创建数据库接口
func CreateService(){
	interfaceDic:=db_config.GetInterfaceDirectory()
	result:=db_config.SelectNoParam("SELECT table_name tableName FROM  information_schema.tables WHERE TABLE_SCHEMA = '"+db_config.DataBase()+"'");
	for _, v := range result {
		tableName:=string(v["tableName"])
		tableName=FirstUpper(tableName)
		ClearTxt(interfaceDic+"\\"+tableName+"_interface.go")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","package go_interface\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","import \"tool/db_config\"\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","import \"entity\"\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","import \"strconv\"\n")
		//新增
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","//新增\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","func Insert_"+tableName+"(mod interface{}){\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     db_config.Insert(mod)\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","}\n\n")

		//编辑
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","//更新\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","func Update_"+tableName+"(num int64,mod interface{}){\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     db_config.Update(num,mod)\n")
		WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","}\n\n")

		

		tableResult:=db_config.SelectNoParam("select column_name,column_type,COLUMN_COMMENT from information_schema.columns where table_name = '"+tableName+"' and table_schema='"+db_config.DataBase()+"'")
		code:=0
		for _,v1:=range tableResult {
			if(code==0){
				//删除
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","//删除\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","func Delete_"+tableName+"(num int64){\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     param := map[string]interface{}{\"id\": num}\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     db_config.DeleteSql(\"delete from "+tableName+" where "+string(v1["column_name"])+"=?id\",param)\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","}\n\n")

				//查询指定ID数据
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","//查询指定ID的数据\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","func Select_"+tableName+"_by_id(num int64) entity."+tableName+"{\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     param := map[string]interface{}{\"id\": num}\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     dbResult:=db_config.Select(\"select * from "+tableName+" where "+string(v1["column_name"])+"=?id\",param)\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     result:=new(entity."+tableName+")\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     for _,v1:=range dbResult {\n")
				for _,v2:=range tableResult {
					columnType:=string(v2["column_type"])
					if strings.Index(columnType,"int")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_value:=string(v1[\""+string(v2["column_name"])+"\"])\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_i, _ := strconv.Atoi("+string(v2["column_name"])+"_value)\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"="+string(v2["column_name"])+"_i\n")
					}else if strings.Index(columnType,"varchar")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}else if strings.Index(columnType,"longtext")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}else if strings.Index(columnType,"text")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}else if strings.Index(columnType,"decimal")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_value:=string(v1[\""+string(v2["column_name"])+"\"])\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_floatvalue, _:= strconv.ParseFloat("+string(v2["column_name"])+"_value, 32)\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"="+string(v2["column_name"])+"_floatvalue\n")
					}else{
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}
				}
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     }\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     return *result\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","}\n\n")

				//查询列表
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","//查询列表\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","func Select_"+tableName+"() []entity."+tableName+"{\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     dbResult:=db_config.SelectNoParam(\"select * from "+tableName+"\")\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     var dt []entity."+tableName+"\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     for _,v1:=range dbResult {\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result:=new(entity."+tableName+")\n")
				for _,v2:=range tableResult {
					columnType:=string(v2["column_type"])
					if strings.Index(columnType,"int")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_value:=string(v1[\""+string(v2["column_name"])+"\"])\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_i, _ := strconv.Atoi("+string(v2["column_name"])+"_value)\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"="+string(v2["column_name"])+"_i\n")
					}else if strings.Index(columnType,"varchar")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}else if strings.Index(columnType,"longtext")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}else if strings.Index(columnType,"text")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}else if strings.Index(columnType,"decimal")!=-1 {
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_value:=string(v1[\""+string(v2["column_name"])+"\"])\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          "+string(v2["column_name"])+"_floatvalue, _:= strconv.ParseFloat("+string(v2["column_name"])+"_value, 32)\n")
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"="+string(v2["column_name"])+"_floatvalue\n")
					}else{
						WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          result."+FirstUpper(string(v2["column_name"]))+"=string(v1[\""+string(v2["column_name"])+"\"])\n")
					}
					
				}
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","          dt=append(dt,*result)\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     }\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","     return dt\n")
				WriteTxt(interfaceDic+"\\"+tableName+"_interface.go","}")

			}
			code++
			
		}
		fmt.Println(tableName+"_interface.go创建完毕，路径："+interfaceDic+"\\"+tableName+".go")
	}
	fmt.Println("MGC_CREATE_INTERFACE任务执行完毕！")
}

//写入文件
func WriteTxt(url string,str string){
	path := url
	file, err := os.OpenFile(path, os.O_APPEND | os.O_CREATE, 0644)
	if err!=nil{
		fmt.Println("file open error:",err)
		return
	}
	defer file.Close()
    //使用缓存方式写入
    writer := bufio.NewWriter(file)
 
    _, w_err := writer.WriteString(str)
	if w_err != nil {
        fmt.Println("写入出错")
    }
	//将缓存中数据刷新到文本中
    writer.Flush()
}

//清理文件内容
func ClearTxt(url string){
	os.OpenFile(url, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {

    if s == "" {

        return ""

    }

    return strings.ToUpper(s[:1]) + s[1:]

}