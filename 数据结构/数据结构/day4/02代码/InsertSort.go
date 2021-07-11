package main

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		//如果当前数据小于有序数据
		if arr[i] < arr[i-1] {
			j := i - 1
			//获取有效数据
			temp := arr[i]
			//一次比较有序数据
			for j >= 0 && arr[j] > temp {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = temp
		}
	}
}
