package main

import "fmt"

func SelectSort(arr *[5]int){

	for j := 0;j<len(arr)-1;j++{
		max := arr[j]
		maxIndex := j
		for i := j+1;i<len(arr);i++{
			if max < arr[i]{// 找最大值
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != 0{
			arr[j],arr[maxIndex] = arr[maxIndex],arr[j]
		}
		fmt.Printf("第%d次: %v \n",j,*arr)
	}
}

func main() {
	arr := [5]int{10,34,19,100,80}
	SelectSort(&arr)
	fmt.Println(arr)
}
