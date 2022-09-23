package logic

var (
	VideoLogic        *videoLogic
	UserLogic         *userLogic
	DictionaryLogic   *dictionaryLogic
	FileMetadataLogic *fileMetadataLogic
)

func init() {
	VideoLogic = &videoLogic{}
	UserLogic = &userLogic{}
	DictionaryLogic = &dictionaryLogic{}
	FileMetadataLogic = &fileMetadataLogic{}
}
