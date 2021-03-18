package storageInterfaces

type StorageRepository interface {
	CreateOrUpdateKey(key, value string) error
	DeleteKey(key string) error
	GetKeyValue(key string) (string, error)
	GetListOfKeys() ([]string, error)
}
