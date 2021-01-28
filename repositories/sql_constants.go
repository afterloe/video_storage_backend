package repositories

const (


	// task
	taskRecordTableName   = "agent_task_record"
	taskRecordCreateTable = `
		CREATE TABLE "agent_task_record"
		(
			"id"          INTEGER PRIMARY KEY AUTOINCREMENT,
			"job_id"      INTEGER,
			"commands"    TEXT,
			"status"      TEXT,
			"result"      TEXT,
			"is_del"      BOOLEAN,
			"create_time" TEXT,
			"modify_time" TEXT
		)`
	taskRecordInsert         = "INSERT INTO agent_task_record (job_id, commands, status, result, is_del, create_time) VALUES (?, ?, ?, ?, ?, ?)"
	taskRecordFindByID       = "SELECT id, job_id, commands, status, result, is_del, create_time, modify_time FROM agent_task_record WHERE id = ?"
	taskRecordQueryTaskByMsg = "SELECT id, job_id, commands, status, result, is_del, create_time, modify_time FROM agent_task_record WHERE job_id = ? AND commands = ?"
	taskRecordCompleteTask   = "UPDATE agent_task_record SET result = ?, status = ?, modify_time = ? WHERE id = ?"

	// job

	jobInsert   = "INSERT INTO agent_job (callback, description, is_del, create_time) VALUES (?, ?, ?, ?)"
	jobFindByID = "SELECT id, callback, description, is_del, create_time, modify_time FROM agent_job WHERE id = ?"
	jobChangeStatus = "UPDATE agent_job SET status = ? WHERE id = ?"
)