package constants

const (
	TableVideo       = "video_describe"
	CreateVideoTable = `
CREATE TABLE IF NOT EXISTS "` + TableVideo + `" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"metadata_id" INTEGER
	"width"	INTEGER,
	"height"	INTEGER,
	"duration"	TEXT,
	"codec_name"	TEXT,
	"display_aspect_ratio" TEXT,
	"codec_long_name"	TEXT
	"create_time"	TEXT,
	"modify_time"	TEXT,
	"is_del"	BLOB
)`
	InsertVideoDescribe = "INSERT INTO " + TableVideo + " (metadata_id, width, height, duration, codec_name, display_aspect_ratio, codec_long_name, create_time, modify_time, is_del) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	UpdateDemandVideo = "UPDATE " + TableVideo + " SET name = ? , size = ? , width = ? , height = ? , duration = ? , path = ? , describe = ? , title = ? , ffmpeg_json = ? , modify_time = ? , is_del = ? WHERE id = ?"
	VideoIsIncluded   = "SELECT id, name, size, width, height, duration, path, describe, title, ffmpeg_json, create_time, modify_time, is_del FROM " + TableVideo + " WHERE path = ?"
	VideoFindByID     = "SELECT id, name, size, width, height, duration, path, describe, title, ffmpeg_json, create_time, modify_time, is_del FROM " + TableVideo + " WHERE id = ?"
	VideoGetList      = "SELECT id, name, size, width, height, duration, path, describe, title, create_time, modify_time FROM " + TableVideo + " WHERE is_del = ? ORDER BY modify_time DESC LIMIT ? OFFSET ?"
	VideoTotalCount   = "SELECT count(1) FROM " + TableVideo + " WHERE is_del = ?"
)
