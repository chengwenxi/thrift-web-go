package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/chengwenxi/thrift-web-go/go/server"
)

func main() {
	go server.Server()

	r := gin.New()
	r.Static("gen-js","./gen-js")
	r.LoadHTMLGlob("./index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.Run()


}
