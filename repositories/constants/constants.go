package constants

const (
	RepositoryTable = "SELECT COUNT(1) FROM sqlite_master WHERE type = ? and name = ?"
)