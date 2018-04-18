package main

import (
	"GoLearn/code/netWork/getLinks"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := link.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// breadthFirst 函数对每个worklist元素调用f
// 并将返回的内容添加到worklist中，对每一个元素，最多调用一次f
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
