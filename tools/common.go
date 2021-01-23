package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"
	"video_storage/config"
)

// 生成UUID
func GeneratorUUID() string {
	code, _ := uuid.NewV4()
	return strings.ToLower(strings.Replace(code.String(), "-", "", -1))
}

// GetTime 获取格式化时间文本
func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 字段检测
func CheckStr(items ...string) error {
	for _, it := range items {
		vt := reflect.TypeOf(it)
		if 0 == len(it) {
			return errors.New(vt.String())
		}
	}
	return nil
}

// 获取文件流
func GeneratorLogStream(logPath, suffix string) io.Writer {
	fullLogPath := fmt.Sprintf("%s/%s.%s_%s.log", logPath, config.Instance.Common.ServerName, suffix, time.Now().Format("20060102"))
	fmt.Printf("[info] will write %s log in %s \r\n", suffix, fullLogPath)
	stream, err := os.OpenFile(fullLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if nil != err {
		fmt.Printf("[error] can't open this %s log in %s. \r\n", suffix, fullLogPath)
		fmt.Println("[info] will return to the default.")
		return os.Stdout
	}
	return stream
}

func ReadFileAsString(filePath string) string {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Error("can't read this file " + filePath)
		return ""
	}
	return string(f)
}

func InterfaceTransformation(source interface{}, pointer interface{}) error {
	jsonByte, _ := json.Marshal(source)
	return json.Unmarshal(jsonByte, pointer)
}
