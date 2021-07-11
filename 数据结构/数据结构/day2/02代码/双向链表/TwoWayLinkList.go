package main

import "fmt"

type LinkNode struct {
	Data interface{} //数据
	Prev *LinkNode   //上一个指针
	Next *LinkNode   //下一个指针
}

//创建双向链表  （数据集合）
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
		//创建新节点
		newNode := new(LinkNode)
		newNode.Data = v
		//新节点指向上一个节点
		newNode.Prev = node
		//当前节点的下一个节点是新节点
		node.Next = newNode
		//当前节点为下一个节点
		node = node.Next
	}

	//头节点的上一个节点Prev 可以指向最后一个节点
	//head.Prev=node
	node = head

}

//打印双向链表
func (node *LinkNode) Print() {
	if node == nil {
		return
	}
	//正序打印数据
	for node != nil {
		if node.Data != nil {
			fmt.Println(node.Data)
		}
		node = node.Next
	}
}

//打印双向链表 倒序
func (node *LinkNode) Print02() {
	if node == nil {
		return
	}
	//指向链表末尾
	for node.Next != nil {
		node = node.Next
	}
	//从后向前打印数据
	for node.Prev != nil {
		if node.Data != nil {
			fmt.Println(node.Data)
		}
		node = node.Prev
	}
}

//数据长度 返回值 个数
func (node *LinkNode) Length() int {
	if node == nil {
		return -1
	}
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
	}
	return i
}

//插入数据 （下标 数据）
func (node *LinkNode) Index(index int, Data interface{}) {
	if node == nil {
		return
	}
	if index < 0 {
		return
	}
	if Data == nil {
		return
	}

	//记录上一个节点
	preNode := node
	//循环找到插入的节点
	for i := 0; i < index; i++ {
		preNode = node
		if node == nil {
			return
		}
		node = node.Next
	}

	//创建新节点
	newNode := new(LinkNode)
	newNode.Data = Data
	//将新节点的指针域分别指向上一个节点和下一个节点
	newNode.Next = node
	newNode.Prev = preNode

	//上一个节点的下一个节点为新节点
	preNode.Next = newNode
	//下一个节点的上一个节点为新节点
	node.Prev = newNode

}

//删除数据 （下标）
func (node *LinkNode) Delete(index int) {
	if node == nil {
		return
	}
	if index < 0 {
		return
	}

	//记录上一个节点
	preNode := node
	for i := 0; i < index; i++ {
		preNode = node
		if node == nil {
			return
		}

		node = node.Next
	}

	//删除节点
	preNode.Next = node.Next
	//当前节点的下一个节点的上一个节点为上一个节点
	node.Next.Prev = preNode

	//销毁当前节点
	node.Data = nil
	node.Next = nil
	node.Prev = nil
	node = nil
}

//链表销毁
func (node *LinkNode) Destroy() {
	if node == nil {
		return
	}

	node.Next.Destroy()

	node.Data = nil
	node.Next = nil
	node.Prev = nil
	node = nil
}
