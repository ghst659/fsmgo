package fsm

import (
	"errors"
)

type machine struct {
	current string
	states  map[string]State
}

func New() (m Machine, err error) {
	mach := new(machine)
	mach.current = ""
	mach.states = make(map[string]State)
	m = mach
	return
}

func (m *machine) RegisterState(s State) (err error) {
	if s == nil {
		err = errors.New("Invalid nil state.")
	} else {
		key = s.Name()
		if _, isRegistered := m.states[key]; isRegistered {
			err = errors.New("State already registered")
		} else {
			m.states[key] = s
			err = nil
		}
	}
	return
}

func (m *machine) CurrentState() (key string, err error) {
	if _, ok := m.states[m.current]; ok {
		key = m.current
	} else {
		key = ""
		err = errors.New("invalid state")
	}
	return
}

func (m *machine) SetCurrentState(key string) (err error) {
	if _, registered := m.states[key]; registered {
		m.current = key
	} else {
		err = errors.New("invalid new state")
	}
	return
}

func (m *machine) Process(inData interface{}) (outData interface{}, err error) {
	if m.current != "" {
		var nextKey string
		s := m.states[m.current]
		nextKey, outData, err = s.Process(inData)
		if err == nil {
			err = m.SetState(nextKey)
		}
	} else {
		outData = nil
		err = errors.New("Invalid nil state")
	}
	return
}
