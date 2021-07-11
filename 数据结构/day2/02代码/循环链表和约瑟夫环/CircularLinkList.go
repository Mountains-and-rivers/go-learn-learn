package main

import "fmt"

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

//创建循环链表
func (node *LinkNode) Create(Data ...interface{}) {
	if node == nil {
		return
	}
	if len(Data) == 0 {
		return
	}

	//记录头节点
	head := node

	for _, v := range Data {

		//创建一个新节点
		newNode := new(LinkNode)
		newNode.Data = v

		//当前节点的下一个节点为新节点
		node.Next = newNode
		node = node.Next
	}
	//形成闭环
	//最后一个节点的下一个节点 为第一个节点
	node.Next = head.Next
	node = head
}

//打印循环链表
func (node *LinkNode) Print() {
	if node == nil {
		return
	}

	//记录循环链表的起始点
	start := node.Next

	//打印链表
	for {
		node = node.Next
		if node.Data != nil {
			fmt.Print(node.Data, " ")

		}
		if start == node.Next {
			break
		}

	}
	//for node.Next != nil {
	//	if node.Data != nil {
	//		fmt.Println(node.Data)
	//	}
	//	node = node.Next
	//}
}

//循环链表长度 返回值 个数
func (node *LinkNode) Length() int {
	if node == nil {
		return -1
	}

	//记录循环链表的起始点
	start := node.Next

	//定义计数器
	i := 0
	for {
		node = node.Next
		i++
		if start == node.Next {
			break
		}
	}
	return i
}

//插入数据（下标 数据）
func (node *LinkNode) Insert(index int, Data interface{}) {
	if node == nil {
		return
	}
	//判断位置是否越界
	if index < 0 || node.Length()+1 < index {
		return
	}
	if Data == nil {
		return
	}
	fmt.Println(node)
	//记录链表第一个节点
	start := node.Next

	//记录上一个节点
	preNode := node

	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}

	//创建新节点
	newNode := new(LinkNode)
	newNode.Data = Data
	//新节点的下一个节点为node
	newNode.Next = node
	//上一个节点的下一个节点为新节点
	preNode.Next = newNode

	//判断如果是第一个节点需要将最后一个节点指向新节点
	if index == 1 {
		for {
			//判断最后一个节点
			if start == node.Next {
				break
			}
			node = node.Next
		}
		//新节点为最后一个节点的下一个节点
		node.Next = newNode
	}

}

//删除数据 （下标）
func (node *LinkNode) Delete(index int) {
	if node == nil {
		return
	}
	if index < 0{
		return
	}

	//起始节点
	start := node.Next

	//记录上一个节点
	preNode := node
	//循环找到删除数据的位置
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}


	//判断如果删除的是第一个节点
	if index == 1 {
		//temp记录最后一个位置
		temp := start
		for {

			if start == temp.Next {
				break
			}

			temp = temp.Next
		}
		//将最后一个节点指向新节点
		temp.Next = node.Next
	}

	//删除数据
	preNode.Next = node.Next

	//数据销毁
	//node.Data = nil
	//node.Next = nil
	node = nil
}
