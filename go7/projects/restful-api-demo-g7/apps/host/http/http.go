package http

import (
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
	"github.com/gin-gonic/gin"
)

func NewHostHTTPHandler(svc host.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// 通过写一个实例类，把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类，会实现Gin的http handler
type Handler struct {
	svc host.Service
}

// 只是完成了Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}
