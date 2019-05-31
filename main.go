package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var waitgroup sync.WaitGroup

func main()  {
	waitgroup.Add(1)  //因为是1下面没有done，所以main会一直等待go程运行完毕。

	Address := flag.String("u","","-u http://baidu.com 访问url")
	Concurrent := flag.Int("c",5,"-c 5 设置并发数，单位时间内的goroutines数量")
	flag.Parse()

	if *Address == "" {
		flag.Usage()
		return
	}

	concurrent := *Concurrent

	c := make(chan *string, concurrent)

	client := &http.Client{}

	go func ()  {
		for {
			c <- Address
		}
	}()

	for i := 1;i < concurrent ;i ++  {//没有waitgroup的话，运行完for循环，主程序就退出了。goroutines也就结束了，看不到结果。
		go func (){
			for {
				address := <- c
				code := httpget(address,client)
				if code == 200 {
					fmt.Println("200")
				}else {
					fmt.Println("failue")
				}
			}
		}()
	}

	waitgroup.Wait()
}

func httpget(url *string ,c *http.Client) int {
	resp, err:= c.Get(*url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return 900
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 901
	}
	if len(body) > 0 {
		//fmt.Println(string(body))
	}
	return resp.StatusCode
}

// func do_somethine1(wg *sync.WaitGroup)  {
// 	wg.Done()
// 	fmt.Println("do_somethine1")
// }
// func do_somethine2(wg *sync.WaitGroup)  {
// 	wg.Done()
// 	fmt.Println("do_somethine1")
// }
// func do()  {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go do_somethine1(&wg)
// 	go do_somethine2(&wg)
// 	wg.Wait()
// }
//几个goroutines 就add多少
//wait会等待goroutines完成后退出
//wg=0时就不等待了。
//每个goroutines运行后都会done -1
