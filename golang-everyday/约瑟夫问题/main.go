package main

import "fmt"

type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个结点 默认值nil
}

// 编写一个函数,构成单向的环形链表
// num 小孩个数
// *Boy  返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  //空节点
	curBoy := &Boy{} // 空节点
	if num < 1 {
		fmt.Println("num不能为1")
		return first
	}
	// 循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		// 分析构成循环链表,需要一个辅助指针[帮忙]
		// 1. 因为第一个小孩比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first // 形成了循环
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //形成循环
		}
	}
	return first
}

// 显示单向的环形链表
func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	curBoy := first
	for {
		fmt.Println(curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

/*
	约瑟夫问题: 1,2...n,你个人围着坐,约定编号为k (1<=k<=n)
	的人从1开始计数,数到m的那个人出列,他的下一位又从1开始报数.数到m那个人出列,知道所有人出列位置
	1. playGame(first *Boy,startNo int,countNum int)
	2. 按照算法要求,留下最后一人
*/

func PlayGame(first *Boy, startNo int, countNum int) {
	//1. 空的链表我们单独处理
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}
	// 留一个,判断startNo <= 小孩的总数
	// 2. 定义辅助指针,帮助删除小孩
	tail := first
	//3. 让tail指向环形链表的最后一个小孩 因为tail在删除链表时需要使用
	for {
		if tail.Next == first {
			// 说明到了最后的小孩
			break
		}
		tail = tail.Next
	}
	// 4. 让first移动到 startNo [后面我们删除小孩,就以first为准]
	for i := 1; i < startNo; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 5. 开始数countNum下 然后删除first指向的节点
	for {
		// 开始数countNum -1 下,因为自身也要数一下
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("出圈的编号为%d\n", first.No)
		// 删除first指向的节点
		first = first.Next
		tail.Next = first

		// 判断如果tail == first 圈子中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出圈的小孩:%d\n", first.No)
}

func main() {
	first := AddBoy(5)
	//ShowBoy(first)
	PlayGame(first, 2, 3)
}
