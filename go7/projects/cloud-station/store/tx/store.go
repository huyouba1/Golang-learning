package tx

func NewTxOssStore() *txOssStore {
	return &txOssStore{}
}

type txOssStore struct {
}

func (s *txOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
