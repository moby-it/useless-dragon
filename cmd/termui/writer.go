package termui

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

const CharDelay = 100 * time.Millisecond

// Narrate prints a message to the console with a delay between each character. If the user presses enter, the message is printed immediately.
func Narrate(message string, w io.Writer) {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer func() {
		term.Restore(int(os.Stdin.Fd()), oldState)
		fmt.Fprintln(w)
	}()
	wg := sync.WaitGroup{}
	mx := sync.RWMutex{}
	canceled := false
	wg.Add(1)
	go func() {
		for i, c := range message {
			mx.RLock()
			if canceled {
				fmt.Fprintf(w, "%s", message[i:])
				wg.Done()
				mx.RUnlock()
				break
			}
			mx.RUnlock()
			fmt.Fprintf(w, "%c", c)
			time.Sleep(100 * time.Millisecond)
			if i == len(message)-1 {
				canceled = true
				wg.Done()
			}
		}
	}()
	for {
		var oneChar [1]byte
		_, err := os.Stdin.Read(oneChar[:])
		if canceled {
			break
		}
		const SPACE = ' '
		if err != nil {
			break
		}
		if oneChar[0] == SPACE {
			mx.Lock()
			canceled = true
			mx.Unlock()
			break
		}
	}
	wg.Wait()
}
