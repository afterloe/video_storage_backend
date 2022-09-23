package sdk

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type sqliteSDK struct {
	instance *sql.DB
}

const (
	driverName = "sqlite3"
)

func initSQLiteSDK(pathOfDBFile string) (*sqliteSDK, error) {
	logrus.Info("连接SQLit3数据库 ... ...")
	logrus.Infof("数据库路径:%s", pathOfDBFile)
	instance, err := sql.Open(driverName, pathOfDBFile)
	if nil != err {
		logrus.Error("初始化sqlite3 sdk失败.")
		logrus.Error(err)
		return nil, err
	}
	if err = instance.Ping(); nil != err {
		logrus.Error("连接SQLite3数据库失败，请检查文件权限.")
		logrus.Error(err)
		return nil, err
	}
	return &sqliteSDK{instance: instance}, nil
}

func (*sqliteSDK) GetContext() context.Context {
	return context.Background()
}

// 查询一条
func (that *sqliteSDK) QueryOne(callback func(sql.Row), sqlStr string, args ...interface{}) {
	if row := that.instance.QueryRowContext(that.GetContext(), sqlStr, args...); nil != row {
		if nil != callback {
			callback(*row)
		}
	} else {
		logrus.Errorf("执行SQL出现异常: %v -> %v", sqlStr, args)
	}
}

// 查询多条
func (that *sqliteSDK) Query(callback func(sql.Rows), sqlStr string, args ...interface{}) {
	logrus.Debugf("执行SQL: %v -> %v", sqlStr, args)
	if rows, err := that.instance.QueryContext(that.GetContext(), sqlStr, args...); nil != err {
		logrus.Errorf("执行SQL出现异常: %v -> %v", sqlStr, args)
	} else {
		if nil != callback {
			callback(*rows)
		}
	}
}

// 执行SQL
func (that *sqliteSDK) Execute(callback func(sql.Result), sqlStr string, args ...interface{}) {
	if result, err := that.instance.ExecContext(that.GetContext(), sqlStr, args...); nil != err {
		logrus.Errorf("执行SQL出现异常: %v -> %v", sqlStr, args)
		logrus.Error(err)
	} else {
		if nil != callback {
			callback(result)
		}
	}
}

func (that *sqliteSDK) ResultToModelBySQL(sql string, value interface{}) []interface{} {
	idx := strings.Index(sql, "FROM")
	columns := strings.Split(sql[6:idx], ",")
	for i, column := range columns {
		columns[i] = strings.Trim(column, " ")
	}
	return that.ResultToModel(columns, value)
}

// 反射 select结果转模型
func (*sqliteSDK) ResultToModel(columns []string, value interface{}) []interface{} {
	var desc []interface{}
	instanceValue := reflect.ValueOf(value).Elem()
	instanceType := reflect.TypeOf(value).Elem()
	for _, column := range columns {
		for i := 0; i < instanceType.NumField(); i++ {
			field := instanceType.Field(i)
			columnName := field.Tag.Get("column")
			fieldType := field.Type.Name()
			if columnName == column {
				value := instanceValue.Field(i).Addr().Interface()
				desc = append(desc, value)
			}
			if fieldType == "Model" {
				modelField := instanceValue.Field(i).Type()
				for j := 0; j < modelField.NumField(); j++ {
					columnName := modelField.Field(j).Tag.Get("column")
					if columnName == column {
						value := instanceValue.Field(i).Field(j).Addr().Interface()
						desc = append(desc, value)
					}
				}
			}
		}
	}

	return desc
}

// 关闭数据库连接
func (that *sqliteSDK) Disconnection() {
	_ = that.instance.Close()
}
