package main

import "fmt"

func main() {
	//创建队列
	queue := new(QueueNode)
	queue.Create(1, 2, 3, 4, 5)
	//打印队列
	queue.Print()
	//阶段队列个数
	count := queue.Length()
	fmt.Println("个数：", count)

	//入列
	queue.Push(666)
	queue.Print()

	//出列
	queue.Pop()
	queue.Print()
	queue.Pop()
	queue.Print()
}
