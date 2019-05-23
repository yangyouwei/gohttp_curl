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
	waitgroup.Add(1)

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

	for i := 1;i < concurrent ;i ++  {
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

