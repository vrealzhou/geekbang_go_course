// 计数器。给定bucket数量和每个bucket的时段长度。创建一个滑动窗口计数器
//
package counter

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/vrealzhou/geekbang_go_course/week6/pkg/event"
)

// 计数器
type BucketedCounter struct {
	buckets     []*Bucket                              // 若干个单元的计数器
	bucketIndex func(startTime int64, t time.Time) int // 用来计算指定时间点落于哪个bucket
	bucketSize  time.Duration                          // 每个单元的时长
	startTime   *int64                                 // int64(len(bucket))结果，暂存下来减少计算时间
}

func NewBucketedCounter(numBuckets int64, bucketSize time.Duration) *BucketedCounter {
	totalBuckets := numBuckets * 2 // 存储双倍buckets防止覆盖
	buckets := make([]*Bucket, totalBuckets)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = NewBucket(bucketSize)
	}
	var start int64 = 0
	return &BucketedCounter{
		buckets:    buckets,
		bucketSize: bucketSize,
		bucketIndex: func(startTime int64, t time.Time) int {
			// 给定时间到起始时间的时间差 / 每段时间差 得出总共多少段
			// 模上总bucket数，得出给定时间在哪个bucket中
			// index在bucket数组中不断循环
			return int((t.UnixNano()-startTime)/bucketSize.Nanoseconds()) % len(buckets)
		},
		startTime: &start,
	}
}

// 启动计数器
func (c *BucketedCounter) StartIfNotStarted(startTime time.Time) {
	start := atomic.LoadInt64(c.startTime)
	if start != 0 {
		return
	}
	atomic.CompareAndSwapInt64(c.startTime, 0, startTime.UnixNano())
}

// 接收一个事件，实现了EventReceiver接口。
// ctx: 上下文
// event: 事件
// value: 事件值
func (c *BucketedCounter) ReceiveEvent(ctx context.Context, event event.Event, value int64) {
	now := time.Now()
	c.StartIfNotStarted(now)
	c.buckets[c.bucketIndex(atomic.LoadInt64(c.startTime), time.Now())].Record(event, now, value)
}

// 取得当前时间往前推指定窗口时段(window)的特定事件统计值
// ctx: 上下文
// window: 窗口时段
// evtType: 事件类型
func (c *BucketedCounter) GetValue(ctx context.Context, evtType event.Event, window time.Duration) (int64, error) {
	now := time.Now()
	c.StartIfNotStarted(now)
	// 计算时间窗口是否过大
	if window.Milliseconds() > c.bucketSize.Milliseconds()*int64(len(c.buckets)/2) {
		return 0, fmt.Errorf("window %d ms is longer than max allowed window", window.Milliseconds())
	}
	windowBegin := now.Add(0 - window)
	startTime := atomic.LoadInt64(c.startTime)
	begin := c.bucketIndex(startTime, windowBegin)
	if begin < 0 {
		begin = 0
	}
	end := c.bucketIndex(startTime, now)
	if end < begin {
		end += len(c.buckets)
	}
	var result int64 = 0
	for i := begin; i < end; i++ {
		result += c.buckets[i%len(c.buckets)].Get(windowBegin, evtType)
	}
	return result, nil
}

// 单一Bucket存储一个单元的计数
type Bucket struct {
	resetDone      *int32
	resetTimestamp *int64
	count          []*int64
	gap            time.Duration
}

func NewBucket(gap time.Duration) *Bucket {
	var resetTimestamp int64 = 0
	var resetDone int32 = 1
	b := &Bucket{
		resetDone:      &resetDone,
		resetTimestamp: &resetTimestamp,
		count:          make([]*int64, event.EventsCount()),
		gap:            gap,
	}
	for i := 0; i < len(b.count); i++ {
		var v int64 = 0
		b.count[i] = &v
	}
	return b
}

// 记录给定事件
func (b *Bucket) Record(evt event.Event, eventTime time.Time, value int64) {
	b.checkAndReset(eventTime.UnixNano())
	if evt.IsCounter() {
		atomic.AddInt64(b.count[evt], value)
	} else if evt.IsMaxUpdater() {
		retry := 3
		for retry > 0 {
			v := atomic.LoadInt64(b.count[evt])
			if v >= value {
				break
			}
			if atomic.CompareAndSwapInt64(b.count[evt], v, value) {
				break
			} else {
				retry--
			}
		}
	}
}

// 取得给定事件的计数值
func (b *Bucket) Get(start time.Time, evt event.Event) int64 {
	ts := atomic.LoadInt64(b.resetTimestamp)
	if start.UnixNano()-ts >= b.gap.Nanoseconds() {
		return 0
	}
	return atomic.LoadInt64(b.count[evt])
}

// 检查是否需要重置
func (b *Bucket) checkAndReset(now int64) {
	ts := atomic.LoadInt64(b.resetTimestamp)
	if now-ts >= b.gap.Nanoseconds() { // 如果重置时间大于每个bucket时间范围，说明计数器还在上一轮没有重置
		if atomic.CompareAndSwapInt32(b.resetDone, 1, 0) { //各个goroutine抢重置标志
			for i := 0; i < len(b.count); i++ { //重置计数器
				atomic.StoreInt64(b.count[i], 0)
			}
			atomic.StoreInt64(b.resetTimestamp, time.Now().UnixNano()) // 设置最新重置时间
			atomic.StoreInt32(b.resetDone, 1)                          // 恢复重置标志
		} else {
			for i := 0; i < 50; i++ {
				ts := atomic.LoadInt64(b.resetTimestamp)
				if ts >= now {
					break
				}
				time.Sleep(b.gap / 100) // 没抢到的睡100微秒 等重置完成
			}
		}
	}
}
