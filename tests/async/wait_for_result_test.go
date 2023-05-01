package async

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/strivesolutions/go-strive-utils/pkg/async"
)

type asyncResult struct {
	intervals int
	count     int
}

func (r *asyncResult) Done() bool {
	return r.count > r.intervals
}

func TestWaitReturnsResultWhenDone(t *testing.T) {
	result := &asyncResult{intervals: 1}

	f := func() (*asyncResult, error) {
		time.Sleep(time.Second * 1)
		result.count++
		return result, nil
	}

	result, err := async.WaitForResult(5, f)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestWaitTimesOut(t *testing.T) {
	result := &asyncResult{intervals: 10}

	f := func() (*asyncResult, error) {
		time.Sleep(time.Second * 1)
		return result, nil
	}

	result, err := async.WaitForResult(1, f)

	assert.NotNil(t, err)
	assert.Nil(t, result)
}
