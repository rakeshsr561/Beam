package entity

import (
	"Beam/utils"
	"fmt"
	"hash/fnv"
	"strings"
)

func get_hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))

	return fmt.Sprint(h.Sum32())
}

func create_file_name_from_url(url string) string {
	lastInd := strings.LastIndex(url, "/")
	fileName := url[lastInd+1:]
	fmt.Println(fileName)
	lastInd = strings.LastIndex(fileName, ".")
	if lastInd == -1 {
		return fileName + "_" + get_hash(url)
	} else {
		name := fileName[:lastInd]
		fileType := fileName[lastInd:]
		return name + "_" + get_hash(url) + fileType
	}

}

func BuildDownloadersFromUrlList(urls []string, config Config) []Downloader {
	var downloaders []Downloader
	for _, url := range urls {
		urlArray := strings.Split(url, "##")
		var userName string
		var password string
		if len(urlArray) == 1 {
			userName = ""
			password = ""

		} else {
			userName = urlArray[1]
			password = urlArray[2]
		}
		if !utils.ValidateUrl(urlArray[0]) {
			continue
		}
		fileName := create_file_name_from_url(strings.Split(url, "##")[0])
		if strings.HasPrefix(url, "https") {
			downloaders = append(downloaders, Https{url, config.Https.Retry, config.Https.SleepTIme, config.Https.DestinationFolder, fileName})
		} else if strings.HasPrefix(url, "http") {
			downloaders = append(downloaders, Http{url, config.Http.Retry, config.Http.SleepTIme, config.Http.DestinationFolder, fileName})
		} else if strings.HasPrefix(url, "ftp") {
			downloaders = append(downloaders, Ftp{urlArray[0], userName, password, config.Ftp.Retry, config.Ftp.SleepTIme, config.Ftp.DestinationFolder, fileName})
		} else if strings.HasPrefix(url, "sftp") {
			fmt.Println(urlArray[0])
			downloaders = append(downloaders, Sftp{urlArray[0], userName, password, config.Sftp.Retry, config.Sftp.SleepTIme, config.Sftp.DestinationFolder, fileName})
		}
	}
	return downloaders

}
