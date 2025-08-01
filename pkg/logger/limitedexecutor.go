package logger

import (
	"sync/atomic"
	"time"

	"github.com/Js41313/Futuer-2/pkg/syncx"
	"github.com/Js41313/Futuer-2/pkg/timex"
)

type limitedExecutor struct {
	threshold time.Duration
	lastTime  *syncx.AtomicDuration
	discarded uint32
}

func newLimitedExecutor(milliseconds int) *limitedExecutor {
	return &limitedExecutor{
		threshold: time.Duration(milliseconds) * time.Millisecond,
		lastTime:  syncx.NewAtomicDuration(),
	}
}

func (le *limitedExecutor) logOrDiscard(execute func()) {
	if le == nil || le.threshold <= 0 {
		execute()
		return
	}

	now := timex.Now()
	if now-le.lastTime.Load() <= le.threshold {
		atomic.AddUint32(&le.discarded, 1)
	} else {
		le.lastTime.Set(now)
		discarded := atomic.SwapUint32(&le.discarded, 0)
		if discarded > 0 {
			Errorf("Discarded %d error messages", discarded)
		}

		execute()
	}
}
