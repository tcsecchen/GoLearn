package main

import (
	"GoLearn/code/spider/cvespider/parse"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//urlTokens 确保请求链接线程在20个以内
var urlTokens = make(chan struct{}, 50)
var infoTokens = make(chan struct{}, 300)
var sqlTokens = make(chan struct{}, 300)

func sqlInsert(db *sql.DB, info parse.Info, logger *log.Logger) bool {
	_, insertErr := db.Exec("INSERT INTO scap VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", nil, info.ID, info.Assortment, info.BugtraqID, info.RemoteOverflow, info.LocalOverflow, info.ReleaseDate, info.UpdateDate, info.Author, info.Version, info.Discuss, info.Exploit, info.Solution, info.Reference)
	if insertErr != nil {
		pattern, _ := regexp.MatchString("1406", insertErr.Error())
		if pattern {
			logger.Println(info.ID, insertErr)
		} else {
			fmt.Println(info.ID, insertErr)
		}
		return false
	}
	fmt.Println(info.ID, "插入成功")
	return true

}

func main() {
	cvelist := make(chan []string, 100)
	cveInfo := make(chan parse.Info, 600)

	db, err := sql.Open("mysql", "root:iloveu@/cvedb")
	err = db.Ping()
	if err == nil {
		fmt.Println("数据库已连接")
	}

	file, err := os.Create("log/log.log")
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	defer file.Close()
	logger := log.New(file, "[Error]", log.LstdFlags)

	go func() {
		for i := 18151; i < 18153; i++ {
			url := "http://cve.scap.org.cn/cve_list.php?p=" + strconv.Itoa(i)
			go func(url string) {
				urlTokens <- struct{}{}
				list, _ := parse.GetLinks(url)
				cvelist <- list
				<-urlTokens
			}(url)
		}
	}()

	go func() {
		for x := range cvelist {
			for _, v := range x {
				go func(url string) {
					infoTokens <- struct{}{}
					info, _ := parse.GetContent(url)
					cveInfo <- info
					<-infoTokens
				}(v)
			}
		}
	}()

	for x := range cveInfo {
		go func(info parse.Info) {
			sqlTokens <- struct{}{}
			sqlInsert(db, info, logger)
			<-sqlTokens
		}(x)
	}

}
