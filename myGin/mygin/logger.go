package mygin

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// 开始 time
		t := time.Now()
		// 处理 request
		c.Next()
		// 计算解析时间
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
