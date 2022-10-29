package entity

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Retry struct {
	Attempts int
	Sleep    int
	F        func() error
}

func (r Retry) RetryF(wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	for i := 0; i < r.Attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(time.Duration(r.Sleep))
			r.Sleep *= 2
		}
		err = r.F()
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", r.Attempts, err)
}
