package repositories

import "video_storage/repositories/constants"

var UserRecordRepository *userRecordRepository

type userRecordRepository struct {
}

func (*userRecordRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableUser, constants.CreateUserTable, needCreate)
	return nil
}

func (*userRecordRepository) SignIn(email, passwd string) error {
	
	return nil
}

//
//func (*userRecordRepository) InsertOne(job *model.Job) error {
//	args := []interface{}{job.Callback, job.Description, job.IsDel, job.CreateTime}
//	var err error
//	sdk.SQLiteSDK.Execute(func(result sql.Result) {
//		taskId, _ := result.LastInsertId()
//		job.ID = taskId
//		if 0 == job.ID {
//			err = errors.New("插入失败")
//		}
//	}, jobInsert, args...)
//
//	return err
//}
//
//func (*userRecordRepository) FindByID(ID int64, job *model.Job) error {
//	args := []interface{}{ID}
//	var err error
//	sdk.SQLiteSDK.Query(func(rows sql.Rows) {
//		if !rows.Next() {
//			err = errors.New("查询无此结果")
//			return
//		}
//		columns, _ := rows.Columns()
//		err := rows.Scan(sdk.SQLiteSDK.ResultToModel(columns, job)...)
//		if nil != err {
//			logrus.Debug(err)
//			logrus.Debug("解构失败.")
//		}
//	}, jobFindByID, args...)
//
//	return err
//}
//
//func (*jobRecordRepository) ChangeJobStatus(ID int64, status string) error {
//	args := []interface{}{ID, status}
//	var err error
//	sdk.SQLiteSDK.Execute(func(result sql.Result) {
//		resultNum, _ := result.RowsAffected()
//		if 1 != resultNum {
//			err = errors.New("更新失败")
//			return
//		}
//	}, jobChangeStatus, args...)
//	return err
//}
