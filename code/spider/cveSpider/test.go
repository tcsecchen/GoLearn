package main

import (
	"GoLearn/code/spider/cvespider/parse"
	"fmt"
	"strconv"
	"sync"
)

func test() {
	cvelist := make(chan []string, 100)
	//cveInfo := make(chan parse.Info, 100)
	var wgSend sync.WaitGroup
	var wgReceive sync.WaitGroup

	for i := 1; i < 50; i++ {
		url := "http://cve.scap.org.cn/cve_list.php?p=" + strconv.Itoa(i)
		wgSend.Add(1)
		go func(url string) {
			defer wgSend.Done()
			list, _ := parse.GetLinks(url)
			cvelist <- list
		}(url)
	}

	wgReceive.Add(1)
	go func() {
		defer wgReceive.Done()
		for x := range cvelist {
			fmt.Println(x)
		}
		fmt.Println(1)
	}()

	go func() {
		wgSend.Wait()
		close(cvelist)
	}()

	wgReceive.Wait()

	/*
		go func() {
			for x := range cvelist {
				for _, v := range x {
					go func(url string) {
						info, _ := parse.GetContent(url)
						cveInfo <- info
					}(v)
				}
			}
		}()*/
}
