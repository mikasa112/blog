package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
	"v1/internal/middleware"
	"v1/internal/model"
	v1 "v1/internal/routers/api/v1"
	"v1/pkg"

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
	model.NewDBEngine()
	gin.SetMode(pkg.Sc.Mode)
	engine := gin.New()
	engine.Use(middleware.Logger())
	engine.Use(middleware.Translations())
	engine.Use(gin.Recovery())

	user := v1.NewUser()
	// apiv1 := engine.Group("api/v1")
	{

	}
	engine.POST("/api/login", user.Login)
	engine.POST("/api/generate", user.GenerateUser)

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
