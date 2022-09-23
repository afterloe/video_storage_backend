package logic

var (
	VideoLogic        *videoLogic
	UserLogic         *userLogic
	DictionaryLogic   *dictionaryLogic
	FileMetadataLogic *fileMetadataLogic
	ObjectLogic       *objectLogic
)

func init() {
	VideoLogic = &videoLogic{}
	UserLogic = &userLogic{}
	DictionaryLogic = &dictionaryLogic{}
	FileMetadataLogic = &fileMetadataLogic{}
	ObjectLogic = &objectLogic{}
}
