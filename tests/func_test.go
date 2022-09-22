package tests

import (
	"strings"
	"testing"
)

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
