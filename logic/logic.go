package logic

var (
	VideoLogic *videoLogic
	UserLogic *userLogic
	DictionaryLogic *dictionaryLogic
)

func init() {
	VideoLogic = &videoLogic{}
	UserLogic = &userLogic{}
	DictionaryLogic = &dictionaryLogic{}
}