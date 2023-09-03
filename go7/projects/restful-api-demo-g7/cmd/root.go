package cmd

import (
	"errors"
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/version"
	"github.com/spf13/cobra"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Long:  "demo-api后端api",
	Short: "demo-api后端api",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return errors.New("no flags find")
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&vers, "version", "v", false, "print demo-api  版本信息")
}
