package work

import (
	"fmt"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bitlux/caches/util"
	"github.com/schollz/progressbar/v3"
)

var writeInterval = 30 * time.Minute

type item struct {
	value    string
	complete bool
}

func (i item) String() string {
	return fmt.Sprintf("%s %d", i.value, func(b bool) int {
		if b {
			return 1
		}
		return 0
	}(i.complete))
}

type Queue struct {
	bar *progressbar.ProgressBar

	mu        sync.Mutex
	state     []item
	index     int
	lastWrite time.Time
}

func stateFromSeq(seq iter.Seq[string]) []item {
	var state []item
	for s := range seq {
		state = append(state, item{value: s})
	}
	return state
}

func stateFromString(s string) ([]item, int) {
	var state []item
	completed := 0
	for line := range strings.Lines(s) {
		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			util.Must(fmt.Errorf("unexpected number of fields: %q", line))
		}
		c, err := strconv.Atoi(tokens[1])
		util.Must(err)
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

func NewFromCheckpoint(fname string) *Queue {
	wq := &Queue{
		lastWrite: time.Now(),
	}

	contents, err := os.ReadFile(fname)
	util.Must(err)
	comp := 0
	wq.state, comp = stateFromString(string(contents))
	wq.bar = bar(len(wq.state))
	util.Must(wq.bar.Add(comp))
	return wq
}

func NewFromSeq(seq iter.Seq[string]) *Queue {
	wq := &Queue{
		state:     stateFromSeq(seq),
		lastWrite: time.Now(),
	}
	wq.bar = bar(len(wq.state))
	return wq
}

func (wq *Queue) MarkFinished(value string) {
	util.Must(wq.bar.Add(1))

	wq.mu.Lock()
	defer wq.mu.Unlock()

	i := slices.IndexFunc(wq.state, func(it item) bool {
		return it.value == value
	})
	wq.state[i].complete = true

	if time.Since(wq.lastWrite) > writeInterval {
		wq.lastWrite = time.Now()
		go dump(slices.Clone(wq.state))
	}
}

func dump(state []item) {
	var sb strings.Builder
	for _, s := range state {
		_, err := sb.WriteString(s.String() + "\n")
		util.Must(err)
	}

	util.Must(os.WriteFile(fmt.Sprintf("%s_%s", "checkpoint", time.Now().Format("06-01-02_150405")),
		[]byte(sb.String()),
		0600))
}

func (wq *Queue) Next() (string, bool) {
	wq.mu.Lock()
	defer wq.mu.Unlock()

	for ; ; wq.index++ {
		if wq.index >= len(wq.state) {
			return "", false
		}
		if !wq.state[wq.index].complete {
			ret := wq.state[wq.index].value
			wq.index++
			return ret, true
		}
	}
}
