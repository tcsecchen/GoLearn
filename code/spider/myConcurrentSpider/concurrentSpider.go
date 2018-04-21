//concurrentSpider 是一个简单的并发爬虫的demo
package main

import (
	"GoLearn/code/netWork/getLinks"
	"fmt"
	"log"
	"os"
)

//令牌是一个计数信号量，确保并发请求在20个以内
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := link.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	//并发爬取web
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
