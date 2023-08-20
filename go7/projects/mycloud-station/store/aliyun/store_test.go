package aliyun_test

import (
	"github.com/stretchr/testify/assert"
	"mycloud-station/store/aliyun"
	"testing"
)

var (
	bucketName    = "images-huyouba1"
	objName       = "store.go"
	localFilePath = "store.go"
	endpoint      = "oss-cn-beijing.aliyuncs.com"
	ak            = "LTAI5t5ynxxLyiHdccpP8wkF"
	sk            = "ovkOEKFV3VDIbVNC5LMOj9OudEHyOR"
)

func TestUpload(t *testing.T) {
	should := assert.New(t)
	uploader, err := aliyun.NewUploader(endpoint, ak, sk)
	err = uploader.Upload(bucketName, objName, localFilePath)
	should.NoError(err)
}
