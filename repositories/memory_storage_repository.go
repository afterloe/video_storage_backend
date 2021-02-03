package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"video_storage/model"
	"video_storage/tools"
)



var storage = make(map[string]map[string]interface{})

const (
	UserCacheDataType = "user"
	fileName = "./.memoryStorage"
)

type memoryStorageRepository struct {
}

func (*memoryStorageRepository) LoadStatusFile () {
	content := tools.ReadFileAsString(fileName)
	if "" == content {
		return
	}
	var a map[string]interface{}
	_ = json.Unmarshal([]byte(content), &a)
	for dataType, instance := range a {
		if UserCacheDataType == dataType {
			storage[UserCacheDataType] = make(map[string]interface{})
			for k, v := range instance.(map[string]interface{}) {
				var user model.User
				tools.InterfaceTransformation(v, &user)
				storage[UserCacheDataType][k] = &user
			}
		}
	}
}

func (*memoryStorageRepository) SaveStatus() {
	_ = os.Remove(fileName)
	file, _ := os.Create(fileName)
	defer file.Close()
	JSONByte, _ := json.Marshal(storage)
	_ = ioutil.WriteFile(fileName, JSONByte, 0666)
}

func (*memoryStorageRepository) Del(dataType string, key string) {
	if nil == storage[dataType] {
		return
	}
	delete(storage[dataType], key)
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

func (*memoryStorageRepository) GetAllTypeValue(dataType string) map[string]interface{} {
	if nil == storage[dataType] {
		return nil
	}
	return storage[dataType]
}

func (that *memoryStorageRepository) Set(dataType string, value interface{}) string {
	key := tools.GeneratorUUID()
	if dataType == UserCacheDataType {
		if nil == storage[UserCacheDataType] {
			storage[UserCacheDataType] = make(map[string]interface{})
		}
		storage[UserCacheDataType][key] = value
	}
	that.SaveStatus()
	return key
}
