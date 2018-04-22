//buffer1Channel 是比较缓冲为1的通道和无缓冲通道的demo
package main

import (
	"fmt"
	"time"
)

//等价于var nobuffChan = make(chan int，0)
var nobuffChan = make(chan string)
var buff1Chan = make(chan string, 1)
var exit = make(chan int)

func main() {

	//无缓冲通道必须有其他goroutine同步接收，否则阻塞
	//本例中一个goroutine负责发送，主goroutine 负责接收
	/*以下代码会被阻塞
	nobuffChan <- a
	fmt.Println(<-nobuffChan)
	*/

	//缓冲为1的通道有一个大小为1的缓冲区，不须有其他goroutine同步接收，但是只能发送一个
	/*以下代码不会阻塞
	buff1Chan <- a
	fmt.Println(<-buff1Chan)
	*/

	go func() {
		fmt.Println("buff1Chan准备发送了")
		buff1Chan <- "buff1Send"
		fmt.Println("buff1Chan发送成功了")
	}()
	go func() {
		fmt.Println("nobuffChan准备发送了")
		nobuffChan <- "nobuffSend"
		fmt.Println("nobuffChan阻塞掉了")
		<-exit
	}()

	time.Sleep(3 * time.Second)
	fmt.Println(<-nobuffChan)
	fmt.Println(<-buff1Chan)
	exit <- 0
}
