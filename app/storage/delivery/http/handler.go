package storageHttp

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	storageInterfaces "key-value-storage/app/storage/interfaces"
	"key-value-storage/app/storage/models"
	"net/http"
)

type storageHandler struct {
	useCase storageInterfaces.StorageUseCase
}

func NewStorageHandler(useCase storageInterfaces.StorageUseCase) *storageHandler {
	return &storageHandler{useCase: useCase}
}

func (h *storageHandler) CreateOrUpdateKey(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var keyValue models.KeyValue

	err = json.Unmarshal(body, &keyValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.useCase.CreateOrUpdateKey(keyValue.Key, keyValue.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *storageHandler) DeleteKey(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	err := h.useCase.DeleteKey(key)
	switch true {
	case errors.Is(err, storageInterfaces.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	case err == nil:
		w.WriteHeader(http.StatusNoContent)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *storageHandler) GetKeyValue(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	var valueModel models.Value

	var err error
	valueModel.Value, err = h.useCase.GetKeyValue(key)
	switch true {
	case errors.Is(err, storageInterfaces.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	case err == nil:
		w.WriteHeader(http.StatusOK)
		jsonValue, _ := json.Marshal(valueModel)
		w.Write(jsonValue)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *storageHandler) GetListOfKeys(w http.ResponseWriter, r *http.Request) {
	list, err := h.useCase.GetListOfKeys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsonList, _ := json.Marshal(list)
	w.Write(jsonList)
}
