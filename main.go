package main

import (
	"ett"
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
	r.Use(ett.Logger()) // global midlleware
	r.GET("/", func(c *ett.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ett</h1>")
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
