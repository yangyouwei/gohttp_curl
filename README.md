# gohttp_curl

没有统计时间，成功返回200，失败返回failue
比较简陋，做压测工具使用。

### usage

      -c int
            -c 5 设置并发数，单位时间内的goroutines数量 (default 5)
      -u string
            -u http://baidu.com 访问url
