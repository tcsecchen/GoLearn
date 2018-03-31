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

func webHandle(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数，默认是不会解析的
    //fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
    //fmt.Println("path", r.URL.Path)
    //fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	
	//模版字符串
	tmpl := `<!DOCTYPE html>
		<html>
			<head>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8"> <title>iast server</title>
			</head>
			<body>
				<div>
					<table border="0">
					{{ range $key, $value := .}}
						<tr>
							<th>{{ $key }}</th>
							<td>{{ parseForm $value}}</td>
						</tr>					
					{{ end }}
					</table>
					<br/>						
				</div>
			</body>
		</html>`

	funcMap := template.FuncMap{"parseForm": parseForm};		
	t := template.New("layout.html").Funcs(funcMap) //创建模版对象
	t, _ = t.Parse(tmpl)  //加载模版字符串
	fmt.Println(t.Name()) 
	t.Execute(w, r.Form);	

	formdata := make(map[string]string);
	//遍历表单数据
    for k, v := range r.Form {
		fmt.Println("key:", k)
		querydata :=  strings.Join(v, "") //表单数据提取后值为数组，使用join转为字符串
		fmt.Println("val:", querydata)
		if k == "id" {
			querydata, _ := strconv.ParseInt(querydata, 10, 64);
			str := sql_query(querydata);
			fmt.Fprintf(w, "查询数据：")
			for _,val := range str {
				fmt.Fprintf(w, " %s ",val);
			}
		} else {
			// s := strings.Split(querydata, " ");
			formdata[k] = querydata;
			
			
			//formdata = append(formdata,querydata);
		}	
	}
	
	row, _ := strconv.ParseInt(formdata["row"],10,64);
	if row > 0 {
		result := sql_insert(formdata["fileName"], formdata["funcName"], row, formdata["issueName"], formdata["thirdLibrary"]);
		if(result){
			fmt.Fprintf(w,"<br/><div>数据插入成功</div><br/>");
		} else {
			fmt.Fprintf(w,"<br/><div>数据插入失败</div><br/>");				
		}
	} else {
		fmt.Fprintf(w,"<br/><div>数据插入失败</div><br/>");				
	}
	

	fmt.Fprintf(w, "<div> Hello astaxie! </div>") //这个写入到w的是输出到客户端的
	data := sql_query_all();
	fmt.Fprintf(w, "<div><br/>全部数据：</div>")
	for _,val := range data {
		fmt.Fprintf(w, "<div> %s </div>",val);
	}
}

//sql查询
func sql_query(queryData int64) []string {
	s := []string{"1","2","3","4","5","6"};

	var err error;
	rows, err := db.Query("SELECT * FROM issues where id=?",queryData);
	if err == nil{
		fmt.Println("查询数据：")
		for rows.Next() {
			var id, fileName, funcName, row, issueName, third_library string;
			// var id, row int64;
			err := rows.Scan(&id, &fileName, &funcName, &row, &issueName, &third_library);
			if err == nil{
				fmt.Println(id, fileName, funcName, row, issueName, third_library);
				s[0] = id;
				s[1] = fileName;
				s[2] = funcName;
				s[3] = row;
				s[4] = issueName;
				s[5] = third_library;
			}			
		}
	}
	return s;
}

func sql_query_all() []string {
	var s []string;

	var err error;
	rows, err := db.Query("SELECT * FROM issues");
	if err == nil{
		fmt.Println("全部数据：")
		for rows.Next() {
			var id, fileName, funcName, row, issueName, third_library string;
			// var id, row int64;
			err := rows.Scan(&id, &fileName, &funcName, &row, &issueName, &third_library);
			if err == nil{
				data := id + " " + fileName + " " + funcName + " " + row + " " + issueName + " " + third_library;
				fmt.Println(data);
				s = append(s, data);
			}			
		}
	}
	return s;
}

func sql_insert(fileName string, funcName string, row int64, issueName string, third_library string) bool{
	var id int64;
	err := db.QueryRow("SELECT id FROM issues where file_name=? and func_name=? and row=? and issue_name=?",fileName, funcName, row, issueName).Scan(&id);
	if err == nil && id != 0{
		fmt.Printf("该条数据已存在,id:%d\n",id);
		return false ;
	} else if err != nil && id == 0{
		fmt.Println(err)
		result, err := db.Exec("INSERT INTO issues VALUES (?,?,?,?,?,?)",nil, fileName, funcName, row, issueName, third_library);
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

func main() {
	var err error;
	db, err = sql.Open("mysql", "root:iloveu@/iast");
	err = db.Ping();
	if err == nil {
		fmt.Println("数据库已连接");
	}

	http.HandleFunc("/", webHandle)
	http.ListenAndServe(":8080", nil)
	
}