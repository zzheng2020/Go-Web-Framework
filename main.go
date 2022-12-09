package main

import (
	"ett"
	"ettcache/lru"
	"fmt"
	"log"
	"net/http"
	"time"
)

func onlyForV2() ett.HandlerFunc {
	return func(c *ett.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := ett.New()
	r.Use(ett.Logger(), ett.Recovery()) // global midlleware

	r.GET("/", func(c *ett.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ett</h1>")
	})

	test := lru.New(int64(0), nil)
	fmt.Println(test)

	r.GET("/panic", func(ctx *ett.Context) {
		names := []string{"test panic"}
		ctx.String(http.StatusOK, names[10000]) // raise panic
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *ett.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
