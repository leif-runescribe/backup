package game

import (
	"errors"
	"sync"

	"github.com/leif-runescribe/westeros-roster/forces"
)

type Realm struct {
	Armies map[string]*forces.Army
	mu     sync.RWMutex
}

func newRealm() *Realm {
	return &Realm{
		Armies: make(map[string]*forces.Army),
	}
}

func (r *Realm) createArmy(armyName string, lordName string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.Armies[armyName || r.Armies[lordName]]; exists {
		return nil, errors.New("imposter ")
	}
}
