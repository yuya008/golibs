//
//	扩展mutex接口
//
package mmutex

import (
	"sync"
	"time"
)

var (
	// trylock未超时期间，重试加锁的delay时间
	trylock_test_delay time.Duration = 1 * time.Millisecond
)

type Mmutex interface {
	sync.Locker
	// 尝试加排它锁
	TryLock() bool
	// 尝试加排它锁，并带有超时
	TryLockTimeOut(time.Duration) bool
}

type mmutex struct {
	w      bool
	locker sync.Locker
	cond   *sync.Cond
}

// 制造一个Mmutex对象
func NewMmutex() Mmutex {
	m := &sync.Mutex{}
	return &mmutex{
		locker: m,
		cond:   sync.NewCond(m),
	}
}

// 排它锁
func (lock *mmutex) Lock() {
	lock.locker.Lock()
	defer lock.locker.Unlock()
	for lock.w {
		lock.cond.Wait()
	}
	lock.w = true
}

// 排它解锁
func (lock *mmutex) Unlock() {
	lock.w = false
	lock.cond.Broadcast()
}

// 尝试加排它锁，成功返回true，反之失败返回false
func (lock *mmutex) TryLock() bool {
	lock.locker.Lock()
	defer lock.locker.Unlock()
	if !lock.w {
		lock.w = true
		return true
	}
	return false
}

// 尝试加排它锁，这个接口同时还具备超时特性，如果timeout为0，那么它的行为
// 将和直接调用TryLock()无异，加锁成功返回true，反之超时失败返回false
func (lock *mmutex) TryLockTimeOut(timeout time.Duration) bool {
	t := time.After(timeout)
	for {
		if lock.TryLock() {
			return true
		}
		select {
		case <-t:
			return false
		default:
			time.Sleep(trylock_test_delay)
		}
	}
}
