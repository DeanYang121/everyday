package stack

import "fmt"

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{}) interface{}
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	dataSource []interface{}
	capSize int //最大范围
	currentSize int //最大大小
}

func (stack *Stack)Clear(){
	stack.dataSource = make([]interface{},0,10)
	stack.capSize = 1000
	stack.currentSize = 0
}
func (stack *Stack) Size() int{
	return stack.currentSize
}

func (stack *Stack) Pop() interface{}{
	if !stack.IsEmpty(){
		last := stack.dataSource[stack.currentSize-1]
		stack.dataSource = stack.dataSource[:stack.currentSize-1] //删除最后一个
		stack.currentSize--
		return last
	}
	return nil
}

func (stack *Stack) Push(data interface{}) interface{}{
	if !stack.IsFull(){
		stack.dataSource =append(stack.dataSource,data)
		stack.currentSize++
	}
	return  nil
}

func (stack *Stack) IsFull() bool{
	if stack.currentSize >= stack.capSize{
		return true
	}else{
		return false
	}
}

func (stack *Stack) IsEmpty() bool{
	if stack.currentSize == 0{
		return true
	}else{
		return false
	}
}


func NewStack()*Stack{
	stack := new(Stack)
	stack.dataSource = make([]interface{},0,10)
	stack.capSize = 1000
	stack.currentSize = 0
	return stack
}

func main() {
	myStack := NewStack()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
}
