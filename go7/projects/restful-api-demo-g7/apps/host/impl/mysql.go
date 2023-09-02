package impl

import (
	"database/sql"
	"gitee.com/go-learn/restful-api-demo-g7/apps"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
	"gitee.com/go-learn/restful-api-demo-g7/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查
//var _ host.Service = (*HostServiceImpl)(nil)

// 这样写，会造成 config.C()还没有准备好，会触发panic
// var impl = NewHostServiceImpl()

// 把对象的注册和对象的初始化这两个逻辑独立出来
var impl = &HostServiceImpl{}

// NewHostServiceImpl 保证该函数之前，全局conf已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host Service 服务的子logger
		// 封装的zap让其满足Logger接口
		// 为什么要封装z:
		// 1. Logger全局实例
		// 2. Logger Level 的动态调整，Logrus不支持Level动态调整
		// 3. 加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	// host service 的具体实现
	l  logger.Logger
	db *sql.DB
}

// 只需要保证，全局对象Config和全局Logger已经加载完成就可以
func (i *HostServiceImpl) Config() {
	// Host Service 服务的子logger
	// 封装的zap让其满足Logger接口
	// 为什么要封装z:
	// 1. Logger全局实例
	// 2. Logger Level 的动态调整，Logrus不支持Level动态调整
	// 3. 加入日志轮转功能的集合
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()
}

// 返回服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

// _ import app 自动执行注册逻辑
func init() {
	// 对象注册到IoC层
	apps.RegistryImpl(impl)
	//apps.HostService = impl
}

// 之前都是在Start的时候，手动吧服务的实现注册到IoC层
// 注册HostService的实例到IOC
// apps.HostService = impl.NewHostServiceImpl()

// mysql 的驱动加载的实现方式
// sql 这个库，是一个框架，驱动是引入依赖的时候加载的
// 我们把app模块，比作一个驱动，ioc比作框架
// _ import app，该app就注册到ioc层
