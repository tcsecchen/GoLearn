package main

import (
	"strconv"
	"fmt"
	"net/http"
	"strings"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB ;

func parseForm(val []string) string{
	return strings.Join(val, "")
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数，默认是不会解析的
    //fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
    //fmt.Println("path", r.URL.Path)
    //fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	
	//模版字符串
	tmpl := `<!DOCTYPE html>
		<html>
			<head>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8"> <title>Go Web Programming</title>
			</head>
			<body>
				<div>
					{{ range $key, $value := .}}
					<h1>key:{{ $key }}</h1>
					<h2>value:{{ parseForm $value}}</h2>					
					{{ end }}	
				</div>
			</body>
		</html>`

	funcMap := template.FuncMap{"parseForm": parseForm};		
	t := template.New("layout.html").Funcs(funcMap) //创建模版对象
	t, _ = t.Parse(tmpl)  //加载模版字符串
	fmt.Println(t.Name()) 
	t.Execute(w, r.Form);	

    for k, v := range r.Form {
		fmt.Println("key:", k)
		querydata :=  strings.Join(v, "") 
		fmt.Println("val:", querydata)
		if k == "url_long" {
			str := sql_query(querydata);
			fmt.Fprintf(w, "查询数据：")
			for _,val := range str {
				fmt.Fprintf(w, " %s ",val);
			}

			data := sql_query_all();
			fmt.Fprintf(w, "<div><br/>全部数据：</div>")
			for _,val := range data {
				fmt.Fprintf(w, "<div> %s </div>",val);
			}
		} else {
			s := strings.Split(querydata, " ");
			id, _ := strconv.ParseInt(s[0],10,32);
			result := sql_insert(id, s[1],s[2]);
			if(result){
				fmt.Fprintf(w,"<br/><div>数据插入成功</div><br/>");
			} else {
				fmt.Fprintf(w,"<br/><div>数据插入失败</div><br/>");				
			}
			
		}	
	}
    fmt.Fprintf(w, "<div> Hello astaxie! </div>") //这个写入到w的是输出到客户端的
}

//sql查询
func sql_query(queryData string) []string {
	s := []string{"hello","hi","nihao"};

	var err error;
	rows, err := db.Query("SELECT user_id,first_name,last_name FROM users where first_name=?",queryData);
	if err == nil{
		fmt.Println("查询数据：")
		for rows.Next() {
			var user_id, first_name, last_name string;
			err := rows.Scan(&user_id, &first_name, &last_name);
			if err == nil{
				fmt.Println(user_id,first_name,last_name);
				s[0] = user_id;
				s[1] = first_name;
				s[2] = last_name;
			}			
		}
	}
	return s;
}

func sql_query_all() []string {
	var s []string;

	var err error;
	rows, err := db.Query("SELECT user_id,first_name,last_name FROM users");
	if err == nil{
		fmt.Println("全部数据：")
		for rows.Next() {
			var user_id, first_name, last_name string;
			err := rows.Scan(&user_id, &first_name, &last_name);
			if err == nil{
				data := user_id + " " + first_name + " " + last_name;
				fmt.Println(data);
				s = append(s, data);
			}			
		}
	}
	return s;
}

func sql_insert(id int64, firstName string, lastName string) bool{
	result, err := db.Exec("INSERT INTO users(user_id,first_name,last_name) VALUES (?,?,?)",id, firstName, lastName); 
	if err == nil{
		fmt.Println(result);
		return true ;
	} else {
		fmt.Println(err)
		return false;
	}
}

func main() {
	var err error;
	db, err = sql.Open("mysql", "root:iloveu@/dvwa");
	err = db.Ping();
	if err == nil {
		fmt.Println("数据库已连接");
	}

	http.HandleFunc("/", sayhelloName)
	http.ListenAndServe(":8080", nil)
	
}