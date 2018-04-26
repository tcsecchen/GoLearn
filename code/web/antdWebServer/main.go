package main

import (
	"GoLearn/code/web/antdWebServer/web/user"
	"net/http"
)

func main() {
	http.HandleFunc("/login", user.Login)
	http.ListenAndServe(":80", nil)
}
