package apps

import (
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
)

// IOC容器层：管理所有的服务的实例

// 1. HostService的实例必须注册过来，HostService才会有具体的实例对象，服务启动的时候完成注册
// 2. HTTP暴露模块，依赖于IOC里面的HostService
var (
	HostService host.Service

	// 维护当前所有的服务
	svcs = map[string]Service{}
)

func Registry(svc Service) {
	// 服务实例注册到svcs map 当中
	if _, ok := svcs[svc.Name()]; ok {
		panic(fmt.Sprintf("Service %s has registried ", svc.Name()))
	}
	svcs[svc.Name()] = svc
	// 根据对象满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 用于初始化 注册到IoC容器里面的所有服务
func Init() {
	for _, v := range svcs {
		v.Config()
	}
}

type Service interface {
	Config()
	Name() string
}
