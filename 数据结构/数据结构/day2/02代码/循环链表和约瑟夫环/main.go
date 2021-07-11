package main

import "fmt"

func main01() {
	//创建循环链表
	list := new(LinkNode)
	list.Create(1, 2, 3, 4, 5)
	//fmt.Println(list)
	//打印循环链表
	//list.Print()

	//插入数据
	//list.Insert(1, 666)
	//list.Print()
	//计算循环链表长度
	//count := list.Length()
	//fmt.Println(count)

	//删除节点
	list.Delete(1)
	list.Print()
}

//约瑟夫环

func main() {
	list := new(LinkNode)
	list.Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32)

	fmt.Println("原始数据：")
	list.Print()
	fmt.Println()
	fmt.Println("删除数据：")
	i := 1
	for list.Length() > 2 {
		i += 3
		if i > list.Length() {
			i = list.Length()%3 + 1
		}
		list.Delete(i)
		list.Print()
		fmt.Println()
	}
}
