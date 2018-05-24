/*此文件模拟一个场景：n个页面，每页有6个子链接，获得子链接后
再进入子链接爬取里面的内容，然后将内容存入数据库
可以作为 m 个 发送者，n个接收者 处理channel及goroutine的模型
*/
package main

import (
	"fmt"
	"sync"
)

func getLinks(page int) (link int) {
	link = page
	return link
}

func getContent(link int) (content int) {
	content = link
	return content
}

func main() {

	go func() { //让程序不会崩溃退出
		for {
		}
	}()

	pageChannel := make(chan int, 10)
	linksChannel := make(chan int, 10)
	contentChannel := make(chan int, 10)

	//将页面URL发送至pageChannel
	go func() {
		defer close(pageChannel)
		for page := 1; page < 10; page++ {
			pageChannel <- page
		}
	}()

	wgGetLink := sync.WaitGroup{}
	//爬取每页子链接后将子链接发送至linkChannel
	go func() {
		defer close(linksChannel)
		for page := range pageChannel {
			wgGetLink.Add(1)
			go func(page int) {
				defer wgGetLink.Done()
				link := getLinks(page)
				linksChannel <- link
			}(page)
		}
		wgGetLink.Wait()
	}()

	wgGetContent := sync.WaitGroup{}
	//爬取子链接页面内容后将内容发送至contentChannel
	go func() {
		defer close(contentChannel)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			fmt.Println("receive1")
			defer wg.Done()
			for link := range linksChannel {
				wgGetContent.Add(1)
				go func(link int) {
					defer wgGetContent.Done()
					content := getContent(link)
					contentChannel <- content
				}(link)
			}
			wgGetContent.Wait()
		}()
		wg.Add(1)
		go func() {
			fmt.Println("receive2")
			defer wg.Done()
			for link := range linksChannel {
				wgGetContent.Add(1)
				go func(link int) {
					defer wgGetContent.Done()
					content := getContent(link)
					contentChannel <- content
				}(link)
			}
			wgGetContent.Wait()
		}()
		wg.Wait()
	}()

	for content := range contentChannel {
		fmt.Println(content)
	}

}
