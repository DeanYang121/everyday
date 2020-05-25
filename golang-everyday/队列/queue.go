package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 模拟队列
	front   int    // 表示指向队列首部
	rear    int    // 表示指向队列的尾部
}

// 添加数据到队列
func (this *Queue) AddQueue(value int) error {
	// 先判断队列是否已满
	if this.rear == this.maxSize - 1{
		return errors.New("队列已满")
	}

	// rear 后移
	this.rear++
	this.array[this.rear] = value
	return nil
}

// 显示队列  找到对首 然后遍历到队尾
func(this *Queue) ShowQueue(){
	fmt.Println("队列当前的情况是:")
	// this.front 不包含队首的元素
	for i := this.front + 1;i < this.rear;i++{
		fmt.Printf("array[%d]=%d\t",i,this.array[i])
	}
}

// 取出队列的数据
func(this *Queue) GetQueue()(int,error){
	// 判断队列是否为空
	if this.front == this.rear{
		return -128,errors.New("queue is empty!")
	}
	this.front++
	val := this.array[this.front]
	return val,nil
}


func main() {
	// 创建一个队列
	queue := Queue{
		maxSize:5,
		front: -1,
		rear: -1,
	}

	var key string
	var val int
	for{
		fmt.Println("1. 输入add 添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 显示队列")
		fmt.Println("4. 输入exit 退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入对的数据")
			fmt.Scanln(&val)
			if err := queue.AddQueue(val);err != nil{
				fmt.Println(err)
			}else{
				fmt.Println("加入队列成功!")
			}
		case "get":
			if val,err := queue.GetQueue();err != nil{
				fmt.Println(err)
			}else{
				fmt.Printf("取出数据:%d \n",val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			fmt.Println("退出成功!")
			os.Exit(0)
		}
	}
}


