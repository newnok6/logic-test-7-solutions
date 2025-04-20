package services

type ProcessFileService interface {
	GetFileName() string
	GetFilePath() string
	GetMeatList() (map[string]uint32, error)
}
