package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gin_web_api/models"
	"gin_web_api/pkg/logging"
	"gin_web_api/pkg/rsa"
	"gin_web_api/pkg/setting"
	"gin_web_api/pkg/shutdown"
	"gin_web_api/routers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	rsa.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
func Run() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	logging.Info("[info] start http server listening ", endPoint)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logging.Fatal("server shutdown err", zap.Error(err))
			}
		},
	)
}
