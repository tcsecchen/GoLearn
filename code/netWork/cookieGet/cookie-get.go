//带有cookie的Get请求
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"os"
)

func main(){
	client := &http.Client{};
	req, err := http.NewRequest("GET", "http://192.168.1.89/dvwa", nil);
	cookie1 := &http.Cookie{Name:"PHPSESSID",Value:"26c2tkqumv2a2l4o34qtdcbs80",HttpOnly:true};
	cookie2 := &http.Cookie{Name:"security",Value:"impossible",HttpOnly:true};
	req.AddCookie(cookie1);
	req.AddCookie(cookie2);
	cookie,err := req.Cookie("security");
	fmt.Println(cookie);
	resp,err := client.Do(req);
	if err != nil {
		panic(err.Error());
	}
	b, err := ioutil.ReadAll(resp.Body);
	resp.Body.Close();
	fmt.Printf("%s",b);
	
}