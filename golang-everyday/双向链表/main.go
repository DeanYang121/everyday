package main

import "fmt"

type HeroNode struct{
	no int
	name string
	nickName string
	next *HeroNode // 指向下一个节点
	prev *HeroNode // 指向前一个节点
}

// 有序插入 按照编号从小到大插入
func InsertHeroNode2(head *HeroNode,newHeroNode *HeroNode){

	temp := head
	flag := true
	for{
		if temp.next == nil{
			break // 到了链表的最后
		}else if temp.next.no > newHeroNode.no{
			// 说明newNode 就应用插入temp后面
			break
		}else if temp.next.no == newHeroNode.no{
			// 链表中已经有了no,不插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag{
		fmt.Println("链表中已经存在no=",newHeroNode.no)
		return
	}else{
		newHeroNode.next = temp.next
		newHeroNode.prev = temp

		if temp.next != nil{
			temp.next.prev = newHeroNode
		}
		temp.next = newHeroNode
	}

}


// 按照插入的前后顺序进行插入
func InsertHeroNode(head *HeroNode,newHeroNode *HeroNode){
	// 思路
	// 1.先找到该链表的最后这个节点
	// 2. 创建一个辅助节点 帮忙查找
	temp := head
	for{
		if temp.next == nil{
			break //找到了最后一个
		}
		temp = temp.next //一直找到链表的最后一个节点
	}

	// 3. 插入
	temp.next = newHeroNode
	// 变成双向链表
	newHeroNode.prev = temp

}

// 删除节点 双向链表
func DelHeroNode(head *HeroNode,no int){
	temp := head
	flag := false
	// 寻找需要删除的节点
	for{
		if temp.next == nil{
			break
		}else if temp.next.no == no{
			// 找到目标节点
			flag = true
			break
		}
		temp = temp.next
	}

	if flag{
		// 找到删除
		temp.next = temp.next.next
		if temp.next != nil{
			temp.next.prev = temp
		}
	}else{
		fmt.Println("要删除的id不存在")
	}
}


// 顺序遍历
func ListHeroNode(head *HeroNode){

	//1. 创建一个辅助节点
	temp := head

	if temp.next == nil{
		fmt.Println("空链表....")
		return
	}

	//2. 遍历这个链表
	for{
		fmt.Printf("[%d,%s,%s]-->",temp.next.no,temp.next.name,temp.next.nickName)
		// 判断是否为链表的最后一个
		temp = temp.next
		if temp.next == nil{
			break
		}
	}
}

// 双向链表 反向遍历
func ListHeroNodeReverse(head *HeroNode){

	//1. 创建一个辅助节点
	temp := head

	if temp.next == nil{
		fmt.Println("空链表....")
		return
	}

	//2. 让temp定位到双向链表的最后结点
	for {
		if temp.next == nil{
			break
		}
		temp = temp.next
	}

	//3. 遍历这个链表
	for{
		fmt.Printf("[%d,%s,%s]-->",temp.no,temp.name,temp.nickName)
		// 判断是否为链表的头部
		temp = temp.prev
		if temp.prev == nil{
			break
		}
	}

}


func main() {
	// 1.创建一个头结点
	head := &HeroNode{}

	// 2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:1,
		name:"宋江",
		nickName:"及时雨",
	}

	hero2 := &HeroNode{
		no:2,
		name:"卢俊义",
		nickName:"玉麒麟",
	}

	hero3 := &HeroNode{
		no:3,
		name:"林冲",
		nickName:"豹子头",
	}

	InsertHeroNode(head,hero1)
	InsertHeroNode(head,hero2)
	InsertHeroNode(head,hero3)
	//DelHeroNode(head,2)
	ListHeroNode(head)
	fmt.Println()
	ListHeroNodeReverse(head)
}
