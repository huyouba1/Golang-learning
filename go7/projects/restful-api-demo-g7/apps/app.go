package apps

import (
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
	"github.com/gin-gonic/gin"
)

// IOC容器层：管理所有的服务的实例

// 1. HostService的实例必须注册过来，HostService才会有具体的实例对象，服务启动的时候完成注册
// 2. HTTP暴露模块，依赖于IOC里面的HostService
var (
	// 40 service  写40个定义？
	// 使用interface{}  + 断言进行抽象
	HostService host.Service

	// 维护当前所有的服务
	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

func RegistryImpl(svc ImplService) {
	// 服务实例注册到svcs map 当中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("Service %s has registried ", svc.Name()))
	}
	implApps[svc.Name()] = svc
	// 根据对象满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 如果制定了具体类型，就导致每次增加一个类型就多一个Get方法
//func GetHostImpl(name string) host.Service

// Get 一个Impl服务的实例： implApps拿
// 返回一个对象，任何类型都可以，使用时由使用方进行断言
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}
	return nil
}

func RegistryGin(svc GinService) {
	// 服务实例注册到svcs map 当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("Service %s has registried ", svc.Name()))
	}
	ginApps[svc.Name()] = svc
}

// 用于初始化 注册到IoC容器里面的所有服务
func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

// 已经加载完成的Gin App有哪些
func LoadGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}
	return
}

func InitGin(r gin.IRouter) {
	// 先初始化好所有的对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成http handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

// 注册Gin编写的Handler
// 比如编写了Http服务A，只需要实现Registry方法，就能吧Handler注册给Root Router
type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}
