package main

import (
	"fmt"
	"reflect"
)

// TierTwoTank Tier Two Tank Struct
type TierTwoTank struct {
	Tank
}

// TierOneTank Tier One Tank Struct
type TierOneTank struct {
	Tank
}

// Tank base of all tanks
type Tank struct {
	Name string
}

// MachineGun Tier Two Tank
var MachineGun TierTwoTank

// FlankGuard Tiew Two Tank
var FlankGuard TierTwoTank

// BaseTank Tier one tank or Base Tank
var BaseTank = TierOneTank{
// TierTwoTanks: []TierTwoTank{&MachineGun, &FlankGuard},
}

func (t Tank) populate() {
	t.Name = fmt.Sprintf("%v", reflect.TypeOf(t))

}

func tankNames() {
	FlankGuard.Name = "FlangGuard"
}

func main() {
	tankNames()
	// tank := []TierTwoTank{FlankGuard, MachineGun}
	// populate([]Tank(tank))

	fmt.Println(FlankGuard.Name)
	fmt.Println(reflect.TypeOf(FlankGuard))

}
