package main

import (
	"fmt"
	"regexp"
)

var (
	rePhone string
	reEmail string
	reLink  string
)

func main() {
	//rePhone = `1[3456789]\d{9}` // 手机号正则 简单正则
	rePhone = `(1[3456789]\d)(\d{4})(\d{4})` // 手机号正则 分组  分了四组 [18620010839 186 2001 0839]
	reEmail = `[\w+\.]+@\w+\.[a-zA-Z0-9_]+(\.[a-zA-Z0-9_]+)?`
	reLink = `<a[\s\S]+?href="(http[\s\S]+?)"`

	//demoString1 := "hahahahah"
	demoString := "这是一个测试字符串,this is a test string,18620010839,&&deanyang1996@126.com,deanyang1996@126.com.cn,18620010835"
	re := regexp.MustCompile(reEmail)
	allString := re.FindAllStringSubmatch(demoString, -1) // -1 表示匹配全部
	for key, value := range allString {
		fmt.Println(key, value)
	}
}
