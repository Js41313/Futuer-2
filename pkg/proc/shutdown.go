//go:build linux || darwin

package proc

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Js41313/Futuer-2/pkg/threading"
)

const (
	//nolint:unused
	wrapUpTime = time.Second
	// why we use 5500 milliseconds is because most of our queue are blocking mode with 5 seconds
	waitTime = 5500 * time.Millisecond
)

var (
	wrapUpListeners          = new(listenerManager)
	shutdownListeners        = new(listenerManager)
	delayTimeBeforeForceQuit = waitTime
)

// AddShutdownListener adds fn as a shutdown listener.
// The returned func can be used to wait for fn getting called.
func AddShutdownListener(fn func()) (waitForCalled func()) {
	return shutdownListeners.addListener(fn)
}

// AddWrapUpListener adds fn as a wrap up listener.
// The returned func can be used to wait for fn getting called.
func AddWrapUpListener(fn func()) (waitForCalled func()) {
	return wrapUpListeners.addListener(fn)
}

// SetTimeToForceQuit sets the waiting time before force quitting.
func SetTimeToForceQuit(duration time.Duration) {
	delayTimeBeforeForceQuit = duration
}

// Shutdown calls the registered shutdown listeners, only for test purpose.
func Shutdown() {
	shutdownListeners.notifyListeners()
}

// WrapUp wraps up the process, only for test purpose.
func WrapUp() {
	wrapUpListeners.notifyListeners()
}

//nolint:unused
func gracefulStop(signals chan os.Signal, sig syscall.Signal) {
	signal.Stop(signals)

	go wrapUpListeners.notifyListeners()

	time.Sleep(wrapUpTime)
	go shutdownListeners.notifyListeners()

	time.Sleep(delayTimeBeforeForceQuit - wrapUpTime)
	_ = syscall.Kill(syscall.Getpid(), sig)
}

type listenerManager struct {
	lock      sync.Mutex
	waitGroup sync.WaitGroup
	listeners []func()
}

func (lm *listenerManager) addListener(fn func()) (waitForCalled func()) {
	lm.waitGroup.Add(1)

	lm.lock.Lock()
	lm.listeners = append(lm.listeners, func() {
		defer lm.waitGroup.Done()
		fn()
	})
	lm.lock.Unlock()

	return func() {
		lm.waitGroup.Wait()
	}
}

func (lm *listenerManager) notifyListeners() {
	lm.lock.Lock()
	defer lm.lock.Unlock()

	group := threading.NewRoutineGroup()
	for _, listener := range lm.listeners {
		group.RunSafe(listener)
	}
	group.Wait()

	lm.listeners = nil
}
