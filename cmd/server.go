package main

import (
	"github.com/Sn0wo2/MeowMetrics/Backend/framework"
	"github.com/gin-gonic/gin"
)

func main() {
	e := framework.NewGin()
	e.Any("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, MeowMetrics~")
	})
	err := e.Start()
	if err != nil {
		panic(err)
	}
}
