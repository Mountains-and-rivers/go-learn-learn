package main

func QuickSort(arr []int, left int, right int) {
	//设置基准值
	temp := arr[left]
	index := left

	i := left
	j := right

	for i <= j {
		//从右面找到比基准值小的数据
		for j >= index && arr[j] >= temp {
			j--
		}
		//获取基准值合适下标
		if j > index {
			arr[index] = arr[j]
			index = j
		}
		//从左面找比基准值大的数据
		for i <= index && arr[i] <= temp {
			i++
		}
		//获取基准值合适下标
		if i <= index {
			arr[index] = arr[i]
			index = i
		}
	}
	//将基准值放在合适位置
	arr[index] = temp

	//递归调用 分步处理数据
	if index-left > 1 {
		QuickSort(arr, left, index-1)
	}
	if right-index > 1 {
		QuickSort(arr, index+1, right)
	}

}
