package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	vers         bool
	ossProvider  string
	aliAccessKey string
	aliSecretKey string
)

var RootCMD = &cobra.Command{
	Use:   "mycloud-station-cli",
	Short: "mycloud-station-cli 文件中转",
	Long:  "mycloud-station-cli 阿里云文件中转",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println("My cloud station's version is v1.0.0")
			return nil
		}
		return errors.New("no flags find")
	},
}

func init() {
	RootCMD.PersistentFlags().BoolVarP(&vers, "version", "v", false, "My cloud station 版本信息")
	RootCMD.PersistentFlags().StringVarP(&ossProvider, "provider", "p", "aliyun", "the oss provider [aliyun/qcloud]")
	RootCMD.PersistentFlags().StringVarP(&aliAccessKey, "accesskey", "a", "", "the ali oss access id")
	RootCMD.PersistentFlags().StringVarP(&aliSecretKey, "secretkey", "s", "", "the ali oss secretkey")
}
