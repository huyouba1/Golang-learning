package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

//var (
//	//ossProvier string
//	//version bool
//)

var RootCmd = &cobra.Command{
	Use:   "demo",
	Long:  "demo后端api",
	Short: "demo后端api",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flags find")
	},
}

//
//func init() {
//	f := RootCmd.PersistentFlags()
//	f.BoolVarP(&version, "version", "v", false, "cloud station 版本信息")
//}
