package main

import (
	"fmt"
)

func test() {
	for i := 0; i < 9; i++ {
		m := 9 / 9
		j := m*i + 1
		k := j + m
		fmt.Println(j, k)
	}

	// cvelist := make(chan []string, 100)
	// //cveInfo := make(chan parse.Info, 100)
	// var wgSend sync.WaitGroup
	// var wgReceive sync.WaitGroup

	// for i := 1; i < 50; i++ {
	// 	url := "http://cve.scap.org.cn/cve_list.php?p=" + strconv.Itoa(i)
	// 	wgSend.Add(1)
	// 	go func(url string) {
	// 		defer wgSend.Done()
	// 		list, _ := parse.GetLinks(url)
	// 		cvelist <- list
	// 	}(url)
	// }

	// wgReceive.Add(1)
	// go func() {
	// 	defer wgReceive.Done()
	// 	for x := range cvelist {
	// 		fmt.Println(x)
	// 	}
	// 	fmt.Println(1)
	// }()

	// go func() {
	// 	wgSend.Wait()
	// 	close(cvelist)
	// }()

	// wgReceive.Wait()

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
	// info, _ := parse.GetContent("http://cve.scap.org.cn/CVE-2018-2819.html")
	// fmt.Println(info.Version)

}
