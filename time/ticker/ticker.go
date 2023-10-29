package ticker

import (
	"errors"
	"time"
)

const (
	// StopMessage message of callCh
	StopMessage = "stop"

	// StopConfirmMessage message of ctrlCh
	StopConfirmMessage = "stopped"
)

var (
	AlreadyStartedError = errors.New("ticker already started")
	NotStartedError     = errors.New("ticker not started")
)

type (
	Tick struct {
		ticker   *time.Ticker
		interval time.Duration
		callback func()
		callCh   chan string // call control message
		respCh   chan string // response for control message
	}
)

// run is a goroutine that wait ticker time timeout.
func (t *Tick) run() {
LOOP:
	for {
		select {
		case <-t.ticker.C: // timer expired.
			t.callback()
		case msg := <-t.callCh:
			if msg == StopMessage {
				t.ticker.Stop()
				break LOOP // end of goroutine.
			}
		}
	}
	t.respCh <- StopConfirmMessage
}

// Start starts tick timer. Return error if already started.
func (t *Tick) Start() error {
	if t.ticker != nil {
		return AlreadyStartedError
	}
	t.ticker = time.NewTicker(t.interval)
	go t.run()
	return nil
}

// Stop stops tick timer and unlink ticker. Return error if not started.
func (t *Tick) Stop() error {
	if t.ticker == nil {
		return NotStartedError
	}
	t.callCh <- StopMessage
	<-t.respCh     // waiting for stop confirm message
	t.ticker = nil // unlink ticker.
	return nil
}

// ChangeInterval changes tick interval.
func (t *Tick) ChangeInterval(interval time.Duration) {
	if t.ticker == nil {
		// if ticker is not running, change property only.
		t.interval = interval
		return
	}
	// if ticker is already running, stop it and change interval.
	t.callCh <- StopMessage
	<-t.respCh     // waiting for stop confirm message
	t.ticker = nil // unlink ticker.
	t.interval = interval
	_ = t.Start() // restart.
}

// NewTick returns new instance of tick. At this moment, tick timer not be running.
func NewTick(interval time.Duration, f func()) *Tick {
	return &Tick{
		ticker:   nil,
		interval: interval,
		callback: f,
		callCh:   make(chan string),
		respCh:   make(chan string),
	}
}
