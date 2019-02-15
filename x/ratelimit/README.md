# 限流

## 采用令牌桶算法
令牌桶算法的原理是系统会以一个恒定的速度往桶里放入令牌，而如果请求需要被处理，则需要先从桶里获取一个令牌，当桶里没有令牌可取时，则拒绝服务。

## 示例

```
package main

import (
	"log"
	"time"

	"github.com/yangwenmai/ratelimit/simpleratelimit"
)

func main() {
    // rate limit: 每秒十个
	rl := simpleratelimit.New(10, time.Second)

	for i := 0; i < 100; i++ {
		log.Printf("limit result: %v\n", rl.Limit())
	}
	log.Printf("limit result: %v\n", rl.Limit())

}
```


## 性能测试

```
go test -bench=. -benchtime="5s"
goos: darwin
goarch: amd64
BenchmarkLimit-4   	200000000	        35.4 ns/op	       0 B/op	       0 allocs/op
```

> 线程安全