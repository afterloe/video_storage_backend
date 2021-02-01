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
) 
