//antdWebServer 是一个前后端分离demo， golang使用Restful API 风格
//全部数据传输采用 json 格式
package main

import (
	"GoLearn/code/web/antdWebServer/web/user"
	"net/http"
)

func main() {
	http.HandleFunc("/login", user.Login)
	http.ListenAndServe(":80", nil)
}
