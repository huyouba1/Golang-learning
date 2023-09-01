package impl

import (
	"database/sql"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
	"gitee.com/go-learn/restful-api-demo-g7/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

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
