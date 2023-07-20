package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.google.com",
		"http://www.taobao.com",
		"http://www.qq.com",
		"http://www.163.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) // 循环接收所有网站的状态信息并打印
	// }
	// for {
	// 	go checkLink(<-c, c) // 无限循环检测所有网站的状态
	// }
	// 和上面相等
	for l := range c {
		go func(link string) {
			// time.Sleep(1 * time.Second)
			checkLink(link, c)
		}(l)
	}

}
func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
