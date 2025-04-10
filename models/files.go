package models

type File struct {
	FileName string `json:"filename"`
	FileSize int32  `json:"filesize"`
}

type Files struct {
	Files []File
}
