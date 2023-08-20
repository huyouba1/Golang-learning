package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-playground/validator"
	"mycloud-station/store"
)

func NewUploader(endpoint, accessKeyId, accessKeySecret string) (store.Uploader, error) {
	p := &aliyun{
		Endpoint:        endpoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		listenr:         NewOssProgressListener(),
	}
	if err := p.validata(); err != nil {
		return nil, err
	}
	return p, nil
}

type aliyun struct {
	//client *oss.Client
	Endpoint        string `validata:"required"`
	AccessKeyId     string `validata:"required"`
	AccessKeySecret string `validata:"required"`
	listenr         oss.ProgressListener
}

var (
	validata = validator.New()
)

func (a *aliyun) validata() error {
	return validata.Struct(a)
}

func (a *aliyun) Upload(bucketName string, objName string, localFilePath string) error {
	// 2. 获取bucket对象
	bucket, err := a.ListBucket(bucketName)
	if err != nil {
		return err
	}
	// 3. 上传文件到该bucket
	// ObjectKey 去掉路径合并到文件名称
	if err := bucket.PutObjectFromFile(objName, localFilePath, oss.Progress(a.listenr)); err != nil {
		return err
	}
	// 4. 打印下载链接
	downloadURL, err := bucket.SignURL(objName, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL: %s \n", downloadURL)
	fmt.Println("请在1天之内下载。")
	return nil
}

func (a *aliyun) ListBucket(bucketName string) (*oss.Bucket, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("The bucketName is empty!")
	}

	client, err := oss.New(a.Endpoint, a.AccessKeyId, a.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
