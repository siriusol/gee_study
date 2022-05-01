package main

import (
	"github.com/siriusol/gee_study/day7-recovery/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello")
	})
	r.GET("/panic", func(c *gee.Context) {
		c.String(http.StatusOK, []string{}[1])
	})

	r.Run(":4444")
}
