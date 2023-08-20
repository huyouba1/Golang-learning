package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

// 修改变量来控制运行逻辑
var (
	// 程序内置
	AccessKey    = "xxx"
	AccessSecret = "xxxx"
	OssEndpoint  = "oss-cn-beijing.aliyuncs.com"

	// 默认配置
	BucketName = "images-xxxx"

	// 用户需要传递的参数
	// 期望用户自己输出(CLI/GUI)
	UploadFile = ""
	help       = false
)

// 实现文件上传的函数
func upload(filepath string) error {
	// 1.实例化client
	client, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		return err
	}
	// 2. 获取bucket对象
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		return err
	}
	// 3. 上传文件到该bucket
	// ObjectKey 去掉路径合并到文件名称
	if err := bucket.PutObjectFromFile(filepath, filepath); err != nil {
		return err
	}
	// 4. 打印下载链接
	downloadURL, err := bucket.SignURL(filepath, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL: %s \n", downloadURL)
	fmt.Println("请在1天之内下载。")
	return nil
}

// 参数合法性检查
func validate() error {
	if OssEndpoint == "" || AccessKey == "" || AccessSecret == "" {
		return fmt.Errorf("endpoint,access_key,access_secret has one empty")
	}

	if UploadFile == "" {
		return fmt.Errorf("upload file path required")
	}
	return nil
}

func loadParams() {
	flag.BoolVar(&help, "h", false, "打印帮助信息")
	flag.StringVar(&UploadFile, "f", "", "需上传文件的名称")
	flag.Parse()

	// 判断CLI 是否需要打印Help信息
	if help {
		usage()
		os.Exit(0)
	}
}

// 打印使用说明
func usage() {
	// 1. 打印一些描述信息
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <upload_file_path>
Options:
    `)
	// 2. 打印有哪些参数可以使用，就像-f
	flag.PrintDefaults()
}

func main() {
	// 参数加载
	loadParams()
	// 参数验证
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常\n")
		usage()
		os.Exit(1)
	}
	if err := upload(UploadFile); err != nil {
		fmt.Printf("上传文件异常：%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("文件: %s 上传完成\n", UploadFile)

}
