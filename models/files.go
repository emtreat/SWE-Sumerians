package models

type File struct {
	FileName string `json:"file_name"`
	FileSize int32  `json:"file_size"`
}

type Files struct {
	Files []File
}
