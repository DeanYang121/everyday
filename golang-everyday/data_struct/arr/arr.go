package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         //获取第几个元素
	Set(index int, newVal interface{}) error    // 修改数据
	Insert(index int, newVal interface{}) error //插入数据
	Append(newVal interface{})                  //追加数据
	Clear()                                     //清空数组
	Delete(index int) error                     //删除数组
	String() string                             //返回字符串
	Iterator()Iterator
}

// 数据结构 字符串 整数 实数
type ArrayList struct {
	dataStore []interface{} // 数组存储
	TheSize   int           // 数组大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.TheSize //返回数组大小
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.TheSize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //重新叠加
	list.TheSize--
	return nil
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newVal // 设置数据
	return nil
}

func (list *ArrayList) checkIsFull() {
	if list.TheSize == cap(list.dataStore) {
		// 判断内存使用
		newDataStore := make([]interface{}, list.TheSize, 2*list.TheSize) // 开辟双倍内存 同时len也需要设置 不然内存空间没有开辟
		copy(newDataStore, list.dataStore)                    //拷贝  copy 有问题
		//for i:=0;i<len(list.dataStore);i++{
		//	newDataStore[i]=list.dataStore[i]

		//}
		list.dataStore = newDataStore
	}
}

func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkIsFull()                               // 检测内存,如果满了,自动追加
	list.dataStore = list.dataStore[:list.TheSize+1] //插入数据,内存移动一位
	for i := list.TheSize; i > index; i-- { // 从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newVal //插入数据
	list.TheSize++ //索引增加
	return nil
}

//func main1() {
//	var list List = NewArrayList()
//	list.Append(1)
//	fmt.Println(list)  // 因为有string 这个方法 所以才能直接打印
//}

//func main()  {
//	var list List = NewArrayList()
//	list.Append("A")
//	list.Append("B")
//	list.Append("C")
//	for i := 0;i<10;i++{
//		list.Insert(1,"D")
//		fmt.Println(list)
//	}
//}

// 第二章 可迭代的数组

type Iterator interface {
	HasNext() bool
	Next()(interface{},error)
	Remove()
	GetIndex()int
}

type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

// 构造指针访问数组
type ArrayListIterator struct {
	list *ArrayList   //数组指针
	currentIndex int //当前索引
}


func (it *ArrayListIterator) HasNext() bool{
	return it.currentIndex < it.list.TheSize //是否有下一个
}

func(it *ArrayListIterator) Next()(interface{},error){
	if !it.HasNext(){
		return nil,errors.New("没有下一个")
	}
	value,err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value,err
}

func (it *ArrayListIterator) Remove(){
	it.currentIndex--
	_ = it.list.Delete(it.currentIndex) //删除一个元素
}

func(it *ArrayListIterator) GetIndex()int{
	return it.currentIndex
}


func (list *ArrayList)Iterator()Iterator{
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

//func main(){
//	var list List = NewArrayList()
//	list.Append("a")
//	list.Append("b")
//	list.Append("c")
//	list.Append("d")
//	for it := list.Iterator();it.HasNext();{
//		item,_ := it.Next()
//		if item == "c"{
//			it.Remove()
//		}
//		fmt.Println(item)
//	}
//	fmt.Println(list)
//}


// 第三章  栈  利用数组实现
type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{}) interface{}
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	myAarry *ArrayList
	capSize int //最大范围
}

func (stack *Stack)Clear(){
	stack.myAarry.Clear()
}
func (stack *Stack) Size() int{
	return stack.myAarry.TheSize
}

func (stack *Stack) Pop() interface{}{
	if !stack.IsEmpty(){
		last := stack.myAarry.dataStore[stack.myAarry.TheSize-1]
		_ = stack.myAarry.Delete(stack.myAarry.TheSize - 1)
		return last
	}
	return nil
}

func (stack *Stack) Push(data interface{}) interface{}{
	if !stack.IsFull(){
		stack.myAarry.Append(data)
	}
	return  nil
}

func (stack *Stack) IsFull() bool{
	if stack.myAarry.TheSize >= stack.capSize{
		return true
	}else{
		return false
	}
}

func (stack *Stack) IsEmpty() bool{
	if stack.myAarry.TheSize == 0{
		return true
	}else{
		return false
	}
}


func NewArrayListStack()*Stack{
	stack := new(Stack)
	stack.myAarry = NewArrayList()
	stack.capSize = 10
	return stack
}

//func main() {
//	myStack := NewArrayListStack()
//	myStack.Push(1)
//	myStack.Push(2)
//	myStack.Push(3)
//	myStack.Push(4)
//	fmt.Println(myStack.Pop())
//	fmt.Println(myStack.Pop())
//	fmt.Println(myStack.Pop())
//	fmt.Println(myStack.Pop())
//}

// 第四章 可迭代的栈
type StackX struct {
	myAarry *ArrayList
	it Iterator
}

func (stack *StackX)Clear(){
	stack.myAarry.Clear()
}
func (stack *StackX) Size() int{
	return stack.myAarry.TheSize
}

func (stack *StackX) Pop() interface{}{
	if !stack.IsEmpty(){
		last := stack.myAarry.dataStore[stack.myAarry.TheSize-1]
		_ = stack.myAarry.Delete(stack.myAarry.TheSize - 1)
		return last
	}
	return nil
}

func (stack *StackX) Push(data interface{}) interface{}{
	if !stack.IsFull(){
		stack.myAarry.Append(data)
	}
	return  nil
}

func (stack *StackX) IsFull() bool{
	if stack.myAarry.TheSize >= 10{
		return true
	}else{
		return false
	}
}

func (stack *StackX) IsEmpty() bool{
	if stack.myAarry.TheSize == 0{
		return true
	}else{
		return false
	}
}


func NewArrayListStackX()*StackX{
	stack := new(StackX)
	stack.myAarry = NewArrayList()
	stack.it = stack.myAarry.Iterator()
	return stack
}

//func main() {
//	myStack := NewArrayListStackX()
//	myStack.Push(1)
//	myStack.Push(2)
//	myStack.Push(3)
//	myStack.Push(4)
//	//fmt.Println(myStack.Pop())
//	//fmt.Println(myStack.Pop())
//	//fmt.Println(myStack.Pop())
//	//fmt.Println(myStack.Pop())
//	for it := myStack.it;it.HasNext();{  //这个地方迭代访问不能直接使用arrayList的  顺序访问 错误的!
//		item,_ := it.Next()
//		fmt.Println(item)
//	}
//}

// 第四章 文件遍历 使用栈实现递归
type StackA struct {
	dataSource []interface{}
	capSize int //最大范围
	currentSize int //最大大小
}

func (stack *StackA)Clear(){
	stack.dataSource = make([]interface{},0,10)
	stack.capSize = 1000
	stack.currentSize = 0
}
func (stack *StackA) Size() int{
	return stack.currentSize
}

func (stack *StackA) Pop() interface{}{
	if !stack.IsEmpty(){
		last := stack.dataSource[stack.currentSize-1]
		stack.dataSource = stack.dataSource[:stack.currentSize-1] //删除最后一个
		stack.currentSize--
		return last
	}
	return nil
}

func (stack *StackA) Push(data interface{}) interface{}{
	if !stack.IsFull(){
		stack.dataSource =append(stack.dataSource,data)
		stack.currentSize++
	}
	return  nil
}

func (stack *StackA) IsFull() bool{
	if stack.currentSize >= stack.capSize{
		return true
	}else{
		return false
	}
}

func (stack *StackA) IsEmpty() bool{
	if stack.currentSize == 0{
		return true
	}else{
		return false
	}
}


func NewStack()*StackA{
	stack := new(StackA)
	stack.dataSource = make([]interface{},0,10)
	stack.capSize = 1000
	stack.currentSize = 0
	return stack
}

func main() {
	path := "/home/docker/everyday"
	files := make([]string,100)
	stack := NewStack()
	stack.Push(path)

	//defer func() {
	//	recover()
	//}()  // recover 捕获panic

	for !stack.IsEmpty(){
		path := stack.Pop().(string)
		files = append(files,path) // 加入列表
		fmt.Println(path)
		read,err := ioutil.ReadDir(path)
		if err != nil{
			fmt.Println(err)
			panic("读取路径错误")
		}
		for _,fi := range read{
			if fi.IsDir(){
				fulldir := path+"/"+fi.Name()
				stack.Push(fulldir)
			}else{
				fullDir := path+"/"+fi.Name()
				files = append(files,fullDir)
			}
		}
	}


	for i := 0;i<len(files);i++{
		fmt.Println(files[i])
	}

}


