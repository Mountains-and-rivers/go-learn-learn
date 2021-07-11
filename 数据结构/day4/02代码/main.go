package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main01() {
	arr := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}

	//冒泡排序
	//BubbleSort(arr)

	//选择排序
	//SelectSort(arr)

	//插入排序
	//InsertSort(arr)

	//希尔排序
	//ShellSort(arr)

	//快速排序
	//QuickSort(arr, 0, len(arr)-1)

	//堆排序
	HeapInit(arr)

	fmt.Println(arr)
}

//变相排序  基于大量重复 在某一个范围内
func main02() {
	//随机数种子
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 0)

	for i := 0; i < 10000; i++ {
		s = append(s, rand.Intn(1000)) //0-999
	}
	fmt.Println(s)

	//统计数据集合中数据出现的次数
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	//fmt.Println(m)

	//排序
	for i := 0; i < 1000; i++ {
		for j := 0; j < m[i]; j++ {
			fmt.Print(i, " ")
		}
	}

}
