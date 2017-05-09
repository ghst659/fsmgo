package fsm

import (
)

type Machine interface {
	Process(inData interface{}) (outData interface{}, err error)
	RegisterState(key string, s State) error
	SetState(key string) error
	GetState() (key string, err error)
}

type State interface {
	Process(inData interface{}) (nextState string, outData interface{}, err error)
}

