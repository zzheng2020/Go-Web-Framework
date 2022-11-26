package main

import (
	"ett"
	"net/http"
)

func main() {
	r := ett.New()

	r.GET("/", func(ctx *ett.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Ett</h1>")
	})

	r.GET("/hello", func(c *ett.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(ctx *ett.Context) {
		ctx.JSON(http.StatusOK, ett.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.Run(":9999")

}
