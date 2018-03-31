package file 

import (
	"os"
	"io/ioutil"
	"strings"
)

func Read(path string) []string {
	fi,err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644);
	if err != nil{
		panic(err);
	}
	defer fi.Close();
	fd,err := ioutil.ReadAll(fi);
	// fmt.Println(string(fd));
	fdArr := strings.Split(string(fd), "\n")
	return fdArr;
}