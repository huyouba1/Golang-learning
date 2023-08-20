package tx

type tx struct {
	//client *oss.Client
	Endpoint        string `validata:"required"`
	AccessKeyId     string `validata:"required"`
	AccessKeySecret string `validata:"required"`
}

func NewTxUploader(endpoint, accessKeyId, accessKeySecret string) *tx {
	return &tx{
		Endpoint:        endpoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}
