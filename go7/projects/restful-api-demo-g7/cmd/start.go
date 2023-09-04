package cmd

import (
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/apps"
	"gitee.com/go-learn/restful-api-demo-g7/protocol"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"os"
	"os/signal"
	"syscall"

	// 注册所有的实例
	_ "gitee.com/go-learn/restful-api-demo-g7/apps/all"

	"gitee.com/go-learn/restful-api-demo-g7/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var (
	confFile string
)

// 程序的启动时组装都在这里进行
// 1.
var StartCmd = &cobra.Command{
	Use:   "start",
	Long:  "启动 demo 后端api",
	Short: "启动 demo 后端api",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}

		// 初始化全局日志Logger
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 加载我们Host Service的实体类
		//service := impl.NewHostServiceImpl()

		// 注册HostService的实例到IOC
		// 采用: _ "gitee.com/go-learn/restful-api-demo-g7/apps/host/impl"  完成注册
		//apps.HostService = impl.NewHostServiceImpl()

		// 如何执行HostService的config方法
		// 因为apps.HostService 是一个host.Service的接口，并没有包含实例初始化（Config）方法
		apps.InitImpl()
		// 通过 Host API handler 提供HTTP restful接口
		//api := http.NewHostHTTPHandler()

		// 从IOC中获取依赖
		//api.Config()

		//// 提供一个 Gin router
		//g := gin.Default()
		//// 注册IoC的所有http handler
		//apps.InitGin(g)
		////api.Registry(g)
		//g.Run(conf.C().App.HttpAddr())
		svc := NewManager()

		ch := make(chan os.Signal, 1)
		// channel是一个复合数据结构，可以党城一个容器，自定义的struct make(chan int,1000) ,8bytes
		// 如果妈呀close gc 是不会回收的
		defer close(ch)
		// Go为了并发编程涉及的（CSP）
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go svc.WaitStop(ch)
		return svc.Start()
	},
}

func NewManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    zap.L().Named("CLI"),
	}
}

// 用于管理所有需要启动的服务
// 1. HTTP 服务的启动
type manager struct {
	http *protocol.HttpService
	l    logger.Logger
}

func (m *manager) Start() error {
	return m.http.Start()
}

// 处理来自外部的中断信号，比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.l.Infof("received signal %s", v)
			m.http.Stop()
		}
	}
}

// 问题：
// 1. http API、GRPC API需要启动，消息总线也需要监听，比如服务注册与配置，这些模块都是独立的
// 	  都需要在程序启动的时候进行启动，都写在start函数里面，start会膨胀到不易维护
// 2. 服务的优雅关闭？ 外部都会发送一个Terminal(中断)信号给程序，程序需要处理该信号
//	  需要实现程序优雅关闭的逻辑的处理，有先后顺序的（从外到内完成资源的释放逻辑处理）
// 		1. api层的关闭(HTTP\GRPC)
//		2. 消息总线关闭
//      3. 关闭数据库链接
//      4. 如果使用了注册猪心，最后才完成下线
//      5. 退出关闭

// 还没有初始化logger实例
// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 增加Config里面的日志配置，求配置全局Logger对象
	lc := conf.C().Log

	//  解析日志level
	//	DebugLevel: "debug",
	//	InfoLevel:  "info",
	//	WarnLevel:  "warning",
	//	ErrorLevel: "error",
	//	FatalLevel: "fatal",
	//	PanicLevel: "panic",
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	//  设置日志的级别 ,使用默认配置初始化Logger的全局配置
	zapConfig := zap.DefaultConfig()

	// 配置日志的Level级别
	zapConfig.Level = level

	// 程序每启动一次，不必要都生成一个新的日志文件
	zapConfig.Files.RotateOnStartup = false

	// 配置日志的输出方式
	switch lc.To {
	case conf.ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没把日志输出到文件中
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	// 配置日志的输出格式
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}

	// 把配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
