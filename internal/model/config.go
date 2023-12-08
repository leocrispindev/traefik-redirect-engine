package model

type Config struct {
	IsRedirectEnable bool   `json:"isRedirectEnable,omitempty"`
	Source           string `json:"source,omitempty"`
	FilePath         string `json:"filePath,omitempty"`
}
