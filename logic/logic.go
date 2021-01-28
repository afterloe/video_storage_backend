package logic

var (
	VideoLogic *videoLogic
	UserLogic *userLogic
)

func init() {
	VideoLogic = &videoLogic{}
	UserLogic = &userLogic{}
}