package utils

import (
	"net/url"
	"strings"
)

func ValidateUrl(urlString string) bool {

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		
		return false
	}

	return true
}

func GetDirFromUrl(urlInput string) string {
	urlString, _ := url.Parse(urlInput)
	path := strings.Split(urlString.Path, "/")
	path = path[:len(path)-1]
	pathSting := strings.Join(path, "/")
	return pathSting
}
