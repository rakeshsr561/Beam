package entity

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

type Http struct {
	Url           string
	RetryCount    int
	SleepDuration int
	Location      string
	FileName      string
}

func (h Http) Download() error {
	resp, err := http.Get(h.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	temp_file_name := uuid.NewString()

	f, err := os.Create(h.Location + "/" + temp_file_name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		os.Remove(h.Location + "/" + temp_file_name)
	} else {
		if _, err1 := os.Stat(h.Location + "/" + h.FileName); err1 == nil {
			os.Remove(h.Location + "/" + h.FileName)
		}
		os.Rename(h.Location+"/"+temp_file_name, h.Location+"/"+h.FileName)
	}

	return err
}

func (f Http) GetUrl() string {
	return f.Url
}
func (f Http) GetDownloadLocation() string {
	return f.Location
}

func (f Http) GetRetryCount() int {
	return f.RetryCount
}

func (f Http) GetSleepDuration() int {
	return f.SleepDuration
}

func (f Http) GetFileName() string {
	return f.FileName
}
