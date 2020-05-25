package main

import "fmt"

type CatNode struct {
	no   int //猫咪编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加的第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 环形,一只猫也是环形的  不能是newCatNode
		fmt.Println(head, "加入到环形的链表")
		return
	}

	//定义一个临时的变量,帮忙找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	// 加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下:")
	temp := head
	if temp.next == nil {
		fmt.Println("空链表....")
		return
	}

	for {
		fmt.Printf("猫的信息为=[id=%d name=%s]->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 删除一个结点
func DeleteCatNode(head *CatNode,no int)*CatNode{
	// 删除一个环形单向链表的思路
	// 先让temp指向head
	// 再让helper指向链表的最后
	// 让temp和要删除的id进行比较,如果相同,则同helper完成删除,必须考虑如果要删除的就是头结点

	temp := head
	helper := head
	if temp.next == nil{
		fmt.Println("这是一个空环形链表,无法删除")
		return head
	}

	// 如果只有一个结点
	if temp.next == head{
		temp.next = nil
		return head
	}

	// helper 定位到链表最后一个节点
	for{
		if helper.next == head{
			break
		}
		helper = helper.next
	}


	// 如果包含两个或两个以上
	flag := true
	for{
		if temp.next == head{
			// 说明已经比较到最后一个,最后一个还没有比较
			break
		}
		if temp.no == no{

			if temp == head{
				// 说明删除的是头结点
				head = head.next
			}

			// 找到了  也可以在这里直接删除
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n",no)
			flag = false
			break
		}
		temp = temp.next  // 移动用来  比较
		helper = helper.next // 移动 如果找到要删除的结点,需要这个结点干掉删除的结点
	}

	// 还要比较一次 最后一个
	if flag {
		// 代表for循环中 没有删除过,需要进行最后一个的比较
		if temp.no == no{
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n",no)
		}else{
			fmt.Printf("没有查询到指定的元素... no=%d\n",no)
		}
	}

	return head
}

func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	//cat2 := &CatNode{
	//	no:   2,
	//	name: "tom2",
	//}

	//cat3 := &CatNode{
	//	no:   3,
	//	name: "tom3",
	//}

	InsertCatNode(head, cat1)
	//InsertCatNode(head, cat2)

	//InsertCatNode(head, cat3)
	DeleteCatNode(head,1)
	ListCircleLink(head)
}
