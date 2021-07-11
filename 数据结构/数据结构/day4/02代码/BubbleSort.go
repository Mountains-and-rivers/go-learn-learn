package main

func BubbleSort(arr []int) {
	flag := false
	//外层控制行
	for i := 0; i < len(arr)-1; i++ {
		//内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			//比较两个相邻元素
			if arr[j] > arr[j+1] {
				//交换数据
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		//判断数据是否是有序
		if !flag {
			return
		} else {
			flag = false
		}
	}
}
