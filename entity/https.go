package entity

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

type Https struct {
	Url           string
	RetryCount    int
	SleepDuration int
	Location      string
	FileName      string
}

func (hs Https) Download() error {
	fmt.Println("hoi")
	resp, err := http.Get(hs.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	temp_file_name := uuid.NewString()

	f, err := os.Create(hs.Location + "/" + temp_file_name)
	fmt.Println(hs.Location + "/" + temp_file_name)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	f.Close()
	//os.Rename(hs.Location+"/"+temp_file_name, hs.Location+"/"+hs.FileName)
	fmt.Println("hello")
	fmt.Println(hs.Location + "/" + hs.FileName)
	fmt.Println(hs.FileName)
	if err != nil {
		//os.Remove(hs.Location + "/" + temp_file_name)
	} else {
		if _, err1 := os.Stat(hs.Location + "/" + hs.FileName); err1 == nil {
			os.Remove(hs.Location + "/" + hs.FileName)
		}
		os.Rename(hs.Location+"/"+temp_file_name, hs.Location+"/"+hs.FileName)
	}
	return err
}

func (f Https) GetUrl() string {
	return f.Url
}
func (f Https) GetDownloadLocation() string {
	return f.Location
}

func (f Https) GetRetryCount() int {
	return f.RetryCount
}

func (f Https) GetSleepDuration() int {
	return f.SleepDuration
}

func (f Https) GetFileName() string {
	return f.FileName
}
