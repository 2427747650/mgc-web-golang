package db_config

import "fmt"
import "github.com/xormplus/xorm"

//@Title XORM查询
//@sqlString SQL语句
//@param 参数map
func Select(sqlString string,param map[string]interface{})[]map[string]xorm.Value{
	var engine=SqlConnetion()
	results, err := engine.SQL(sqlString, &param).QueryValue()
	if(err!=nil){
		fmt.Println(err)
	}
	return results
}

//@Title XORM查询
//@sqlString SQL语句
func SelectNoParam(sqlString string)[]map[string]xorm.Value{
	var engine=SqlConnetion()
	results, err := engine.QueryValue(sqlString)
	if(err!=nil){
		fmt.Println(err)
	}
	return results
}

//@Title XORM插入
func Insert(mod interface{}) int64{
	var engine=SqlConnetion()
	affected, err := engine.Insert(mod)
	if(err!=nil){
		fmt.Println(err)
	}
	return affected
}

//@Title sql语句插入
//@sqlString SQL语句
//@param 参数map
func InsertSql(sqlString string,param map[string]interface{}){
	var engine=SqlConnetion()
	_, err := engine.SqlMapClient(sqlString, &param).Execute()
	if(err!=nil){
		fmt.Println(err)
	}
}

//@Title XORM更新
//@id ID
//@mod 实体
func Update(id int64,mod interface{}) int64{
	var engine=SqlConnetion()
	affected, err := engine.ID(id).Update(mod)
	if(err!=nil){
		fmt.Println(err)
	}
	return affected
}

//@Title sql语句更新
//@sqlString SQL语句
//@param 参数map
func UpdateSql(sqlString string,param map[string]interface{}){
	var engine=SqlConnetion()
	_, err := engine.SqlMapClient(sqlString, &param).Execute()
	if(err!=nil){
		fmt.Println(err)
	}
}

//@Title XORM删除
//@id ID
//@mod 实体
func Delete(id int64,mod interface{}) int64{
	var engine=SqlConnetion()
	affected, err := engine.ID(id).Delete(mod)
	if(err!=nil){
		fmt.Println(err)
	}
	return affected
}

//@Title sql语句删除
//@sqlString SQL语句
//@param 参数map
func DeleteSql(sqlString string,param map[string]interface{}){
	var engine=SqlConnetion()
	_, err := engine.SqlMapClient(sqlString, &param).Execute()
	if(err!=nil){
		fmt.Println(err)
	}
}