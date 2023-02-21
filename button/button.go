package button

import (
	"machine"
	"time"
)

var DebounceDelay = 10 * time.Millisecond // used to be 50ms

type Combination uint8

const (
	MultiOr Combination = iota
	MultiAnd
	MultiXor
)

func NewButton(pin machine.Pin, mode machine.PinMode) *PushButton {
	pin.Configure(machine.PinConfig{Mode: mode})
	return &PushButton{
		btn: pin,
	}
}

func NewToggleButton(button *PushButton, initialState ...bool) *ToggleButton {
	state := false
	if len(initialState) > 0 {
		state = initialState[0]
	}
	return &ToggleButton{
		btn:   button,
		state: state,
	}
}

func NewMultiButton(combination Combination, buttons ...Button) *MultiButton {
	return &MultiButton{buttons: buttons, combination: combination}
}

type Button interface {
	Get() bool
}

type PushButton struct {
	state      bool
	lastUpdate time.Time
	btn        machine.Pin
}

func (b *PushButton) Get() bool {
	s := b.btn.Get()
	// debounce
	if s != b.state {
		if time.Since(b.lastUpdate) > DebounceDelay {
			b.state = s
		}
	} else {
		b.lastUpdate = time.Now()
	}
	return b.state
}

type ToggleButton struct {
	btn       *PushButton
	state     bool
	lastState bool
}

func (t *ToggleButton) Get() bool {
	s := t.btn.Get()
	if s != t.lastState && s {
		t.state = !t.state
	}

	t.lastState = s
	return t.state
}

type MultiButton struct {
	buttons     []Button
	combination Combination
}

func (m *MultiButton) Get() bool {
	state := false
	switch m.combination {
	case MultiAnd:
		state = true
		for _, b := range m.buttons {
			if !b.Get() {
				state = false
				break
			}
		}
	case MultiOr:
		for _, b := range m.buttons {
			if b.Get() {
				state = true
				break
			}
		}
	case MultiXor:
		for _, b := range m.buttons {
			state = state != b.Get()
		}
	}
	return state
}

func (m *MultiButton) GetAll() []bool {
	states := make([]bool, len(m.buttons))
	for i, b := range m.buttons {
		states[i] = b.Get()
	}
	return states
}
