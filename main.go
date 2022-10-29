package main

import (
	"Beam/entity"
	"Beam/managers"
	"Beam/utils"
)

func main() {
	urls := utils.ReadFileByLine("site_detail")
	config := entity.LoadConfiguration("config.json")
	managers.DownLoadAll(urls, config)

}
