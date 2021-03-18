package storageMap

import storageInterfaces "key-value-storage/app/storage/interfaces"

type storageRepository struct {
	storage map[string]string
}

func NewStorageRepository() *storageRepository {
	return &storageRepository{storage: map[string]string{}}
}

func (r *storageRepository) CreateOrUpdateKey(key, value string) error {
	r.storage[key] = value

	return nil
}

func (r *storageRepository) DeleteKey(key string) error {
	if _, ok := r.storage[key]; !ok {
		return storageInterfaces.ErrNotFound
	}

	delete(r.storage, key)

	return nil
}

func (r *storageRepository) GetKeyValue(key string) (string, error) {
	value, ok := r.storage[key]
	if !ok {
		return "", storageInterfaces.ErrNotFound
	}

	return value, nil
}

func (r *storageRepository) GetListOfKeys() ([]string, error) {
	var list []string

	for key, _ := range r.storage {
		list = append(list, key)
	}

	return list, nil
}
