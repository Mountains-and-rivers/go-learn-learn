package main

import "fmt"

//二分查找 BinarySearch(数据，元素) 返回值为下标
func BinarySearch(arr []int, num int) int {
	//定义起始下标
	start := 0
	//定义结束下标
	end := len(arr) - 1
	//中间基准值
	mid := (start + end) / 2

	for i := 0; i < len(arr); i++ {
		//基准值为查找值
		if num == arr[mid] {
			return mid
		} else if num > arr[mid] {
			//比基准值大  查找右侧
			start = mid + 1
		} else {
			//比基准值小  查找左侧
			end = mid - 1
		}
		//再次设置中间基准值位置
		mid = (start + end) / 2
	}
	return -1
}
func main() {
	//有序数据
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 666

	index := BinarySearch(arr, num)
	fmt.Println(index)
}
