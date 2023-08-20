package aliyun

import (
	"cloud-station/store"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint,access_key,access_secret has one empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	AccessKey := "LTAI5t5ynxxLyiHdccpP8wkF"
	AccessSecret := "ovkOEKFV3VDIbVNC5LMOj9OudEHyOR"
	OssEndpoint := "oss-cn-beijing.aliyuncs.com"
	return NewAliOssStore(&Options{
		Endpoint:     OssEndpoint,
		AccessKey:    AccessKey,
		AccessSecret: AccessSecret,
	})
}

// AliOssStore 对象的构造函数
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	// 校验参数
	if err := opts.Validate(); err != nil {
		return nil, err
	}
	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client:   c,
		listener: NewDefaultProgressListener(),
	}, nil
}

type AliOssStore struct {
	// 阿里云 OSS client，私有变量，不运行外部
	client *oss.Client

	// 依赖listener的实现
	listener oss.ProgressListener
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 3. 上传文件到该bucket
	// ObjectKey 去掉路径合并到文件名称
	if err := bucket.PutObjectFromFile(objectKey, fileName, oss.Progress(s.listener)); err != nil {
		return err
	}
	// 4. 打印下载链接
	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL: %s \n", downloadURL)
	fmt.Println("请在1天之内下载。")
	return nil
}
