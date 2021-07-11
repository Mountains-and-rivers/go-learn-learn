package main

func SelectSort(arr []int) {

	//外层控制行
	for i := 0; i < len(arr); i++ {
		//记录最大值下标
		index := 0
		//内层控制列
		//遍历数据 查找最大值
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] > arr[index] {
				//记录下标
				index = j
			}
		}

		//交换数据
		arr[index], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[index]
	}
}
