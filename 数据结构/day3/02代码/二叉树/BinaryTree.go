package main

import (
	"fmt"
	"reflect"
)

type BinaryNode struct {
	Data   interface{} //数据
	lChild *BinaryNode //左子树
	rChild *BinaryNode //右子树
}

//创建二叉树
func (node *BinaryNode) Create() {
	node = new(BinaryNode)
}

//先序遍历
func (node *BinaryNode) PreOrder() {
	if node == nil {
		return
	}
	//DLR — 先序遍历，即先根再左再右
	fmt.Println(node.Data)

	//递归遍历左子树
	node.lChild.PreOrder()
	//递归遍历右子树
	node.rChild.PreOrder()

}

//中序遍历
func (node *BinaryNode) MidOrder() {
	if node == nil {
		return
	}

	//LDR — 中序遍历，即先左再根再右
	//递归遍历左子树
	node.lChild.MidOrder()
	//打印数据
	fmt.Println(node.Data)
	//递归遍历右子树
	node.rChild.MidOrder()

}

//后序遍历
func (node *BinaryNode) RearOrder() {
	if node == nil {
		return
	}

	//LRD — 后序遍历，即先左再右再根
	//递归遍历左子树
	node.lChild.RearOrder()
	//递归遍历右子树
	node.rChild.RearOrder()
	//打印数据
	fmt.Println(node.Data)
}

//二叉树高度 深度
func (node *BinaryNode) TreeHeight() int {
	if node == nil {
		return 0
	}
	//进入下一层遍历
	lh := node.lChild.TreeHeight()
	rh := node.rChild.TreeHeight()
	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}

}

//二叉树叶子节点个数
//叶子节点是没有后继的节点
func (node *BinaryNode) LeafCount(num *int) {
	if node == nil {
		return
	}
	//判断没有后继的节点为叶子节点
	if node.lChild == nil && node.rChild == nil {
		(*num)++
	}
	node.lChild.LeafCount(num)
	node.rChild.LeafCount(num)
}

//二叉树数据查找
func (node *BinaryNode) Search(Data interface{}) {
	if node == nil {
		return
	}

	//== 不支持slice 和 map
	//reflect.DeepEqual()
	if reflect.TypeOf(node.Data) == reflect.TypeOf(Data) && node.Data == Data {
		fmt.Println("找到数据", node.Data)
		return
	}
	node.lChild.Search(Data)
	node.rChild.Search(Data)
}

//二叉树销毁
func (node *BinaryNode) Destroy() {
	if node == nil {
		return
	}

	node.lChild.Destroy()
	node.lChild = nil
	node.rChild.Destroy()
	node.rChild = nil
	node.Data = nil

}

//二叉树反转
//如果想反转二叉树 二叉树必须是一个满二叉树
func (node *BinaryNode) Reverse() {
	if node == nil {
		return
	}

	//交换节点  多重赋值
	node.lChild, node.rChild = node.rChild, node.lChild

	node.lChild.Reverse()
	node.rChild.Reverse()

}

//二叉树拷贝
func (node *BinaryNode) Copy() *BinaryNode {
	if node == nil {
		return nil
	}
	lChild:=node.lChild.Copy()
	rChild:=node.rChild.Copy()

	//创建写节点 拷贝数据
	newNode:=new(BinaryNode)
	newNode.Data=node.Data
	newNode.lChild=lChild
	newNode.rChild=rChild
	return newNode
}
