package aliyun_test

import (
	"cloud-station/store"
	"cloud-station/store/aliyun"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	uploader store.Uploader
)

var (
	BucketName = "images-huyouba1"
)

// Aliyun Oss Store Upload 测试用例
func TestUpload(t *testing.T) {
	// 使用assert编写测试用例的断言语句
	// 通过New获取一个断言实例
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		// 没有error开启下一个步骤
		t.Log("upload ok")
	}
}

func TestUploadError(t *testing.T) {
	// 使用assert编写测试用例的断言语句
	// 通过New获取一个断言实例
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_testxxx.go")
	should.Error(err, "open store_testxxx.go: no such file or directory")

}

// 通过init编写uploader实例化逻辑
func init() {
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali
}
