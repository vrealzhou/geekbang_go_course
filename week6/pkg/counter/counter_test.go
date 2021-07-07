package counter

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/vrealzhou/geekbang_go_course/week6/pkg/event"
)

func BenchmarkSingleBucketRecord(b *testing.B) {
	bkt := NewBucket(100 * time.Millisecond)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			evt := event.Event(rand.Int() % event.EventsCount())
			now := time.Now()
			bkt.Record(evt, now, 1)
		}
	})
}

func BenchmarkBucketCounterRecord(b *testing.B) {
	c := NewBucketedCounter(10, 100*time.Millisecond)
	c.StartIfNotStarted(time.Now())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			evt := event.Event(rand.Int() % event.EventsCount())
			c.ReceiveEvent(context.Background(), evt, 1)
		}
	})
}

func TestBucketCounter(t *testing.T) {
	var bucketCnt int64 = 10
	bucketSize := 100 * time.Millisecond
	window := time.Duration(bucketCnt) * bucketSize
	c := NewBucketedCounter(bucketCnt, bucketSize)
	c.StartIfNotStarted(time.Now())
	successCount := 1453
	failureCount := 325
	receiveCh := make(chan bool)
	go func() {
		for i := 0; i < successCount; i++ {
			time.Sleep(2 * time.Millisecond)
			go func() {
				c.ReceiveEvent(context.Background(), event.SUCCESS, 1)
				receiveCh <- true
			}()
		}
	}()
	go func() {
		for i := 0; i < failureCount; i++ {
			time.Sleep(5 * time.Millisecond)
			go func() {
				c.ReceiveEvent(context.Background(), event.FAILURE, 1)
				receiveCh <- false
			}()
		}
	}()
	getCh := time.Tick(window)
	var totalSuccess int64 = 0
	var totalFailure int64 = 0
	countDownSuccess := successCount
	countDownFailure := failureCount
	looping := true
	finished := false
	for looping {
		select {
		case v := <-receiveCh:
			if v {
				countDownSuccess--
			} else {
				countDownFailure--
			}
			if countDownSuccess == 0 && countDownFailure == 0 {
				close(receiveCh)
				finished = true
			}
		case <-getCh:
			successCnt, failureCnt := getResults(t, c, window)
			totalSuccess += successCnt
			totalFailure += failureCnt
			if finished {
				looping = false
			}
		}
	}
	if totalSuccess != int64(successCount) {
		t.Errorf("Incorrect totalSuccess: %d", totalSuccess)
	}
	if totalFailure != int64(failureCount) {
		t.Errorf("Incorrect totalFailure: %d", totalFailure)
	}
}

func getResults(t *testing.T, c *BucketedCounter, window time.Duration) (int64, int64) {
	successCnt, err := c.GetValue(context.Background(), event.SUCCESS, window)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%v, event.SUCCESS: %d", time.Now(), successCnt)
	failureCnt, err := c.GetValue(context.Background(), event.FAILURE, window)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%v, event.FAILURE: %d", time.Now(), failureCnt)
	return successCnt, failureCnt
}
