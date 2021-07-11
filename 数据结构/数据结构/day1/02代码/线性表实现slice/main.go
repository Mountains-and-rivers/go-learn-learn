package main

import "fmt"

func main() {

	var s Slice
	//创建切片
	s.Create(5, 5, 1, 2, 3, 4, 5)
	s.Append(6, 7, 8)
	s.Append(9,10,11)
	//打印切片
	s.Print()

	fmt.Println("\n长度：", s.len)
	fmt.Println("容量：", s.cap)

	//根据下标获取元素
	value:=s.GetdData(11)
	fmt.Println("值：",value)

	//根据元素获取下标
	index:=s.Search(666)
	fmt.Println("下标：",index)

	//删除元素
	s.Delete(5)
	s.Delete(5)
	s.Print()
	fmt.Println()

	//指针偏移
	s.Insert(8,666)
	s.Print()

	//slice销毁
	fmt.Println(s)
	s.Destroy()
	fmt.Println(s)

	//s1:=[]int{1,2,3,4,5}
	////将一个切片指向地址的值设置为nil  就会被gc回收
	//s1=nil
}
