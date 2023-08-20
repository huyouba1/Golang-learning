package store

type Uploader interface {
	Upload(bucketName string, objName string, localFilePath string) error
}
