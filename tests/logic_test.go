package tests

import (
	"testing"
	"video_storage/logic"
	"video_storage/repositories"
	"video_storage/sdk"
)

var sqlInitFlag = false

func initSQL() {
	if false == sqlInitFlag {
		sdk.Init("/home/afterloe/Desktop/video_storage.db")
		repositories.Init()
		sqlInitFlag = true
	}
}

func TestSignIn(t *testing.T) {
	initSQL()
	const (
		email = "605728727@qq.com"
		passwd = "11111111"
	)
	user, err := logic.UserLogic.SignIn(email, passwd)
	if nil != err {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestSignUp(t *testing.T) {
	initSQL()
	const (
		email = "605728727@qq.com"
		passwd = "11111111"
	)
	user, err := logic.UserLogic.SignUp(email, passwd)
	if nil != err {
		t.Error(err)
		return
	}
	t.Log("INSERT SUCCESS")
	t.Log(user)
}