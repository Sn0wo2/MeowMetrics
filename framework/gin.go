package framework

import (
	"net/http"
	"time"

	"github.com/Sn0wo2/MeowMetrics/Backend/debug"
	"github.com/gin-gonic/gin"
)

type GinF struct {
	*gin.Engine
}

func NewGin() *GinF {
	if !debug.IsDebugging() {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(gin.Recovery())

	engine.HandleMethodNotAllowed = true

	return &GinF{
		Engine: engine,
	}
}

func (gf *GinF) Start() error {
	server := &http.Server{
		Addr:              ":3000",
		Handler:           gf,
		ReadHeaderTimeout: 30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	return server.ListenAndServe()
}
