package repositories

import (
	"database/sql"
	"errors"
	"video_storage/model"
	"video_storage/repositories/constants"
	"video_storage/sdk"
	"video_storage/tools"
)

var UserRecordRepository *userRecordRepository

type userRecordRepository struct {
}

func (*userRecordRepository) repositoryTable(needCreate bool) error {
	tableRepository(constants.TableUser, constants.CreateUserTable, needCreate)
	return nil
}

func (that *userRecordRepository) SignUp(email, passwd string) (*model.User, error) {
	user := that.FindUserByEmail(email)
	var err error
	if "" != user.Mail {
		err = errors.New("该账号已被注册")
	} else {
		user.Mail = email
		user.Passwd = tools.MD5(email + passwd)
		err = that.InsertOne(user)
	}
	return user, err
}

func (*userRecordRepository) InsertOne(user *model.User) error {
	//mail, pwd, nickname, avatar, create_time, modify_time, is_del
	createTime := tools.GetTime()
	args := []interface{}{user.Mail, user.Passwd, user.Mail, "", createTime, createTime, false}
	var err error
	sdk.SQLiteSDK.Execute(func(result sql.Result) {
		id, _ := result.LastInsertId()
		user.ID = id
		if 0 == user.ID {
			err = errors.New("插入失败")
		}
	}, constants.InsertOne, args...)

	return err
}

func (*userRecordRepository) FindUserByEmail(email string) *model.User {
	args := []interface{}{email}
	var user = model.User{}
	sdk.SQLiteSDK.QueryOne(func(row sql.Row) {
		row.Scan(&user.ID, &user.Mail, &user.IsDel)
	}, constants.FindUserByEmail, args...)
	return &user
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
