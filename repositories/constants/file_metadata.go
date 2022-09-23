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
	FindAllFileMetaData            = "SELECT id, filename, file_type, file_size, create_time, fullpath, is_link FROM " + TableFileMetadata + " WHERE is_del = false LIMIT ? OFFSET ?"
	FileMetadataCount              = "SELECT count(1) FROM " + TableFileMetadata + " WHERE is_del = ?"
	FindFileMetadataByID           = "SELECT id, hex_code, filename, file_type, file_size, create_time, fullpath, is_link, is_del FROM " + TableFileMetadata + " WHERE id = ?"
	FindFileMetadataByKeyword      = "SELECT id, filename, file_type, file_size, create_time, fullpath, is_link FROM " + TableFileMetadata + " WHERE fullpath LIKE ? AND is_del = ? LIMIT ? OFFSET ?"
	FindFileMetadataByKeywordCount = "SELECT count(1) FROM " + TableFileMetadata + " WHERE fullpath LIKE ? AND is_del = ?"
)
