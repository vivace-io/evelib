package util

import (
	"errors"
	"time"
)

// Limiter is used primarily to limit calls to services and APIs that have rate
// limiting rules. To use effectively, code that requires rate limiting should
// first call Limiter.Connect() which will block until the rules/resources
// permit a new connection to proceed. After completing the connection,
// Limiter.Disconnect() should be called to release the connection back to the
// pool to allow for use elsewhere.
// Limiter does not actually Connect/Disconnect from a server, only manages when
// new connections may be created.
type Limiter struct {
	maxRate  int
	maxBurst int
	maxConn  int
	conn     int
	clear    chan bool
}

// NewLimiter returns a new, managed Limiter instance.
func NewLimiter(rate, burst, conn int) (*Limiter, error) {
	if rate <= 0 || burst <= 0 || conn <= 0 {
		return nil, errors.New("NewLimiter values rate/burst/conn must be greater than zero in value")
	}
	lim := &Limiter{
		maxRate:  rate,
		maxBurst: burst,
		maxConn:  conn,
		conn:     0,
	}
	lim.manage()
	return lim, nil
}

// Connect blocks until a new connection is allowed under the rules defined in
// the Limiter allow it and increments the current connection count. In other
// words, once there exists an available open connection AND it is within the
// rate/burst limit, the function returns and allows execution to continue.
func (lim *Limiter) Connect() {
	select {
	case <-lim.clear:
		for {
			if lim.conn < lim.maxConn {
				return
			}
		}
	}
}

// Disconnect will decrement the number of current connections registered with
// the Limiter, allowing waiting connections to begin. Should be called after
// Limiter.Connect() is called, otherwise connections will never be released and
// goroutines waiting to connect will be stuck.
func (lim *Limiter) Disconnect() {
	lim.conn--
}

func (lim *Limiter) manage() {
	lim.clear = make(chan bool, lim.maxBurst)
	go func() {
		for {
			lim.clear <- true
			time.Sleep(time.Duration(lim.maxRate/1000) * time.Millisecond)
		}
	}()
}
