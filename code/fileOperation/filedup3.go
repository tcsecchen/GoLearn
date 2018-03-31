package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//找出重复行
func main(){
	counts := make(map[string]int);
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename);
		if err != nil {
			fmt.Println(err);
			continue;
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line]++ ;
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line);
		}
	}
	
}