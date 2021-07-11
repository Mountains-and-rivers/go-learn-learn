package main

import "fmt"

func main() {
	//创建树
	tree := new(BinaryNode)

	//创建节点
	node1 := BinaryNode{1, nil, nil}
	node2 := BinaryNode{2, nil, nil}
	node3 := BinaryNode{3, nil, nil}
	node4 := BinaryNode{4, nil, nil}
	node5 := BinaryNode{5, nil, nil}
	node6 := BinaryNode{6, nil, nil}
	//node7 := BinaryNode{7, nil, nil}

	//创建树
	tree.Data = 0
	tree.lChild = &node1
	tree.rChild = &node2
	node1.lChild = &node3
	node1.rChild = &node4
	node2.lChild = &node5
	node2.rChild = &node6
	//node3.lChild = &node7

	//先序遍历
	//tree.PreOrder()
	//中序遍历
	//tree.MidOrder()
	//后序遍历
	//tree.RearOrder()

	//二叉树高度 深度
	//h := tree.TreeHeight()
	//fmt.Println("高度：", h)

	//二叉树叶子节点
	//num := 0
	//tree.LeafCount(&num)
	//fmt.Println("叶子节点：",num)

	//二叉树数据查找
	//tree.Search(6)

	//销毁二叉树
	//tree.Destroy()
	//tree.RearOrder()

	//fmt.Println("翻转前：")
	//tree.PreOrder()
	////二叉树反转
	//tree.Reverse()
	//
	//fmt.Println("翻转后：")
	//tree.PreOrder()

	//拷贝二叉树
	tree1 := tree.Copy()

	tree1.lChild.Data=666

	fmt.Println("原始树：")
	tree.PreOrder()
	fmt.Println("新树：")
	tree1.PreOrder()
}
