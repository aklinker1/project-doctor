package log

import (
	"fmt"
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

	stop := func(e error) {
		success = e == nil
		done = true
		time.Sleep(sleepDuration)
	}
	return stop, func() {
		position := 0
		for !done {
			fmt.Fprintf(w, " %s %s\n", Loading(frames[position]), status())
			time.Sleep(sleepDuration)
			position = (position + 1) % frameCount
		}
		if success {
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
