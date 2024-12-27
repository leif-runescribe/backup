package forces

import (
	"errors"

	"github.com/leif-runescribe/westeros-roster/units"
)

type Army struct {
	armyName string
	Lord     string
	Units    []*units.Unit
	Gold     float32
}

func newArmy(lord string, name string) *Army {
	return &Army{
		armyName: name,
		Lord:     lord,
		Units:    []*units.Unit{},
		Gold:     400.0,
	}
}

func (a *Army) purchaseUnit(u *units.Unit) error {
	if a.Gold < u.price {
		return errors.New("Unsufficient gold for :", u.name)
	}
	a.Units = append(a.Units, u)
	a.Gold -= u.price
	return nil

}
func (a *Army) RemoveUnit(unitID int) error {
	for i, u := range a.Units {
		if u.unitID == unitID {
			a.Units = append(a.Units[:i], a.Units[i+1:]...)
			return nil
		}
	}
	return errors.New("unit not found")
}
