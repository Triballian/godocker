package test2

import "fmt"

// TierTwoTank Tier Two Tank Struct
type TierTwoTank struct {
	Tank
	tierThreeTanks []TierThreeTank
}

// TierOneTank Tier One Tank Struct
type TierOneTank struct {
	Tank
	TierTwoTanks []TierTwoTank
}
// Tank base of all tanks
type Tank struct {
	Name string
}

// MachineGun Tier Two Tank
var MachineGun = TierTwoTank{
	Name:           "Machine Gun",
	tierThreeTanks: []TierThreeTank{Destroyer, Gunner, Sprayer},
}

// FlankGuard Tiew Two Tank
var FlankGuard = TierTwoTank{
	Name:           "Flank Gaurd",
	tierThreeTanks: []TierThreeTank{TriAngle, QuadTankGuard, TwinFlankGuard, Auto3},
}

// BaseTank Tier one tank or Base Tank
var BaseTank = TierOneTank{
	Name:         "Base Tank",
	TierTwoTanks: []TierTwoTank{MachineGun, FlankGuard},
}

func main(){
	if Basetank
}