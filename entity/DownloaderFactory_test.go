package entity

import (
	"Beam/utils"
	"fmt"
	"testing"
)

func TestBuildDownloadersFromUrlList(t *testing.T) {
	urls := utils.ReadFileByLine("../site_detail")
	config := LoadConfiguration("config.json")
	downloaders := BuildDownloadersFromUrlList(urls, config)
	actual := 4

	expected := len(downloaders)
	fmt.Println(len(downloaders))
	if 4 != expected {
		t.Errorf("Expected (%d) is not same as"+
			" actual (%d)", expected, actual)
	}
}
