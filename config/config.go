package config

import (
	"gopkg.in/ini.v1"
	_ "io/ioutil"
	"reflect"
)

var Instance *ApplicationConfig

// 读取ini文件
func ReadInICfgFile(pathOfInI string) {
	Instance = &ApplicationConfig{}
	iniCfg, err := ini.Load(pathOfInI)
	if nil != err {
		return
	}
	instanceValue := reflect.ValueOf(Instance).Elem()
	instanceType := reflect.TypeOf(Instance).Elem()
	for i := 0; i < instanceType.NumField(); i++ {
		field := instanceType.Field(i)
		sectionName := field.Tag.Get("ini")
		section, err := iniCfg.GetSection(sectionName)
		if nil != err {
			continue
		}
		structFields := instanceValue.Field(i).Type()
		for j := 0; j < structFields.NumField(); j++ {
			key := structFields.Field(j).Tag.Get("ini")
			value, err := section.GetKey(key)
			if nil != err {
				continue
			}
			instanceValue.Field(i).Field(j).Set(reflect.ValueOf(value.String()))
		}
	}
}
