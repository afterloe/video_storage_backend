package constants

const (
	TableVideo       = "on_demand_video"
	CreateVideoTable = `
CREATE TABLE IF NOT EXISTS "` + TableVideo + `" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name"	TEXT,
	"size"	INTEGER,
	"width"	INTEGER,
	"height"	INTEGER,
	"duration"	INTEGER,
    "path" TEXT,
	"describe"	TEXT,
	"title"	TEXT,
	"ffmpeg_json"	TEXT,
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`
	InsertDemandVideo = "INSERT INTO " + TableVideo + " (name, size, width, height, duration, path, describe, title, ffmpeg_json, create_time, modify_time, is_del) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	UpdateDemandVideo = "UPDATE " + TableVideo + " SET name = ? , size = ? , width = ? , height = ? , duration = ? , path = ? , describe = ? , title = ? , ffmpeg_json = ? , modify_time = ? , is_del = ? WHERE id = ?"
	VideoIsIncluded   = "SELECT id, name, size, width, height, duration, path, describe, title, ffmpeg_json, create_time, modify_time, is_del FROM " + TableVideo + " WHERE path = ?"
	VideoFindByID     = "SELECT id, name, size, width, height, duration, path, describe, title, ffmpeg_json, create_time, modify_time, is_del FROM " + TableVideo + " WHERE id = ?"
	VideoGetList      = "SELECT id, name, size, width, height, duration, path, describe, title, create_time, modify_time FROM " + TableVideo + " WHERE is_del = ? ORDER BY modify_time DESC LIMIT ? OFFSET ?"
	VideoTotalCount   = "SELECT count(1) FROM " + TableVideo + " WHERE is_del = ?"
)
