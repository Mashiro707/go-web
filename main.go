package main

import (
	"gow"
	"net/http"
)

func main() {
	r := gow.New()
	r.GET("/hello", func(c *gow.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/", func(c *gow.Context) {
		c.JSON(http.StatusOK, gow.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9001")
}
