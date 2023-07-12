package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"v1/internal/middleware"
	"v1/pkg"
	"v1/pkg/app"
	"v1/pkg/err"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

// 启动
func (s *Server) Run() {
	//读取配置
	pkg.ReadConfig()
	boot()
}

func boot() {
	gin.SetMode(pkg.Sc.Mode)
	engine := gin.New()
	engine.Use(middleware.Logger())
	engine.Use(gin.Recovery())
	engine.GET("/", func(ctx *gin.Context) {
		app.NewResponse(ctx).ErrTo(err.InvalidParams)
	})
	srv := &http.Server{
		Addr:    pkg.Sc.Port,
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			pkg.Log.Fatal("Listen", zap.String("error", err.Error()))
		}
	}()
	pkg.Log.Info("Server Started")
	//优雅的关闭服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	pkg.Log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		pkg.Log.Fatal("Server Shutdown", zap.String("error", err.Error()))
	}
	pkg.Log.Info("Server exiting")
}
