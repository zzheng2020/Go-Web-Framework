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

	r.GET("/hello", func(ctx *ett.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(ctx *ett.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})

	r.GET("/assets/*filepath", func(c *ett.Context) {
		c.JSON(http.StatusOK, ett.H{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":9999")

}
