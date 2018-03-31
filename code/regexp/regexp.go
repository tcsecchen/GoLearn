package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "中国hello你好hi";
	reg := regexp.MustCompile("[\\p{Han}]+");
	//查找
	//fmt.Println(reg.FindAllString(str,-1));
	//替换
	fmt.Println(reg.ReplaceAllString(str,"\n$0 "));
}