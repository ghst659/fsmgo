package main

import (
	"flag"
	"fmt"
	"github.com/ghst659/fsmgo/fsm"
	"strings"
)

type Ping struct{}
type Pong struct{}

func (s *Ping) Name() string {
	return "Ping"
}

func (s *Pong) Name() string {
	return "Pong"
}

func (s *Ping) Process(inData interface{}) (nextKey string, outData interface{}, err error) {
	err = nil
	inStr := inData.(string)
	outData = "Ping got " + inStr
	if strings.HasPrefix(inStr, "i") {
		nextKey = "Ping"
	} else {
		nextKey = "Pong"
	}
	return
}

func (s *Pong) Process(inData interface{}) (nextKey string, outData interface{}, err error) {
	err = nil
	inStr := inData.(string)
	outData = "Pong got " + inStr
	if strings.HasPrefix(inStr, "o") {
		nextKey = "Pong"
	} else {
		nextKey = "Ping"
	}
	return
}

func main() {
	// var fp_fixed = flag.Bool("verbose", false, "run verbosely.")
	flag.Parse()
	if m, err := fsm.New(); err == nil {
		ping := new(Ping)
		pong := new(Pong)
		m.RegisterState(ping)
		m.RegisterState(pong)
		m.SetCurrentState("Ping")
		for _, arg := range flag.Args() {
			odata, _ := m.Process(arg)
			if sName, err := m.CurrentState(); err == nil {
				fmt.Println(sName, arg, odata)
			}
		}
	}
}
