package constants

const (
	TableVideo = "on_demand_video"
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
	UpdateDemandVideo = "UPDATE " + TableVideo + " SET name = ? AND size = ? AND width = ? AND height = ? AND duration = ? AND path = ? AND describe = ? AND title = ? AND ffmpeg_json = ? AND modify_time = ? AND is_del = ? WHERE id = ?"
) 
