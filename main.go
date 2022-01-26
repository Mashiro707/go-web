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

	r.GET("/hello/:name", func(c *gow.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gow.Context) {
		c.JSON(http.StatusOK, gow.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
