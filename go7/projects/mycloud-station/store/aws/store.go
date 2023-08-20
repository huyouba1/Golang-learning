package aws

type aws struct {
	Endpoint        string `validata:"required"`
	AccessKeyId     string `validata:"required"`
	AccessKeySecret string `validata:"required"`
}

func NewAwsUploader(endpoint, accessKeyId, accessKeySecret string) *aws {
	return &aws{
		Endpoint:        endpoint,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}
