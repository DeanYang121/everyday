package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [4]int
	head    int
	tail    int // 初始化 head tail 都是0
}

// 入队列
func (this *CircleQueue) PushQueue(val int) error {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	// 分析出this.tail 在队列尾部,但是不包含最后的元素
	this.array[this.tail] = val //把值给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return nil
}

// 出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return -128, errors.New("queue is empty")
	}
	// 取出 head 队首 且包含队首元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, nil
}

// 显示队列
func (this *CircleQueue) Show() {
	fmt.Println("环形队列情况如下:")
	// 取出当前队列有多少元素
	size := this.Size()
	if size == 0 {
		fmt.Println("当前队列为空")
		return
	}

	// 设计一个辅助变量,指向head
	tempHead := this.head
	for i := this.head; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 判断环形队列满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断环形队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.head == this.tail
}

// 取出环形队列有多少元素
func (this *CircleQueue) Size() int {
	// 这是一个关键的算法.
	return (this.tail + this.maxSize - 1 - this.head) % this.maxSize
}

func main() {
	// 创建一个队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 显示队列")
		fmt.Println("4. 输入exit 退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入对的数据")
			fmt.Scanln(&val)
			if err := queue.PushQueue(val); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入队列成功!")
			}
		case "get":
			if val, err := queue.Pop(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("取出数据:%d \n", val)
			}
		case "show":
			queue.Show()
		case "exit":
			fmt.Println("退出成功!")
			os.Exit(0)
		}
	}
}
