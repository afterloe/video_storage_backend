package repositories

import (
	"errors"
	"video_storage/tools"
)

var MemoryStorageRepository *memoryStorageRepository

var storage = make(map[string]map[string]interface{})

const (
	UserCacheDataType = "user"
)

type memoryStorageRepository struct {
}

func (*memoryStorageRepository) Get(dataType string, key string) (interface{}, error) {
	if nil == storage[dataType] {
		return nil, errors.New("没有这个类型的存储")
	}
	if val := storage[dataType][key]; nil == val {
		return nil, errors.New(dataType + " 没有 " + key + "对应的值")
	} else {
		return val, nil
	}
}

func (*memoryStorageRepository) Set(dataType string, value interface{}) string {
	key := tools.GeneratorUUID()
	if dataType == UserCacheDataType {
		if nil == storage[UserCacheDataType] {
			storage[UserCacheDataType] = make(map[string]interface{})
		}
		storage[UserCacheDataType][key] = value
	}

	return key
}
