//遍历 map[string]interface{} 类型的value值
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	txt := `{"a":1,"b":2,"c":[{"name":"1","group":"2"},{"name":"3","group":"4"}]}`
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(txt), &m); err != nil {
		panic(err)
	}
	for i, v := range m {
		fmt.Printf("%v : %v \n", i, v)
		fmt.Println(reflect.TypeOf(v))
	}
	//interface 类型转换
	//value.(int)转为int，value.(string) 转为 string
	for _, v := range m["c"].([]interface{}) {
		fmt.Println(v)
	}
}
