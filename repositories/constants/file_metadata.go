package constants

const (
	TableFileMetadata       = "file_metadata"
	CreateFileMetaDataTable = `
	CREATE TABLE IF NOT EXISTS "` + TableFileMetadata + `" (
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"hex_code"	TEXT NOT NULL,
		"filename"	TEXT NOT NULL,
		"file_size"	INTEGER,
		"file_type"	TEXT,
		"fullpath"	TEXT,
		"create_time"	TEXT,
		"is_del"	BLOB,
		"is_link"	BLOB
)`
	FindAllFileMetaData = `SELECT id, filename, file_type, file_size, create_time, fullpath, is_link FROM ` + TableFileMetadata + ` WHERE is_del = true LIMIT ? OFFSET ?`
)
