package example_test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

var (
	// 全局client实例，在包加载的时候初始化(init)
	client *oss.Client
)

var (
	AccessKey    = os.Getenv("ALI_ak")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

// 测试阿里云oss SDK BucketList接口
func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

// 测试阿里云oss SDK PutObjectFromFile接口
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		t.Log(err)
	}

	// 上传文件到bucket中
	// 云商ossserver会根据key的路径结构自动创建目录
	// objectKey 上传到bucket里面的对象的名称，包含路径
	// mydir/test.go ossserver会自动创建一个mydir目录
	// 把当前的文件上传到了mydir下
	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}

// 初始化一个oss的client，等下给其他所有测试用例使用
func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	fmt.Println(OssEndpoint)
	if err != nil {
		panic(err)
	}
	client = c
}
