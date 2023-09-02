package apps

import "gitee.com/go-learn/restful-api-demo-g7/apps/host"

// IOC容器层：管理所有的服务的实例

// 1. HostService的实例必须注册过来，HostService才会有具体的实例对象，服务启动的时候完成注册
// 2. HTTP暴露模块，依赖于IOC里面的HostService
var (
	HostService host.Service
)
