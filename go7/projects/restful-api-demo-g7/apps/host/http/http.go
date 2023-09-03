package http

import (
	"gitee.com/go-learn/restful-api-demo-g7/apps"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
	"github.com/gin-gonic/gin"
)

//	func NewHostHTTPHandler() *Handler {
//		return &Handler{}
//	}
var handler = &Handler{}

// 通过写一个实例类，把内部的接口通过HTTP协议暴露出去
// 所以需要依赖内部接口的实现
// 该实体类，会实现Gin的http handler
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	//if apps.HostService == nil {
	//	panic("Dependence host service required")
	//}
	//// 从 IOC 里面获取HostService的实例对象
	//h.svc = apps.HostService
	h.svc = apps.GetImpl(host.AppName).(host.Service)
}

// 只是完成了Http Handler的注册
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
	r.GET("/hosts", h.queryHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

// 完成http Handler的注册
func init() {
	apps.RegistryGin(handler)
}
