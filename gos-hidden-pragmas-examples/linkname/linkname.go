package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("goroutineid:", goroutineid())
		}()
	}
	wg.Wait()
}

func goroutineid() int64 {
	m := (*m)(unsafe.Pointer(runtime_getm()))
	g := m.curg
	return g.goid
}

//go:linkname runtime_getm runtime.getm
func runtime_getm() uintptr

type m struct {
	g0      *g     // goroutine with scheduling stack
	morebuf gobuf  // gobuf arg to morestack
	divmod  uint32 // div/mod denominator for arm - known to liblink

	// Fields not known to debuggers.
	procid     uint64       // for debuggers, but offset not hard-coded
	gsignal    *g           // signal-handling g
	goSigStack gsignalStack // Go-allocated signal handling stack
	sigmask    sigset       // storage for saved signal mask
	tls        [6]uintptr   // thread-local storage (for x86 extern register)
	mstartfn   func()
	curg       *g // current running goroutine
}

type g struct {
	// Stack parameters.
	// stack describes the actual stack memory: [stack.lo, stack.hi).
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	// stackguard1 is the stack pointer compared in the C stack growth prologue.
	// It is stack.lo+StackGuard on g0 and gsignal stacks.
	// It is ~0 on other goroutine stacks, to trigger a call to morestackc (and crash).
	stack       stack   // offset known to runtime/cgo
	stackguard0 uintptr // offset known to liblink
	stackguard1 uintptr // offset known to liblink

	_panic       unsafe.Pointer // innermost panic - offset known to liblink
	_defer       unsafe.Pointer // innermost defer
	m            unsafe.Pointer // current m; offset known to arm liblink
	sched        gobuf
	syscallsp    uintptr        // if status==Gsyscall, syscallsp = sched.sp to use during gc
	syscallpc    uintptr        // if status==Gsyscall, syscallpc = sched.pc to use during gc
	stktopsp     uintptr        // expected sp at top of stack, to check in traceback
	param        unsafe.Pointer // passed parameter on wakeup
	atomicstatus uint32
	stackLock    uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid         int64
}

type sigset uint32

type guintptr uintptr

type gobuf struct {
	// The offsets of sp, pc, and g are known to (hard-coded in) libmach.
	//
	// ctxt is unusual with respect to GC: it may be a
	// heap-allocated funcval so write require a write barrier,
	// but gobuf needs to be cleared from assembly. We take
	// advantage of the fact that the only path that uses a
	// non-nil ctxt is morestack. As a result, gogo is the only
	// place where it may not already be nil, so gogo uses an
	// explicit write barrier. Everywhere else that resets the
	// gobuf asserts that ctxt is already nil.
	sp   uintptr
	pc   uintptr
	g    guintptr
	ctxt unsafe.Pointer // this has to be a pointer so that gc scans it
	ret  uint64
	lr   uintptr
	bp   uintptr // for GOEXPERIMENT=framepointer
}

// Stack describes a Go execution stack.
// The bounds of the stack are exactly [lo, hi),
// with no implicit data structures on either side.
type stack struct {
	lo uintptr
	hi uintptr
}

// gsignalStack saves the fields of the gsignal stack changed by
// setGsignalStack.
type gsignalStack struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr
	stktopsp    uintptr
}
