package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

func NewOssProgressListener() *OssProgressListener {
	return &OssProgressListener{}
}

type OssProgressListener struct {
	bar *progressbar.ProgressBar
}

func (p *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	//fmt.Println(event.EventType, event.TotalBytes, event.RwBytes)
	switch event.EventType {
	case oss.TransferStartedEvent:
		p.bar = progressbar.DefaultBytes(event.TotalBytes, "文件上传中...")
	case oss.TransferDataEvent:
		p.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Println("上传完成")
	case oss.TransferFailedEvent:
		fmt.Println("上传失败")
	}
}
