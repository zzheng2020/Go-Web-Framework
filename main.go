package main

import (
	"ett"
	"net/http"
)

func main() {
	r := ett.New()
	r.GET("/index", func(c *ett.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *ett.Context) {
			c.HTML(http.StatusOK, "<h1>Hello ett</h1>")
		})

		v1.GET("/hello", func(c *ett.Context) {
			// expect /hello?name=ettktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ett.Context) {
			// expect /hello/ettktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *ett.Context) {
			c.JSON(http.StatusOK, ett.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
