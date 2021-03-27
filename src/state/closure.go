package state

import "binchunk"
import . "api"

type closure struct {
	proto *binchunk.Prototype
	goFunc GoFunction
}

func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}

func newGoClosure(f GoFunction) *closure {
	return &closure{goFunc: f}
}
