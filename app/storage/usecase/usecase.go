package usecase

import storageInterfaces "key-value-storage/app/storage/interfaces"

type storageUseCase struct {
	repository storageInterfaces.StorageRepository
}

func NewStorageUseCase(repository storageInterfaces.StorageRepository) *storageUseCase {
	return &storageUseCase{repository: repository}
}

func (u *storageUseCase) CreateOrUpdateKey(key, value string) error {
	return u.repository.CreateOrUpdateKey(key, value)
}

func (u *storageUseCase) DeleteKey(key string) error {
	return u.repository.DeleteKey(key)
}

func (u *storageUseCase) GetKeyValue(key string) (string, error) {
	return u.repository.GetKeyValue(key)
}

func (u *storageUseCase) GetListOfKeys() ([]string, error) {
	list, err := u.repository.GetListOfKeys()

	if len(list) == 0 {
		list = []string{}
	}

	return list, err
}
