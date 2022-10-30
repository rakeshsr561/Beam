package entity

import (
	"errors"
	"sync"
	"testing"
)

type RetryRest struct {
	value int
}

func (r *RetryRest) f1() error {
	r.value -= 1
	return errors.New("test error")
}

func TestValidUrl(t *testing.T) {
	var wg sync.WaitGroup
	testS := RetryRest{5}
	retry := Retry{3, 1, testS.f1}
	wg.Add(1)
	go retry.RetryF(&wg)
	wg.Wait()
	actual := testS.value

	expected := 2
	if actual != expected {
		t.Errorf("Expected (%d) is not same as"+
			" actual (%d)", expected, actual)
	}
}
