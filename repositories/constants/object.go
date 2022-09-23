package constants

const (
	TableObject       = "object"
	CreateObjectTable = `
	CREATE TABLE IF NOT EXISTS "` + TableObject + `" (
		"id"			INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"virtual_path"	TEXT NOT NULL,
		"metadata_id"	INTEGER NOT NULL,
		"create_time"	TEXT,
		"is_del"		BLOB,
		"modify_time"	TEXT
)`
	FindObjectByMetadataID = "SELECT id, virtual_path, metadata_id, create_time, is_del, modify_time FROM " + TableObject + " WHERE metadata_id = ?"
	InsertObject           = "INSERT INTO " + TableObject + " (virtual_path, metadata_id, create_time, is_del, modify_time) VALUES (?, ?, ?, ?, ?)"
)
