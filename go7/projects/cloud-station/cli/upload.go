package cli

import (
	"cloud-station/store"
	"cloud-station/store/aliyun"
	"cloud-station/store/aws"
	"cloud-station/store/tx"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	ossProvier   string
	ossEndpoint  string
	accessKey    string
	accessSecret string
	bucketName   string
	uploadFile   string
)

const (
	default_ak = "xx"
	default_sk = "xx"
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvier {
		case "aliyun":

			aliOpts := &aliyun.Options{
				Endpoint:     ossEndpoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			}
			setAliDefault(aliOpts)
			uploader, err = aliyun.NewAliOssStore(aliOpts)
		case "tx":
			uploader = tx.NewTxOssStore()
		case "aws":
			uploader = aws.NewAwsOssStore()
		default:
			return fmt.Errorf("not support oss storage provider")

		}
		if err != nil {
			return err
		}
		// 使用upload来上传文件
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

func setAliDefault(options *aliyun.Options) {
	if options.AccessKey == "" {
		options.AccessKey = default_ak
	}
	if options.AccessSecret == "" {
		options.AccessSecret = default_sk
	}
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvier, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provider endpoint")
	f.StringVarP(&bucketName, "bucket_name", "b", "images-huyouba1", "oss storage provider bucket name")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	f.StringVarP(&accessKey, "access_key", "k", "", "oss storage provider ak")
	f.StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provider sk")
	RootCmd.AddCommand(UploadCmd)
}
