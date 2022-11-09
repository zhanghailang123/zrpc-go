package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//声明一个等待数组 有点类似于java中的countDownLatch
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.github.com/",
		"http://www.baidu.com/",
		"http://c.biancheng.net/view/108.html",
	}

	for _, url := range urls {
		//遍历切片并且给计数器+1
		wg.Add(1)

		go func(url string) {

			//使用defer，表示函数完成时将等待组值-1
			defer wg.Done()
			resp, err := http.Get(url)
			fmt.Println(resp)
			fmt.Println(url, err)
		}(url)
	}

	wg.Wait()
	fmt.Println("==over==")
}
