package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("port", 3000, "webserver port")
	mode := flag.String("mode", "release", "dev or release")

	flag.Parse()

	if *mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	app.Static("/public", "./public")

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	app.Run(fmt.Sprintf(":%v", *port))
}
