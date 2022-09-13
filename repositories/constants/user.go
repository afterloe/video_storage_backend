package constants

const (
	TableUser       = "member"
	CreateUserTable = `
CREATE TABLE IF NOT EXISTS "` + TableUser + `" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"mail"	TEXT NOT NULL,
	"pwd"	TEXT NOT NULL,
	"nickname"	TEXT,
	"avatar"	TEXT,
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`
	FindUserByEmail = `SELECT id, mail, is_del FROM ` + TableUser + ` WHERE mail = ?`
	InsertOne       = "INSERT INTO " + TableUser + " (mail, pwd, nickname, avatar, create_time, modify_time, is_del) VALUES (?, ?, ?, ?, ?, ?, ?)"
	SignIn          = "SELECT id, mail, nickname, avatar, is_del FROM " + TableUser + " WHERE mail = ? AND pwd = ?"
	FindByID        = "SELECT id, mail, pwd, nickname, avatar, create_time, modify_time, is_del FROM " + TableUser + " WHERE id = ?"
)
