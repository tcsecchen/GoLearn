这是我个人学习go时踩的一些坑，后面学习途中会逐步的更新

1. 全局变量可以声明不使用，其他的变量声明后必须使用！！
2. 交换两个变量的值，则可以简单地使用 a, b = b, a。
3. for 循环条件不能加圆括号！
4. go里的自增运算符只有——“后++”，i++为语句，不是表达式，j = i++ 是错的
5. 数组传递的是拷贝，不是引用
6. go 结构体转json时候：结构体必须是大写字母开头的成员才会被处理到！！！
7. json 转结构体时，同样结构体成员也必须为大写字母开头
8. golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
9. out chan<- int 是单向通道中的发送通道, in <-chan int 是单向通道中的接收通道 
10. 无缓冲通道是缓冲为0的通道，发送时必须有其他goroutine同步接收，否则会被阻塞;而缓冲为1的通道有一个大小为1的缓冲区，无须同步接收。形象的例子：无缓冲通道好比必须要你本人签收的快递，你不来他不走，进入阻塞状态；缓冲为1的通道好比你有一个容量为1的丰巢快递柜，快递员把快递放进去就可以走，但是必须取出后才能放第2个(demo可参考[buffer1channel.go](https://github.com/tcsecchen/GoLearn/blob/master/code/goroutine/buffer1Channel/buffer1channel.go))
11. 修改web服务器response响应头可以用 w.Header().Set(k,v) 和 w.writeHeader(code), 设置http status code(http状态码) 只能通过 w.writeHeader(404) 来实现
12. w.Header().Set(k,v) 只能在 w.writeHeader(code) 之前使用，否则会无效
