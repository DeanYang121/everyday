package main

import "fmt"

/*
	1. left 表示数组左边的下标
	2. right 表示数组右边的下标
    3. array 表示要排序的数组
*/
func QuickSort(left int,right int,array *[6]int){
	l := left
	r := right
	privot := array[(left+right)/2]
	temp := 0
	// 将比privot小的放到左边
	for ;l<r;{
		// 先从privot左边 找到大于等于privot的值
		for ;array[l] < privot;{
			l++
		}
		for ;array[r] > privot;{
			r--
		}
		// 两头分割查找的相等 表示本次查找任务完成
		if l >= r{
			break
		}
		// 进行交换 把左边和右边找出来的两个数交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		// 相等不交换 直接往下走一步
		if array[l] == privot{
			r--
		}
		if array[r] == privot{
			l++
		}
	}
	// 如果l == r 再移动一位
	if l == r{
		l++
		r--
	}
	// 递归
	if left < r{
		QuickSort(left,r,array)
	}
	if right > l{
		QuickSort(l,right,array)
	}
}

func main() {
	arr := [6]int{-9,78,0,23,-234,98}
	QuickSort(0,len(arr)-1,&arr)
	fmt.Println(arr)
}
