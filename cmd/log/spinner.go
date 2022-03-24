package log

import (
	"fmt"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

func Spinner(frames []string, status func() string) (func(error), func()) {
	frameCount := len(frames)
	w := uilive.New()
	w.Start()
	done := false
	success := false
	sleepDuration := 100 * time.Millisecond

	var m sync.Mutex
	isDone := func() bool {
		m.Lock()
		res := done
		m.Unlock()
		return res
	}
	isSuccess := func() bool {
		m.Lock()
		res := success
		m.Unlock()
		return res
	}

	var wg sync.WaitGroup
	wg.Add(1)
	stop := func(e error) {
		m.Lock()
		success = e == nil
		done = true
		m.Unlock()
		wg.Wait()
	}
	return stop, func() {
		position := 0
		for !isDone() {
			fmt.Fprintf(w, " %s %s\n", Loading(frames[position]), status())
			time.Sleep(sleepDuration)
			position = (position + 1) % frameCount
		}
		wg.Done()
		if isSuccess() {
			fmt.Fprintf(w, Success(" ✔ %s\n"), status())
		} else {
			fmt.Fprintf(w, Error(" ✘ %s\n"), status())
		}
		w.Stop()
	}
}

// BrailSpinner just calls `Spinner` with frames that look like brail
func BrailSpinner(status func() string) (func(error), func()) {
	return Spinner([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}, status)
}
