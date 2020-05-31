package main

import "fmt"

func InsertSort(arr *[5]int){

	for i := 1;i<len(arr);i++{
		// 一个有序列表  一个无序  把无序的一个个移入有序
		insertVal := arr[i]
		insertIndex := i -1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal{
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex-- // 往前查找
		}
		// 插入
		if insertIndex + 1 != i{
			arr[insertIndex + 1] = insertVal
		}
		fmt.Printf("第%d次插入之后:%v\n",i,*arr)
	}
}

func main() {
	arr := [5]int{23,0,12,56,34}
	InsertSort(&arr)
}
