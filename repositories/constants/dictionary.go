package constants

const (
	TableDictionary       = "dictionary_group"
	CreateDictionaryTable = `
CREATE TABLE IF NOT EXISTS "` + TableDictionary + `" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name"	TEXT,
	"group_type" TEXT,
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`
	FindDictionaryGroupByName = "SELECT id, name, group_type FROM " + TableDictionary + " WHERE name = ?"
	CreateDictionaryGroup     = "INSERT INTO " + TableDictionary + " (name, group_type, create_time, modify_time, is_del) VALUES (?, ?, ?, ?, ?)"
)
