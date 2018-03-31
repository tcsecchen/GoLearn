package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB ;

type Context struct{
	Status string
	Data DataObeject
}
type DataObeject struct{
	Type int32
}

func requestLine(id string) string {
	var context Context;
	client := &http.Client{};
	url := "http://www.vulbox.com/user/get_business/?name=" + id;
	req, err := http.NewRequest("GET", url, nil);
	cookie1 := &http.Cookie{Name:"Hm_lvt_ca048dd0dbb114747cc58907f7e20022",Value:"1518081837,1518138228,1519606977,1519607008",HttpOnly:true};
	cookie2 := &http.Cookie{Name:"Key_auth",Value:"mUx8Dkn2xTElkijQRduOAtV38gOpw76LyOc9z%2BvSBKU%3D",HttpOnly:true};
	cookie3 := &http.Cookie{Name:"VulBox_uid",Value:"TUGxAsO47liQ1batmBIm52zyA1uccw0hquzZICHlF7L9%2FDRfUC8LCAjflxC%2Fvino",HttpOnly:true};
	cookie4 := &http.Cookie{Name:"PHPSESSID",Value:"342p3l2ik99sad3ljtb0ssf8s4",HttpOnly:true};
	cookie5 := &http.Cookie{Name:"Hm_lpvt_ca048dd0dbb114747cc58907f7e20022",Value:"1519613514",HttpOnly:true};
	req.AddCookie(cookie1);
	req.AddCookie(cookie2);
	req.AddCookie(cookie3);
	req.AddCookie(cookie4);
	req.AddCookie(cookie5);
	resp,err := client.Do(req);
	if err != nil {
		panic(err.Error());
	}
	b, err := ioutil.ReadAll(resp.Body);
	resp.Body.Close();
	text := json.Unmarshal(b, &context);
	if text == nil{
		//fmt.Printf("%d",context.Data.Type);
		switch(context.Data.Type){
			case 0:
			  return "A";
			case 1:
			  return "B";
			case 2:
			  return "C";
			default:
			  return "0";
		}
	}else{
		return "0";
	}
}

func sql_insert(firmName string, firmUrl string, frimLevel string) bool{
	var id int64;
	err := db.QueryRow("SELECT firm_id FROM firmlist where firm_name=? and firm_url=? and firm_level=?",firmName, firmUrl, frimLevel).Scan(&id);
	if err == nil && id != 0{
		fmt.Printf("该条数据已存在,id:%d\n",id);
		return false ;
	} else if err != nil && id == 0{
		fmt.Println(err)
		result, err := db.Exec("INSERT INTO firmlist VALUES (?,?,?,?)",nil,firmName, firmUrl, frimLevel);
		if err == nil{
			fmt.Println(result);
			return true ;
		} else {
			fmt.Println(err)
			return false;
		}
	} else{
		return false;
	}
	//
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
		level := requestLine(v1);
		fmt.Println(v1,v2,level);
		result := sql_insert(v1,v2,level);
		if !result {
			fmt.Println("数据插入失败");
		}
	}
	return true;
}

func main(){
	var err error;
	db, err = sql.Open("mysql", "root:iloveu@/vulbox");
	err = db.Ping();
	if err == nil {
		fmt.Println("数据库已连接");
	}

	args := os.Args;
	if len(args) < 2 {
		fmt.Println("请输入文件名")
	} else {
		file := args[1];
		readLine(file);
	}
}