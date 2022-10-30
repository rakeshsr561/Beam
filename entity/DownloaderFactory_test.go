package entity

import (
	"Beam/utils"
	"fmt"
	"strings"
	"testing"
)

func TestBuildDownloadersFromUrlList(t *testing.T) {
	urls := utils.ReadFileByLine("../site_detail_test")
	config := LoadConfiguration("../config_test.json")
	downloaders := BuildDownloadersFromUrlList(urls, config)
	actual := 4

	expected := len(downloaders)
	fmt.Println(len(downloaders))
	if 4 != expected {
		t.Errorf("Expected (%d) is not same as"+
			" actual (%d)", expected, actual)
	}
}

func TestBuildDownloadersFromUrlListData(t *testing.T) {
	urls := utils.ReadFileByLine("../site_detail_test")
	config := LoadConfiguration("../config_test.json")
	downloaders := BuildDownloadersFromUrlList(urls, config)
	ftp := downloaders[0]
	http := downloaders[1]
	https := downloaders[2]
	sftp := downloaders[3]

	actual := ftp.GetUrl() == "ftp://abc:21/folder/one.end_one"
	actual = actual && http.GetUrl() == "http://abc.com/folder/two.end_two"
	actual = actual && https.GetUrl() == "https://abc.com/folder/three.end_three"
	actual = actual && sftp.GetUrl() == "sftp://abc:21/folder/four.end_four"

	actual = actual && ftp.GetDownloadLocation() == "ftp_folder"
	actual = actual && http.GetDownloadLocation() == "http_folder"
	actual = actual && https.GetDownloadLocation() == "https_folder"
	actual = actual && sftp.GetDownloadLocation() == "sftp_folder"

	actual = actual && ftp.GetRetryCount() == 4
	actual = actual && http.GetRetryCount() == 2
	actual = actual && https.GetRetryCount() == 3
	actual = actual && sftp.GetRetryCount() == 5

	actual = actual && ftp.GetSleepDuration() == 40
	actual = actual && http.GetSleepDuration() == 20
	actual = actual && https.GetSleepDuration() == 30
	actual = actual && sftp.GetSleepDuration() == 50

	actual = actual && strings.HasPrefix(ftp.GetFileName(), "one_")
	actual = actual && strings.HasPrefix(http.GetFileName(), "two_")
	actual = actual && strings.HasPrefix(https.GetFileName(), "three_")
	actual = actual && strings.HasPrefix(sftp.GetFileName(), "four_")

	actual = actual && strings.HasSuffix(ftp.GetFileName(), ".end_one")
	actual = actual && strings.HasSuffix(http.GetFileName(), ".end_two")
	actual = actual && strings.HasSuffix(https.GetFileName(), ".end_three")
	actual = actual && strings.HasSuffix(sftp.GetFileName(), ".end_four")

	expected := true

	if actual != expected {
		t.Errorf("Expected (%t) is not same as"+
			" actual (%t)", expected, actual)
	}

}
