BEGIN TRANSACTION;
DROP TABLE IF EXISTS `video_target_bind`;
CREATE TABLE IF NOT EXISTS `video_target_bind` (
	`video_id`	INTEGER,
	`target_id`	INTEGER,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB
);
DROP TABLE IF EXISTS `video_target`;
CREATE TABLE IF NOT EXISTS `video_target` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`name`	TEXT,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB
);
DROP TABLE IF EXISTS `video_screenshot`;
CREATE TABLE IF NOT EXISTS `video_screenshot` (
	`name`	TEXT,
	`size`	INTEGER,
	`on_demand_video`	TEXT,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB,
	`group_id`	INTEGER
);
DROP TABLE IF EXISTS `role_permit_bind`;
CREATE TABLE IF NOT EXISTS `role_permit_bind` (
	`role_id`	INTEGER,
	`permit_id`	INTEGER,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB
);
DROP TABLE IF EXISTS `power_role`;
CREATE TABLE IF NOT EXISTS `power_role` (
	`id`	INTEGER,
	`name`	TEXT,
	`describe`	TEXT,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	INTEGER
);
DROP TABLE IF EXISTS `power_permit`;
CREATE TABLE IF NOT EXISTS `power_permit` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`name`	TEXT,
	`url`	TEXT,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB,
	`method`	TEXT
);
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE IF NOT EXISTS `operation_log` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`user_id`	INTEGER,
	`create_time`	TEXT,
	`permit_id`	INTEGER,
	`content`	TEXT
);
DROP TABLE IF EXISTS `on_demand_video`;
CREATE TABLE IF NOT EXISTS `on_demand_video` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`name`	TEXT,
	`size`	INTEGER,
	`width`	INTEGER,
	`height`	INTEGER,
	`duration`	INTEGER,
	`describe`	TEXT,
	`title`	TEXT,
	`ffmpeg_json`	TEXT,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB
);
DROP TABLE IF EXISTS `member_role_bind`;
CREATE TABLE IF NOT EXISTS `member_role_bind` (
	`member_id`	INTEGER,
	`role_id`	INTEGER,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB
);
DROP TABLE IF EXISTS `member_real_name`;
CREATE TABLE IF NOT EXISTS `member_real_name` (
	`member_id`	INTEGER NOT NULL,
	`real_name`	TEXT,
	`hometown`	TEXT,
	`birthday`	TEXT,
	`real_id`	TEXT,
	`sex`	BLOB,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB
);
DROP TABLE IF EXISTS `member`;
CREATE TABLE IF NOT EXISTS `member` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`mail`	TEXT NOT NULL,
	`pwd`	TEXT NOT NULL,
	`nickname`	TEXT,
	`avatar`	TEXT,
	`create_time`	TEXT,
	`modify_time`	TEXT,
	`is_del`	BLOB
);
COMMIT;
