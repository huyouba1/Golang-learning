package impl

import (
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
)
import "context"

// 业务处理层（controller层）
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 直接打印日志
	//i.l.Named("Create").Debug("create host")
	//i.l.Info("create host")
	//// 带Format的日志打印,fmt.Sprintf()
	//i.l.Debugf("create host %s", ins.Name)
	//// 携带额外的meta数据，常用于Trace系统
	//i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")

	// 校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 默认值填充
	ins.InjectDefault()

	// 有dao模块负责把对象入库
	err := i.save(ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequst) (*host.HostSet, error) {

	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequst) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequst) (*host.Host, error) {
	return nil, nil
}
