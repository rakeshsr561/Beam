package entity

type Downloader interface {
	Download() error
	GetUrl() string
	GetDownloadLocation() string
	GetRetryCount() int
	GetSleepDuration() int
	GetFileName() string
}
