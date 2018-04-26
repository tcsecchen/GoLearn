package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type user struct {
	Username string
	Password string
}

//Login 是用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		username := strings.Join(r.Form["username"], "")
		password := strings.Join(r.Form["password"], "")
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		if username == "chen" && password == "123456" {
			fmt.Fprintf(w, "<h1>登录成功！</h1>")
		}
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
		var s user
		err := json.Unmarshal(result, &s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	}

}
