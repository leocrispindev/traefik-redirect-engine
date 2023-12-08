package utils

import (
	"fmt"
	"net/http"
	"os"
)

func GetFullUrl(r *http.Request) string {
	result := fmt.Sprintf("%s://%s%s", getProtocol(r), r.Host, r.RequestURI)

	return result
}

func getProtocol(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func Readfile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
