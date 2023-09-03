package protocol

import (
	"context"
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/apps"
	"gitee.com/go-learn/restful-api-demo-g7/conf"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"
)

// HttpService 构造函数
func NewHttpService() *HttpService {
	// new gin router 实例，并没有加载Handler
	r := gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAddr(),
		Handler:           r,
	}
	return &HttpService{
		server: server,
		l:      zap.L().Named("HTTP Service"),
		r:      r,
	}
}

type HttpService struct {
	server *http.Server
	l      logger.Logger
	r      gin.IRouter
}

func (s *HttpService) Start() error {
	// 加载Handler,把所有的模块的handler注册给了Gin Router
	apps.InitGin(s.r)

	// 已加载App的日志信息
	apps := apps.LoadGinApps()
	s.l.Infof("loaded gin apps :%v", apps)

	// 该操作是阻塞点，简单端口，等待请求
	// 如果服务的正常的关闭
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service stopped success")
			return nil
		}
		return fmt.Errorf("start service err, %s", err.Error())
	}
	return nil
}

func (s *HttpService) Stop() {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("shut down http service error, %s", err)
	}
}
