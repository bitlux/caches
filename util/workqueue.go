// TODO:
//   - loading checkpoint
//   - option for output filename
package util

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	defaultCheckpointInterval = 60 * time.Minute
	defaultNumWorkers         = 10
)

type Option func(*WorkQueue)

// SetCheckpointInterval sets the interval at which a work checkpoint is written to disk.
func SetCheckpointInterval(d time.Duration) Option {
	return func(wq *WorkQueue) {
		wq.checkpointInterval = d
	}
}

// SetNumWorkers sets the number of goroutines that are concurrently doing work.
func SetNumWorkers(i int) Option {
	return func(wq *WorkQueue) {
		wq.numWorkers = i
	}
}

type item struct {
	value    string
	complete bool
}

// WorkQueue stores a list of tasks to complete, hands each task to a user-provided function, and
// writes any "successes", as defined by the user-provided function, to disk. It checkpoints its
// progress along the way so that it can be stopped and restarted from a checkpoint while needing
// minimal work to be redone.
type WorkQueue struct {
	bar *progressbar.ProgressBar

	checkpointInterval time.Duration
	numWorkers         int

	processItem func(string, chan<- string)

	mu    sync.Mutex
	state []item
	index int
}

func stateFromString(s string) ([]item, int) {
	var state []item
	completed := 0
	for line := range strings.Lines(s) {
		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			Must(fmt.Errorf("unexpected number of fields: %q", line))
		}
		c, err := strconv.Atoi(tokens[1])
		Must(err)
		completed += c
		state = append(state, item{value: tokens[0], complete: c == 1})
	}
	return state, completed
}

func bar(size int) *progressbar.ProgressBar {
	return progressbar.NewOptions(size,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowIts(),
		progressbar.OptionShowCount(),
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionThrottle(time.Second),
	)
}

// NewFromCheckpoint creates a WorkQueue from a checkpoint file that an earlier WorkQueue has
// written.
func NewFromCheckpoint(fname string, options ...Option) *WorkQueue {
	wq := &WorkQueue{
		checkpointInterval: defaultCheckpointInterval,
		numWorkers:         defaultNumWorkers,
	}

	for _, opt := range options {
		opt(wq)
	}

	contents, err := os.ReadFile(fname)
	Must(err)
	comp := 0
	wq.state, comp = stateFromString(string(contents))
	wq.bar = bar(len(wq.state))
	Must(wq.bar.Add(comp))
	return wq
}

// NewWorkQueue creates a WorkQueue with a list of work to perform.
func NewWorkQueue(items []string, f func(string, chan<- string), options ...Option) *WorkQueue {
	wq := &WorkQueue{
		checkpointInterval: defaultCheckpointInterval,
		numWorkers:         defaultNumWorkers,
		processItem:        f,
	}

	for _, opt := range options {
		opt(wq)
	}

	for _, s := range items {
		wq.state = append(wq.state, item{value: s})
	}
	wq.bar = bar(len(wq.state))
	return wq
}

// Run starts the work and blocks until it is completed.
func (wq *WorkQueue) Run() {
	go wq.checkpoint()

	successCh := make(chan string)
	go writeSuccess(successCh)

	sem := make(chan struct{}, wq.numWorkers)
	for value, ok := wq.next(); ok; value, ok = wq.next() {
		//fmt.Println("value is", value)
		sem <- struct{}{}
		go func() {
			wq.processItem(value, successCh)
			wq.markFinished(value)
			<-sem
		}()
	}

	// Wait for completion
	for n := wq.numWorkers; n > 0; n-- {
		sem <- struct{}{}
	}

	close(successCh)
	fmt.Println()
}

func writeSuccess(ch <-chan string) {
	fp, err := os.Create(fmt.Sprintf("success_%s", time.Now().Format("06-01-02_150405")))
	Must(err)

	for t := range ch {
		_, err := fmt.Fprintf(fp, "%s %s\n", t, time.Now().Format(time.DateTime))
		Must(err)
	}
	Must(fp.Close())
}

func (wq *WorkQueue) markFinished(value string) {
	Must(wq.bar.Add(1))

	wq.mu.Lock()
	defer wq.mu.Unlock()

	i := slices.IndexFunc(wq.state, func(it item) bool {
		return it.value == value
	})
	wq.state[i].complete = true
}

func (wq *WorkQueue) next() (string, bool) {
	wq.mu.Lock()
	defer wq.mu.Unlock()

	if wq.index >= len(wq.state) {
		return "", false
	}
	ret := wq.state[wq.index].value
	wq.index++
	return ret, true
}

func (wq *WorkQueue) checkpoint() {
	for range time.NewTicker(wq.checkpointInterval).C {
		var sb strings.Builder
		wq.mu.Lock()
		for _, s := range wq.state {
			i := 0
			if s.complete {
				i = 1
			}
			_, err := sb.WriteString(fmt.Sprintf("%s %d\n", s.value, i))
			Must(err)
		}

		wq.mu.Unlock()

		fname := fmt.Sprintf("checkpoint_%s", time.Now().Format("06-01-02_150405"))
		Must(os.WriteFile(fname, []byte(sb.String()), 0600))
	}
}
