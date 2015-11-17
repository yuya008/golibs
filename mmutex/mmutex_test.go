/*
	测试
*/
package mmutex

import (
	_ "log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestLockAndUnlock(t *testing.T) {
	mm := NewMmutex()
	// mm := sync.Mutex{}
	var sum int // is 499500
	var wg sync.WaitGroup
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(n int) {
			mm.Lock()
			sum = sum + n
			mm.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log(sum)
	if sum != 499999500000 {
		t.Error("mmutex fail!")
	}
}

func TestLockAndUnlockAndTryLock0(t *testing.T) {
	mm := NewMmutex()
	var wg sync.WaitGroup
	wg.Add(1)
	mm.Lock()
	go func() {
		if mm.TryLock() {
			t.Error("mmutex fail!")
		}
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	mm.Unlock()
	go func() {
		if !mm.TryLock() {
			t.Error("mmutex fail!")
		}
		wg.Done()
	}()
	wg.Wait()
}

func TestLockAndUnlockAndTryLock1(t *testing.T) {
	mm := NewMmutex()
	var a int32
	var b int32
	var wg sync.WaitGroup
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(n int) {
			if mm.TryLock() {
				a++
				mm.Unlock()
			} else {
				atomic.AddInt32(&b, 1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log(a)
	t.Log(b)
	if a+b != 1000000 {
		t.Error("mmutex fail!")
	}
}

func TestTryLockTimeOut(t *testing.T) {
	var timeout0 time.Duration = time.Second * 3
	var timeout1 time.Duration = timeout0 - time.Second*2
	mm := NewMmutex()
	mm.Lock()
	go func() {
		if mm.TryLockTimeOut(timeout1) {
			t.Error("mmutex fail!")
		}
	}()
	time.Sleep(timeout0)
	mm.Unlock()
	if !mm.TryLock() {
		t.Error("mmutex fail!")
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if !mm.TryLockTimeOut(timeout0) {
			t.Error("mmutex fail!")
		}
		wg.Done()
	}()
	time.Sleep(timeout1)
	mm.Unlock()
	wg.Wait()
}
