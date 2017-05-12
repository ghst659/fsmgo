package fsm

import ()

type Machine interface {
	Process(inData interface{}) (outData interface{}, err error)
	RegisterState(s State) error
	SetCurrentState(key string) error
	CurrentState() (key string, err error)
}

type State interface {
	Name() string
	Process(inData interface{}) (nextState string, outData interface{}, err error)
}
