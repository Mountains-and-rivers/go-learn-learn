package main

//初始化堆
func HeapInit(arr []int) {

	//将切片转成二叉树模型  实现大根堆
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		HeapSort(arr, i, length-1)
	}

	//根节点存储最大值
	for i := length - 1; i > 0; i-- {
		//如果只剩下根节点和跟节点下的左子节点
		if i == 1 && arr[0] <= arr[i] {
			break
		}
		//将根节点和叶子节点数据交换
		arr[0], arr[i] = arr[i], arr[0]
		HeapSort(arr, 0, i-1)
	}

}

//获取堆中最大值  放在根节点
func HeapSort(arr []int, startNode int, maxNode int) {

	//最大值放在根节点
	var max int
	//定义做左子节点和右子节点
	lChild := startNode*2 + 1
	rChild := lChild + 1
	//子节点超过比较范围 跳出递归
	if lChild >= maxNode {
		return
	}
	//左右比较  找到最大值
	if rChild <= maxNode && arr[rChild] > arr[lChild] {
		max = rChild
	} else {
		max = lChild
	}

	//和跟节点比较
	if arr[max] <= arr[startNode] {
		return
	}

	//交换数据
	arr[startNode], arr[max] = arr[max], arr[startNode]
	//递归进行下次比较
	HeapSort(arr, max, maxNode)
}
