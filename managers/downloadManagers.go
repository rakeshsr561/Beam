package managers

import (
	"Beam/entity"
	"Beam/utils"
	"os"
	"strings"
	"sync"
)

func DownLoadAll(urls []string, config entity.Config) {
	urls = removeDuplicateUrls(urls)
	downloaders := entity.BuildDownloadersFromUrlList(urls, config)
	var wg sync.WaitGroup
	for _, downloader := range downloaders {
		err := utils.CreateFolderIfNotExists(downloader.GetDownloadLocation())
		if err != nil {

		}
		retry := entity.Retry{downloader.GetRetryCount(), downloader.GetSleepDuration(), downloader.Download}
		wg.Add(1)
		go retry.RetryF(&wg)
	}
	wg.Wait()

}

func removeDuplicateUrls(urls []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range urls {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func replace_old_copy_by_new_copy(downloader entity.Downloader) {
	copy_name := downloader.GetDownloadLocation() + "/" + downloader.GetFileName()
	name := strings.Split(downloader.GetFileName(), "__")
	fileName := downloader.GetFileName()
	fileType := fileName[strings.LastIndex(fileName, "."):]
	real_name := downloader.GetDownloadLocation() + "/" + name[0] + fileType
	e1 := os.Remove(real_name)
	if e1 != nil {
		return
	}
	e2 := os.Rename(copy_name, real_name)

	if e2 != nil {
		return
	}
}
