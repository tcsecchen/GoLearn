package main;

import(
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
)

func read(path string) string {
	fi,err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644);
	if err != nil{
		panic(err);
	}
	defer fi.Close();
	fd,err := ioutil.ReadAll(fi);
	//fmt.Println(string(fd));

	return string(fd);
}

func readLine(path string) bool {
	fi,err := os.OpenFile(path,os.O_APPEND|os.O_CREATE, 0644);
	if err != nil {
		panic(err);
	}
	defer fi.Close();
	for {
		var v1,v2 string;
		//_,err := fmt.Fscanln(fi,&v1,&v2);
		_,err := fmt.Fscanf(fi,"%s %s\n",&v1,&v2);//格式化读取
		if err != nil{
			break
		}
		fmt.Println(v1,v2);
	}
	return true;
}

func write(filename string, context string) bool {
	var d1 = []byte(context);
	err2 := ioutil.WriteFile(filename, d1, 0666) //写入文件(字节数组)
	if err2 != nil{
		return false;
	}else{
		return true;
	}
}
          
func main(){
	args := os.Args;
	reg := regexp.MustCompile("[\\p{Han}]+");
	if len(args) < 2 {
		fmt.Println("请输入文件名")
	} else {
		file := args[1];
		context := read(file);
		//fmt.Println(context);
		//readLine(file);
		context2 := reg.ReplaceAllString(context,"\n$0 ");
		write("list.txt", context2);
	}
}