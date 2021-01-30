package constants

const (
	TableDictionaryGroup       = "dictionary_group"
	CreateDictionaryGroupTable = `
CREATE TABLE IF NOT EXISTS "` + TableDictionaryGroup + `" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name"	TEXT,
	"group_type" TEXT,
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`
	FindDictionaryGroupByName = "SELECT id, name, group_type FROM " + TableDictionaryGroup + " WHERE name = ?"
	CreateDictionaryGroup     = "INSERT INTO " + TableDictionaryGroup + " (name, group_type, create_time, modify_time, is_del) VALUES (?, ?, ?, ?, ?)"
	FindDictionaryGroupByID = "SELECT id, name, group_type, create_time, modify_time, is_del FROM " + TableDictionaryGroup + " WHERE id = ?"
	FindAllDictionaryGroup = "SELECT id, name, group_type, create_time, modify_time FROM " + TableDictionaryGroup + " WHERE is_del = ?"
	TableDictionary = "dictionary"
	CreateDictionaryTable = `
CREATE TABLE IF NOT EXISTS "` + TableDictionary + `" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name"	TEXT,
	"data"	TEXT,
	"group_id" INTEGER,
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`
	CreateDictionary = "INSERT INTO " + TableDictionary + " (name, data, group_id, create_time, modify_time, is_del) VALUES (?, ?, ?, ?, ?, ?)"
	FindAllDictionary = "SELECT id, name, data, group_id, create_time, modify_time, is_del FROM " + TableDictionary + " WHERE group_id = ?"
)
