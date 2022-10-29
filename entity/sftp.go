package entity

import (
	"Beam/utils"
	"github.com/google/uuid"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Sftp struct {
	Url           string
	Username      string
	Password      string
	RetryCount    int
	SleepDuration int
	Location      string
	FileName      string
}

func (f Sftp) Download() error {
	c, err := ftp.Dial(f.Url, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(f.Username, f.Password)
	if err != nil {
		log.Fatal(err)
	}
	c.ChangeDir(utils.GetDirFromUrl(f.Url))

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

func (f Sftp) GetUrl() string {
	return f.Url
}
func (f Sftp) GetDownloadLocation() string {
	return f.Location
}

func (f Sftp) GetRetryCount() int {
	return f.RetryCount
}

func (f Sftp) GetSleepDuration() int {
	return f.SleepDuration
}

func (f Sftp) GetFileName() string {
	return f.FileName
}
