package constants

const (
	TableUser   = "member"
	CreateUserTable = `
CREATE TABLE IF NOT EXISTS "member" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"mail"	TEXT NOT NULL,
	"pwd"	TEXT NOT NULL,
	"nickname"	TEXT,
	"avatar"	TEXT,
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`

	)
