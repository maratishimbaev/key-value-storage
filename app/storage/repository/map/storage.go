package storageMap

type storageRepository struct {
	storage map[string]string
}

func NewStorageRepository() *storageRepository {
	return &storageRepository{storage: map[string]string{}}
}
