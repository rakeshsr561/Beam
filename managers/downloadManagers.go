package managers

import (
	"Beam/entity"
	"Beam/utils"
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
