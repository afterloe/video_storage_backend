package constants

const (
	TableDictionary = "dictionary_group"
	CreateDictionaryTable = `
CREATE TABLE IF NOT EXISTS "` + TableDictionary + `" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name"	TEXT,
	"group_type" TEXT,
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`

)