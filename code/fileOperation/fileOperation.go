package main ;

import (
	"fmt"
	"os"
	"io/ioutil"
)

//读取整个文件
func read(path string) string {
	fi,err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644);
	if err != nil{
		panic(err);
	}
	defer fi.Close();
	fd,err := ioutil.ReadAll(fi);
	fmt.Println(string(fd));

	s := []byte("Hello World!");

	// ioutil 写文件会清空之前的内容
	//ioutil.WriteFile(path,s,os.ModeAppend);

	fi.Write(s); //在文件末尾追加
	return string(fd);
}

//按行读取文件
func readLine(path string) bool {
	fi,err := os.OpenFile(path,os.O_APPEND|os.O_CREATE, 0644);
	if err != nil {
		panic(err);
	}
	defer fi.Close();
	for {
		var v1,v2,v3 string;
		_,err := fmt.Fscanln(fi,&v1,&v2,&v3);
		//_,err := fmt.Fscanf(fi,"%s%s%s",&v1,&v2,&v3);格式化读取
		if err != nil{
			break
		}
		fmt.Println(v1,v2,v3);
	}
	return true;
}

func write(filename string) bool {
	wireteString := "测试n";
	var d1 = []byte(wireteString);
	err2 := ioutil.WriteFile(filename, d1, 0666) //写入文件(字节数组)
	if err2 != nil{
		return false;
	}else{
		return true;
	}
}

func main(){
	args := os.Args;
	if len(args) < 2 {
		fmt.Println("请输入文件名")
	} else {
		file := args[1];
		//read(file);
		readLine(file);
		write("test1.txt");
	}
}
