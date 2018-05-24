package main

import (
	"GoLearn/code/spider/cvespider/parse"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//urlTokens 确保请求链接线程在50个以内
var urlTokens = make(chan struct{}, 50)
var infoTokens = make(chan struct{}, 300)
var sqlTokens = make(chan struct{}, 300)

func sqlInsert(db *sql.DB, info parse.Info, logger *log.Logger) bool {
	_, insertErr := db.Exec("INSERT INTO scap VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", nil, info.ID, info.Name, info.Assortment, info.BugtraqID, info.RemoteOverflow, info.LocalOverflow, info.ReleaseDate, info.UpdateDate, info.Author, info.Version, info.Discuss, info.Exploit, info.Solution, info.Reference)
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
	cvepage := make(chan string, 10000)
	cvelist := make(chan []string, 100)
	cveInfo := make(chan parse.Info, 600)

	timeStart := time.Now()

	db, err := sql.Open("mysql", "root:iloveu@/cvedb")
	err = db.Ping()
	if err == nil {
		fmt.Println("数据库已连接")
	}

	//输出日志到文件
	file, err := os.Create("log/log.log")
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	defer file.Close()
	logger := log.New(file, "[Error]", log.LstdFlags)

	go func() {
		defer close(cvepage)
		for i := 1; i < 50; i++ {
			pageURL := "http://cve.scap.org.cn/cve_list.php?p=" + strconv.Itoa(i)
			cvepage <- pageURL
		}
	}()

	wgGetLink := sync.WaitGroup{}
	//启动一个goroutine循环爬取每页cve链接
	go func() {
		defer close(cvelist)
		for x := range cvepage {
			wgGetLink.Add(1)
			go func(url string) {
				defer wgGetLink.Done()
				urlTokens <- struct{}{}
				list, _ := parse.GetLinks(url)
				cvelist <- list
				<-urlTokens
			}(x)
		}
		wgGetLink.Wait()
	}()

	wgGetCveInfo := sync.WaitGroup{}
	//启动一个goroutine 循环接收cvelist，爬取每个cve的内容
	go func() {
		defer close(cveInfo)
		for x := range cvelist {
			for _, v := range x {
				wgGetCveInfo.Add(1)
				go func(url string) {
					defer wgGetCveInfo.Done()
					infoTokens <- struct{}{}
					info, _ := parse.GetContent(url)
					cveInfo <- info
					<-infoTokens
				}(v)
			}
		}
		wgGetCveInfo.Wait()
	}()

	wgSQLInsert := sync.WaitGroup{}
	//循环接收cve信息， 启动goroutine 插入数据库
	for x := range cveInfo {
		wgSQLInsert.Add(1)
		go func(info parse.Info) {
			defer wgSQLInsert.Done()
			sqlTokens <- struct{}{}
			sqlInsert(db, info, logger)
			<-sqlTokens
		}(x)
	}

	wgSQLInsert.Wait()
	elapsed := time.Since(timeStart)
	fmt.Println("cve爬取完成，耗时：", elapsed)
}
