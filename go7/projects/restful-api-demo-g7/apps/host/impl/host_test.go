package impl_test

import (
	"context"
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host"
	"gitee.com/go-learn/restful-api-demo-g7/apps/host/impl"
	"gitee.com/go-learn/restful-api-demo-g7/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-02"
	ins.Name = "test"
	ins.Region = "cn-hangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func TestQuery(t *testing.T) {
	should := assert.New(t)
	req := host.NewQueryHostRequst()
	req.Keywords = "接口测试"
	set, err := service.QueryHost(context.Background(), req)
	if should.NoError(err) {
		for i := range set.Items {
			fmt.Println(set.Items[i].Id)
		}
	}
}

func TestDescribe(t *testing.T) {
	should := assert.New(t)
	req := host.NewDescribeHostRequestWithId("ins-09")
	ins, err := service.DescribeHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}

func init() {
	// 测试用例的配置文件
	//err := conf.LoadConfigFromToml("/Users/v_zhenxiyao/Desktop/Golang/go7/projects/restful-api-demo-g7/etc/demo.toml")
	err := conf.LoadConfigFromToml("../../../etc/demo.toml")
	if err != nil {
		panic(err)
	}

	// 需要初始化全局Logger
	// 为什么不设计为默认打印，因为性能
	fmt.Println(zap.DevelopmentSetup())
	// host service 的具体实现
	service = impl.NewHostServiceImpl()
}
