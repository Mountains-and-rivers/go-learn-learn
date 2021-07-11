package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main(){
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db,_:=sql.Open("mysql","root:root@(127.0.0.1:3306)/itcast") // 设置连接数据库的参数
	defer db.Close()	//关闭数据库
	err:=db.Ping()		//连接数据库
	if err!=nil{
		fmt.Println("数据库连接失败")
		return
	}

	//操作一：执行数据操作语句
	/*
	sql:="insert into stu values (2,'berry')"
	result,_:=db.Exec(sql)		//执行SQL语句
	n,_:=result.RowsAffected();	//获取受影响的记录数
	fmt.Println("受影响的记录数是",n)
	*/

	//操作二：执行预处理
	/*
	stu:=[2][2] string{{"3","ketty"},{"4","rose"}}
	stmt,_:=db.Prepare("insert into stu values (?,?)")		//获取预处理语句对象
	for _,s:=range stu{
		stmt.Exec(s[0],s[1])			//调用预处理语句
	}
	*/

	//操作三：单行查询
	/*
	var id,name string
	rows:=db.QueryRow("select * from stu where id=4")   //获取一行数据
	rows.Scan(&id,&name)		//将rows中的数据存到id,name中
	fmt.Println(id,"--",name)
	*/

	//操作四：多行查询
	rows,_:=db.Query("select * from stu")		//获取所有数据
	var id,name string
	for rows.Next(){		//循环显示所有的数据
		rows.Scan(&id,&name)
		fmt.Println(id,"--",name)
	}
}
