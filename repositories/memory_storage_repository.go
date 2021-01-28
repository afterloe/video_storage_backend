package repositories

import "video_storage/tools"

var MemoryStorageRepository *memoryStorageRepository

var storage = make(map[string]map[string]interface{})
const (
	UserCacheDataType = "user"
)

type memoryStorageRepository struct {
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