package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type valNode struct {
	row int
	col int
	val int
}

func main() {
	// 1. 创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1  // 黑棋
	chessMap[2][3] = 2  // 蓝棋
	var rows int
	var cols int
	// 2. 查看原始数组
	for _, v := range chessMap {
		rows++
		cols = 0
		for _,v2 := range v{
			cols++
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组
	// 思路
	// (1) 遍历chessMap 如果发现有一个元素的值不为0,创建一个node结构体
	// (2) 将其放入到对应的切片即可

	var sparseArr []valNode

	// 标准的稀疏数组 还需要记录规模 几行几列 默认值是什么
	valnode := valNode{
		row: rows,
		col: cols,
		val: 0,
	}
	sparseArr = append(sparseArr,valnode)

	for i, v := range chessMap {
		for j,v2 := range v{
			if v2 != 0{
				_valNode := valNode{
					row:i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr, _valNode)
			}
		}
	}
	fmt.Println("**********: 转换成稀疏数组")
	err := os.Remove("/home/docker/everyday/golang-everyday/稀疏数组/chess.data")
	if err != nil{
		fmt.Println("删除文件失败:",err.Error())
		return
	}

	file,err := os.OpenFile("/home/docker/everyday/golang-everyday/稀疏数组/chess.data",os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)

	for key, value := range sparseArr {
		str := fmt.Sprintf("%d %d %d\n",value.row,value.col,value.val)
		fmt.Printf("%d: %d %d %d\n",key,value.row,value.col,value.val)
		writer.WriteString(str)
	}

	// 将稀疏数组存盘
	writer.Flush()


	//从文件读取稀疏数组
	file,err = os.Open("/home/docker/everyday/golang-everyday/稀疏数组/chess.data")
	if err != nil{
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)

	var chessMap2 [11][11]int
	flag := 0
	for{
		str,err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF{
			break
		}
		flag++
		if flag != 1{
			fmt.Println("123")
			str = strings.ReplaceAll(str,"\n","")
			_rows := strings.Split(str," ")
			_row_index,_ := strconv.Atoi(_rows[0])
			_col_index,_ := strconv.Atoi(_rows[1])
			_value,err := strconv.Atoi(_rows[2])
			if err != nil{
				fmt.Println(err)
			}
			fmt.Println(_row_index,_col_index,_value)
			chessMap2[_row_index][_col_index] = _value
		}
	}

	fmt.Println("*******:::: 还原回原始数组")
	for _, v := range chessMap2 {
		rows++
		cols = 0
		for _,v2 := range v{
			cols++
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

}



// 加餐彩蛋
// 1. 一次性将文件读取到位
func readFileAll(){

	file := "./1.txt"
	var content []byte
	var err error
	if content,err = ioutil.ReadFile(file);err != nil{   // ReadFile 已经封装了open 和close 所以不需要close文件
		fmt.Printf("read file err=%v",err.Error())
		return
	}

	fmt.Printf("%v",string(content))
}


// 2. 读取文件
func openFile(){
	file,err := os.Open("./1.txt")
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("file=%v",file)
	err = file.Close()
	if err != nil{
		fmt.Println(err)
	}

}


//3. 读取文件 带缓冲
func readFIle2(){
	file,err := os.Open("./1.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一个 *Reader 带缓冲的
	const (
		defaultBufferSize = 4096 // 默认缓冲区大小4096
	)
	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF{
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束")
}

//4. 写入文件
func wirteFile(){
	filePath := "./1.txt"
	file,err := os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	str := "hello \n"
	writer := bufio.NewWriter(file)
	for i := 0;i<5;i++{
		writer.WriteString(str)
	}
	// 将写入缓存的数据 flush写入文件中
	writer.Flush()
}