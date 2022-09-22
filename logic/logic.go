package logic

var (
	VideoLogic        *videoLogic
	UserLogic         *userLogic
	DictionaryLogic   *dictionaryLogic
	FileMeatdataLogic *fileMeatdataLogic
)

func init() {
	VideoLogic = &videoLogic{}
	UserLogic = &userLogic{}
	DictionaryLogic = &dictionaryLogic{}
	FileMeatdataLogic = &fileMeatdataLogic{}
}
