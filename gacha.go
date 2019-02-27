package gacha

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Capsule struct {
	Item        interface{}
	Probability float64
}

type blindBox struct {
	capsules []Capsule
	rand     *rand.Rand
	sync.RWMutex
}

type noResultError struct {
	boxLen int64
	em     float64
	sum    float64
	pos    float64
}

func (e *noResultError) Error() string {
	return fmt.Sprintf("gacha: boxLen=%d, em=%f, sum=%f, pos=%f: %s",
		e.boxLen,
		e.em,
		e.sum,
		e.pos,
		"No result error")
}

func New() *blindBox {
	return &blindBox{
		capsules: make([]Capsule, 0),
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (b *blindBox) Empty() bool {
	return b.Len() == 0
}

func (b *blindBox) Get() (*Capsule, error) {
	ret, err := b.Peek()

	if err == nil {
		b.remove(*ret)
	}

	return ret, err
}

func (b *blindBox) Len() int {
	b.Lock()
	defer b.Unlock()

	return len(b.capsules)
}

func (b *blindBox) Peek() (*Capsule, error) {
	if b.Empty() {
		return nil, &noResultError{
			boxLen: int64(len(b.capsules)),
			em:     0,
			sum:    0,
			pos:    0,
		}
	}

	b.Lock()
	defer b.Unlock()

	var em, sum, pos float64
	for _, capsule := range b.capsules {
		sum += capsule.Probability
	}

	pos = b.rand.Float64() * sum

	for _, capsule := range b.capsules {
		if em += capsule.Probability; em > pos {
			return &capsule, nil
		}
	}

	return nil, &noResultError{
		boxLen: int64(len(b.capsules)),
		em:     em,
		sum:    sum,
		pos:    pos,
	}
}

func (b *blindBox) Put(capsules ...Capsule) {
	if len(capsules) == 0 {
		return
	}

	b.Lock()
	defer b.Unlock()

	// Remove duplicates
	for _, capsule := range capsules {
		if pos := b.indexOf(capsule); pos > 0 {
			b.capsules = append(b.capsules[:pos], b.capsules[pos+1:]...)
		}
	}

	b.capsules = append(b.capsules, capsules...)
}

func (b *blindBox) indexOf(capsule Capsule) int {
	for i, c := range b.capsules {
		if c.Item == capsule.Item {
			return i
		}
	}

	return -1
}

// No locks, use for internally
func (b *blindBox) remove(capsules ...Capsule) {
	for _, capsule := range capsules {
		idx := b.indexOf(capsule)

		if idx == -1 {
			continue
		}

		b.capsules = append(b.capsules[:idx], b.capsules[idx+1:]...)
	}
}

func (b *blindBox) Remove(capsules ...Capsule) {
	b.Lock()
	defer b.Unlock()

	b.remove(capsules...)
}
