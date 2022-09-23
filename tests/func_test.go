package tests

import (
	"strings"
	"testing"
	"video_storage/tools"
)

func TestExecuteCMD(t *testing.T) {
	cmdStr := "ffprobe -v quiet -print_format json -show_format -show_streams %s"
	receive := tools.Execute(cmdStr, `/mnt/wehouse_3`)
	// if receive.HasError() != nil {
	// 	// t.Errorf("%s\n", receive.HasError().Error())
	// }
	t.Logf("%s\n", receive.ToString())
}

func TestSQLToModel(t *testing.T) {
	TableDictionaryGroup := "User"
	// s := "SELECT id, name, group_type, create_time, modify_time, is_del FROM " + TableDictionaryGroup + " WHERE id = ?"
	s := "SELECT id, name, group_type, create_time, modify_time FROM " + TableDictionaryGroup + " WHERE is_del = ?"
	i := strings.Index(s, "FROM")
	cs := strings.Split(s[6:i], ",")
	for _, c := range cs {
		c = strings.Trim(c, " ")
		t.Log(c)
	}
}
