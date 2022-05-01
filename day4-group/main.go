package main

import (
	"github.com/siriusol/gee_study/day4-group/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee at v1</h1>")
		})
		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=Ther
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee at v2</h1>")
		})
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/Ther2
			c.String(http.StatusOK, "hello %s v2, you are at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
		v2.POST("/login/", func(c *gee.Context) {
			c.String(http.StatusOK, "Duplicated!")
		})
		v2.GET("/person/:pid/book/:bid", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"pid": c.Param("pid"),
				"bid": c.Param("bid"),
			})
		})
	}

	r.Run(":4444")
}
