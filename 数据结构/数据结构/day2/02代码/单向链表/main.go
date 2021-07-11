package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  string
	age  int
}

func main() {

	s1 := Student{1001, "张三", "男", 18}
	s2 := Student{1008, "李四", "男", 20}
	s3 := Student{1010, "王五", "男", 24}
	s4 := Student{1007, "赵六", "男", 19}
	s5 := Student{1666, "刘七", "男", 21}

	//s6 := Student{1000, "陈八", "女", 30}
	//s6 := Student{1008, "李四", "男", 20}

	var list *LinkNode = new(LinkNode)
	//list.Create(1,2,3,4,5)
	//创建链表
	list.Create(s1, s2, s3, s4, s5)
	//打印链表
	//list.Print()

	//计算链表个数
	//count := list.Length()
	//fmt.Println(count)

	//数据插入（头插）
	//list.InsertByHead(s6)
	//数据插入（尾插）
	//list.InsertByTail(s6)
	//数据插入 （位置）
	//list.InserrByIndex(2, s6)
	//数据删除 （位置）
	//list.DeleteByIndex(2)
	//list.DeleteByIndex(2)

	//删除数据 （数据）
	//list.DeleteByData(s6)

	//查找数据
	//index:=list.Search(s6)
	//fmt.Println(index)
	list.Print()

	//销毁链表
	list.Destroy()
	fmt.Println("销毁后：")
	list.Print()
}
