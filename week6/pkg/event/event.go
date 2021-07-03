package event

import (
	"context"
	"time"
)

type Event int

const (
	SUCCESS Event = iota
	FAILURE
	TIMEOUT
	SHORT_CIRCUITED
	THREAD_POOL_REJECTED
	SEMAPHORE_REJECTED
	BAD_REQUEST
	FALLBACK_SUCCESS
	FALLBACK_FAILURE
	FALLBACK_REJECTION
	FALLBACK_DISABLED
	FALLBACK_MISSING
	EXCEPTION_THROWN
	COMMAND_MAX_ACTIVE
	EMIT
	FALLBACK_EMIT
	THREAD_EXECUTION
	THREAD_MAX_ACTIVE
	COLLAPSED
	RESPONSE_FROM_CACHE
	COLLAPSER_REQUEST_BATCHED
	COLLAPSER_BATCH
)

type eventType int

const (
	counter    eventType = 1
	maxUpdater eventType = 2
)

type eventMeta struct {
	_type eventType
	name  string
}

var events = []eventMeta{
	{
		name:  "SUCCESS",
		_type: 1,
	},
	{
		name:  "FAILURE",
		_type: 1,
	},
	{
		name:  "TIMEOUT",
		_type: 1,
	},
	{
		name:  "SHORT_CIRCUITED",
		_type: 1,
	},
	{
		name:  "THREAD_POOL_REJECTED",
		_type: 1,
	},
	{
		name:  "SEMAPHORE_REJECTED",
		_type: 1,
	},
	{
		name:  "BAD_REQUEST",
		_type: 1,
	},
	{
		name:  "FALLBACK_SUCCESS",
		_type: 1,
	},
	{
		name:  "FALLBACK_FAILURE",
		_type: 1,
	},
	{
		name:  "FALLBACK_REJECTION",
		_type: 1,
	},
	{
		name:  "FALLBACK_DISABLED",
		_type: 1,
	},
	{
		name:  "FALLBACK_MISSING",
		_type: 1,
	},
	{
		name:  "EXCEPTION_THROWN",
		_type: 1,
	},
	{
		name:  "COMMAND_MAX_ACTIVE",
		_type: 2,
	},
	{
		name:  "EMIT",
		_type: 1,
	},
	{
		name:  "FALLBACK_EMIT",
		_type: 1,
	},
	{
		name:  "THREAD_EXECUTION",
		_type: 1,
	},
	{
		name:  "THREAD_MAX_ACTIVE",
		_type: 2,
	},
	{
		name:  "COLLAPSED",
		_type: 1,
	},
	{
		name:  "RESPONSE_FROM_CACHE",
		_type: 1,
	},
	{
		name:  "COLLAPSER_REQUEST_BATCHED",
		_type: 1,
	},
	{
		name:  "COLLAPSER_BATCH",
		_type: 1,
	},
}

func EventsCount() int {
	return len(events)
}

func (t Event) IsCounter() bool {
	return events[t]._type == counter
}

func (t Event) IsMaxUpdater() bool {
	return events[t]._type == maxUpdater
}

func (t Event) Name() string {
	return events[t].name
}

type EventReceiver interface {
	ReceiveEvent(ctx context.Context, t time.Time, event Event, value int64)
}
