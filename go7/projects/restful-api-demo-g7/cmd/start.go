package cmd

import (
	"gitee.com/go-learn/restful-api-demo-g7/apps"
	// 注册所有的实例
	_ "gitee.com/go-learn/restful-api-demo-g7/apps/all"

	"gitee.com/go-learn/restful-api-demo-g7/apps/host/http"
	"gitee.com/go-learn/restful-api-demo-g7/conf"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var (
	version  bool
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

		// 加载我们Host Service的实体类
		//service := impl.NewHostServiceImpl()

		// 注册HostService的实例到IOC
		// 采用: _ "gitee.com/go-learn/restful-api-demo-g7/apps/host/impl"  完成注册
		//apps.HostService = impl.NewHostServiceImpl()

		// 如何执行HostService的config方法
		// 因为apps.HostService 是一个host.Service的接口，并没有包含实例初始化（Config）方法
		apps.Init()
		// 通过 Host API handler 提供HTTP restful接口
		api := http.NewHostHTTPHandler()

		// 从IOC中获取依赖
		api.Config()

		// 提供一个 Gin router
		g := gin.Default()
		api.Registry(g)

		return g.Run(conf.C().App.HttpAddr())
	},
}

//还没有初始化logger实例

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
