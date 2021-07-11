package main

import "fmt"

func main() {

	//创建链栈
	s := CreateStack(1, 2, 3, 4, 5)
	//打印链栈
	//PrintStack(s)

	//打印链栈个数
	//count:=LengthStack(s)
	//fmt.Println("个数：",count)
	//fmt.Println(s)
	//入栈
	s = Push(s, 6)
	s = Push(s, 7)
	s = Push(s, 8)
	PrintStack(s)
	//出栈
	s=Pop(s)
	fmt.Println("\n出栈后：")
	PrintStack(s)

	//出栈
	s=Pop(s)
	fmt.Println("\n出栈后：")
	PrintStack(s)

	//出栈
	s=Pop(s)
	fmt.Println("\n出栈后：")
	PrintStack(s)

	//清空链栈
	s=Clear(s)
	fmt.Println("\n清空栈：")
	PrintStack(s)

}
