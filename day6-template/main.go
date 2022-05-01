package main

import (
	"fmt"
	"github.com/siriusol/gee_study/day6-template/gee"
	"html/template"
	"net/http"
	"time"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	s1 := &Student{
		Name: "Alice",
		Age:  20,
	}
	s2 := &Student{
		Name: "Bob",
		Age:  30,
	}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": []*Student{s1, s2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "age",
			"now":   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":4444")
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

type Student struct {
	Name string
	Age  int8
}
