package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"mycloud-station/store"
	"mycloud-station/store/aliyun"
	"net"
	"os"
	"path"
	"strings"
	"time"
)

var (
	uploadFilePath string
	bucketName     string
	bucketEndpoint string
)

const (
	defaultBuckName = "images-huyouba1"
	defaultEndpoint = "oss-cn-beijing.aliyuncs.com"
	defaultALIAK    = "LTAI5t5ynxxLyiHdccpP8wkF"
	defaultALISK    = "ovkOEKFV3VDIbVNC5LMOj9OudEHyOR"
)

var UploadCMD = &cobra.Command{
	Use:   "upload",
	Short: "上传文件到中转站",
	Long:  "upload 上传文件到中转站",
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getprovider()
		if err != nil {
			return err
		}
		if uploadFilePath == "" {
			return fmt.Errorf("upload file path is missing")
		}
		// 为了防止文件都堆在一个文件夹里面 无法查看
		// 我们采用日期进行编码
		day := time.Now().Format("20060102")

		// 为了防止不同用户同一时间上传相同的文件
		// 我们采用用户的主机名作为前置
		hn, err := os.Hostname()
		if err != nil {
			ipAddr := getOutBindIp()
			if ipAddr == "" {
				hn = "unknown"
			} else {
				hn = ipAddr
			}
		}
		fn := path.Base(uploadFilePath)
		ok := fmt.Sprintf("%s/%s/%s", day, hn, fn)
		err = p.Upload(bucketName, ok, uploadFilePath)
		if err != nil {
			return err
		}
		return nil
	},
}

func getOutBindIp() string {
	conn, err := net.Dial("udp", "baidu.com")
	if err != nil {
		return ""
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")
	if len(ip) == 0 {
		return ""
	}
	return ip[0]
}

func getprovider() (p store.Uploader, err error) {
	switch ossProvider {
	case "aliyun":
		fmt.Printf("上传云商: 阿里云[%s]\n", defaultEndpoint)
		if aliAccessKey == "" {
			aliAccessKey = defaultALIAK
		}
		if aliSecretKey == "" {
			aliSecretKey = defaultALISK
		}
		fmt.Printf("上传用户: %s\n", aliAccessKey)
		getSecretKeyFromInput()
		p, err = aliyun.NewUploader(bucketEndpoint, aliAccessKey, aliSecretKey)
		return
	case "tx":
		return nil, fmt.Errorf("not impl")
	case "aws":
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknown oss privier options [aliyun/qcloud]")
	}
}

func getSecretKeyFromInput() {
	fmt.Printf("请输入access secret key: ")
	fmt.Scanln(&aliSecretKey)

}

func init() {
	UploadCMD.PersistentFlags().StringVarP(&uploadFilePath, "uploadFilePath", "f", "", "upload file path")
	UploadCMD.PersistentFlags().StringVarP(&bucketName, "bucketName", "b", defaultBuckName, "upload oss bucket name")
	UploadCMD.PersistentFlags().StringVarP(&bucketEndpoint, "bucketEndpoint", "e", "", "upload oss endpoint")
	RootCMD.AddCommand(UploadCMD)
}
