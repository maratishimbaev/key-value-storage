package storageInterfaces

type StorageUseCase interface {
	CreateOrUpdateKey(key, value string) error
	DeleteKey(key string) error
	GetKeyValue(key string) (string, error)
	GetListOfKeys() ([]string, error)
}
