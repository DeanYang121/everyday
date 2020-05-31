package main

import (
	"fmt"
	"os"
)

/*
 实际需求: google一个上机题目: 有一个公司,当有新员工来报道时,要求该员工的信息加入 id,性别,年龄,住址... 当输入该员工的id时,要求查找到该员工的所有信息
 要求: 不使用数据库,尽量节省内存,速度越快越好,哈希表(散列)
 */

type Emp struct {
	Id int
	Name string
	Next *Emp
}
// 方法待定
func (this *Emp) ShowMe(){
	fmt.Printf("链表%d 找到该雇员 %d,%s\n",this.Id %7,this.Id,this.Name)
}

//定义EmpLink
// 这里的EmpLink不带表头,第一个节点就存放雇员信息
type EmpLink struct {
	Head *Emp
}
// 方法待定

func (this *EmpLink)Insert(emp *Emp){ //保证添加时 编号时从小到大
	cur := this.Head // 辅助指针
	var pre *Emp = nil // 辅助指针 协助插入
	// 如果当前的EmpLink 就是空链表
	if cur == nil{
		this.Head = emp
		return
	}
	// 如果不是空链表 找到emp的位置并插入
	// 思路让cur和emp 比较,pre始终保持在cur前面
	for{
			if cur != nil{
				if cur.Id >= emp.Id {
					// 找到位置
					break
				}
				pre = cur
				cur = cur.Next
			}else{
				break
			}
	}

	// 退出时 看下是否将emp添加到链表最后
	pre.Next = emp
	pre.Next.Next = cur
}

// 显示链表信息
func(this *EmpLink)ShowLink(no int){
	if(this.Head == nil){
		fmt.Printf("链表%d 为空\n",no)
		return
	}
	cur := this.Head
	for{
		if cur != nil{
			fmt.Printf("链表%d 雇员id=%d,名字=%s ->",no,cur.Id,cur.Name)
			cur = cur.Next
		}else{
			break
		}
	}
	fmt.Println()
}

// 根据id查找雇员

func(this *EmpLink)FindById(id int)*Emp{
	cur := this.Head

	for{
		if cur != nil && cur.Id == id{
			return cur
		}else if cur == nil{
			break
		}
		cur = cur.Next
	}
	return nil
}

// 定义hashtable 含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// hashTable 编写insert雇员的方法
func(this *HashTable)Insert(emp *Emp){
	// 使用散列函数,确定将该雇员添加到那个链表
	linkNo := this.HashFunc(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp) //
}

// 散列方法
func (this *HashTable)HashFunc(id int)int{
	return id%7 //得到一个值,就是对于链表的下标
}

// 显示hashtabl
func(this *HashTable) ShowAll(){
	for i := 0;i<len(this.LinkArr);i++{
		this.LinkArr[i].ShowLink(i)
	}
}

// 查找
func(this *HashTable)FindById(id int)*Emp{
	linkNo := this.HashFunc(id)
	return this.LinkArr[linkNo].FindById(id)
}

func main() {
	// 实现思路
	// 使用链表实现哈希表,该链表不带表头,第一个节点就存储信息
	var  (
		key string
		id int
		name string
		hashTable HashTable
	)
	for{
		fmt.Println("==========雇员系统菜单===========")
		fmt.Println("input 添加雇员")
		fmt.Println("show  显示雇员")
		fmt.Println("find  查找雇员")
		fmt.Println("exit  退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id: id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("请输入ID号")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil{
				fmt.Printf("id=%d的雇员不存在\n",id)
			}else{
				//编写一个方法显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			fmt.Println("退出成功!")
			os.Exit(0)
		}
	}
}
