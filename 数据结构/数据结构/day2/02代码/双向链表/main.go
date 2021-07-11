package main

import "fmt"

func main() {
	//创建链表
	list := new(LinkNode)
	list.Create(1, 2, 3, 4, 5)

	//打印链表
	//list.Print()
	//fmt.Println(list)
	//链表长度
	//count := list.Length()
	//fmt.Println(count)

	//插入数据
	//list.Index(2, 666)
	//删除数据
	list.Delete(2)
	list.Print()

	//销毁链表
	list.Destroy()
	fmt.Println("销毁后：")
	list.Print()
}
