package logic

import (
	"errors"
	"strings"
	"video_storage/model"
	"video_storage/repositories"
	"video_storage/tools"
)

type userLogic struct {
}

func (*userLogic) SignIn(email, passwd string) (*model.User, error) {
	user := repositories.UserRecordRepository.FindUserByPwd(email, tools.MD5(email+passwd))
	var err error
	if 0 == user.ID {
		err = errors.New("账号密码错误")
	}
	if true == user.IsDel {
		err = errors.New("该账号已被冻结")
	}
	return user, err
}

func (*userLogic) SignUp(email, passwd string) (*model.User, error) {
	user := repositories.UserRecordRepository.FindUserByEmail(email)
	var err error
	if "" != user.Mail {
		err = errors.New("该账号已被注册")
	} else {
		user.Mail = email
		user.Passwd = tools.MD5(email + passwd)
		user.Nickname = strings.Split(email, "@")[0]
		user.Avatar = "anyOne.png"
		err = repositories.UserRecordRepository.InsertOne(user)
	}
	return user, err
}

func (*userLogic) CheckLoginStatus(uid int64) (string, error) {
	group := repositories.MemoryStorageRepository.GetAllTypeValue("user")
	for token, v := range group {
		if v.(*model.User).ID == uid {
			return token, errors.New("当前账号已登陆")
		}
	}
	return "", nil
}

func (*userLogic) Cancellation(token string)  {
	repositories.MemoryStorageRepository.Del("user", token)
}