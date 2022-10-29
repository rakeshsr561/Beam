package entity

import (
	"github.com/google/uuid"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

type Ftp struct {
	Url           string
	RetryCount    int
	SleepDuration int
	Location      string
	FileName      string
}

func (f Ftp) Download() error {
	urlString, _ := url.Parse(f.Url)
	c, err := ftp.Dial(urlString.Host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("", "")
	if err != nil {
		log.Fatal(err)
	}

	c.ChangeDir(urlString.Path)

	segments := strings.Split(f.Url, "/")
	serverFileName := segments[len(segments)-1]

	res, err := c.Retr(serverFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	temp_file_name := uuid.NewString()

	outFile, err := os.Create(f.Location + "/" + temp_file_name)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, res)

	if err != nil {
		os.Remove(f.Location + "/" + temp_file_name)
	} else {
		if _, err1 := os.Stat(f.Location + "/" + f.FileName); err1 == nil {
			os.Remove(f.Location + "/" + f.FileName)
		}
		os.Rename(f.Location+"/"+temp_file_name, f.Location+"/"+f.FileName)
	}

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (f Ftp) GetUrl() string {
	return f.Url
}

func (f Ftp) GetDownloadLocation() string {
	return f.Location
}

func (f Ftp) GetRetryCount() int {
	return f.RetryCount
}

func (f Ftp) GetSleepDuration() int {
	return f.SleepDuration
}

func (f Ftp) GetFileName() string {
	return f.FileName
}
