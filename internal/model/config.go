package model

type Config struct {
	IsRedirectEnable bool   `json:"isRedirectEnable"`
	Source           string `json:"source"`
	FilePath         string `json:"filePath"`
}
