package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type dtank struct {
	dtname string
	gconsole
}

type gconsole struct {
	gameStatsBuild
	colors []netReplaceColor
}

type gameStatsBuild struct {
	stats []int
}

type netReplaceColor struct {
	id    int
	color []string
}

// var dtanks = map[string]dtank

// not sure yet, how I'm going to pick the initial tank
// type initialTank struct {
// 	tierOneTank
// 	tierThreeTank
// }

// Tank allows all tanks to inherit Name string property
type Tank struct {
	Name string
}

// TierFourTank Tier Four Tank Struct
type TierFourTank struct {
	Tank
}

// TierFourSpecial Tier Four Tank Special, for Spike, Auto Smasher and LandMine
type TierFourSpecial struct {
	Tank
}

// TierThreeTank Tier Three Tank Struct
type TierThreeTank struct {
	Tank
	tierFourTanks []TierFourTank
}

// TierThreeSpecial Tier 3 special tank for smasher, only available to the Base Tank if the tier 2 tank is not slected.
type TierThreeSpecial struct {
	Tank
	tierFourTanks []TierFourTank
}

// TierTwoTank Tier Two Tank Strucute if the Machine Gun does not select a tier 3 tank then the Tier 4 special tank Spayer becomes Available
type TierTwoTank struct {
	Tank
	tierThreeTanks  []TierThreeTank
	tierFourSpecial TierFourSpecial
}

// TierOneTank Tier One Tank Struct if tier two tank is not selected then the tier 3 special tank Smasher becomes available.
type TierOneTank struct {
	Tank
	tierTwoTanks     []TierTwoTank
	tierThreeSpecial TierThreeSpecial
}

// LandMine Tier Four Tank
var LandMine TierFourTank

// AutoSmasher Tier Four Tank
var AutoSmasher TierFourTank

// Spike Tier Four Tank
var Spike TierFourTank

// Triplet Tier Four Tank
var Triplet TierFourTank

// PentaShot Tier Four Tank
var PentaShot TierFourTank

// SpreadShot Tier Four Tank
var SpreadShot TierFourTank

// OctoTankQuad Tier Four Tank Quad Derivitve
var OctoTankQuad TierFourTank

// Auto5 Tier Four Tank
var Auto5 TierFourTank

// TripleTwinTwin Tier Four Tank Twin Derivitive
var TripleTwinTwin TierFourTank

// BattleshipTwin Tier Four Tank Twin Derivitive
var BattleshipTwin TierFourTank

// Ranger Tier Four Tank
var Ranger TierFourTank

// Stalker Tier Four Tank
var Stalker TierFourTank

// Overlord Tier Four Tank
var Overlord TierFourTank

// Necromancer Tier Four Tank
var Necromancer TierFourTank

// Manager Tier Four Tank
var Manager TierFourTank

// OverTrapperSeer Tier Four Tank Overseer Derivitive
var OverTrapperSeer TierFourTank

// BattleshipSniper Tier Four Tank Derives from Sniper
var BattleshipSniper TierFourTank

// Factory Tier Four Tank
var Factory TierFourTank

// Predator Tier Four Tank
var Predator TierFourTank

// StreamlinerHunter Tier Four Tank Derives from Hunter
var StreamlinerHunter TierFourTank

// Hibrid Tier Four Tank
var Hibrid TierFourTank

// Annihilator Tier Four Tank
var Annihilator TierFourTank

// Sprayer Tier Four Tank only if the Machine Gun tank didn't upgrade to tier three.
var Sprayer TierFourSpecial

// Skimmer Tier Four Tank
var Skimmer TierFourTank

// AutoGunnerGunner Tier Four Tank Derives from Gunner Tank
var AutoGunnerGunner TierFourTank

// GunnerTrapperGunner Tier Four Tank Derives from Gunner Tank
var GunnerTrapperGunner TierFourTank

// StreamlinerGunner Tier Four Tank Derives from Gunner Tank
var StreamlinerGunner TierFourTank

// TriTrapper Tier Four Tank
var TriTrapper TierFourTank

// GunnerTrapperSniper Tier Four Tank Derives from Sniper Tank
var GunnerTrapperSniper TierFourTank

// OverTrapperSniper Tier Four Tank Derives from Sniper Tank
var OverTrapperSniper TierFourTank

// MegaTrapper Tier Four Tank
var MegaTrapper TierFourTank

// AutoTrapper Tier Four Tank
var AutoTrapper TierFourTank

// Booster Tier Four Tank
var Booster TierFourTank

// Fighter Tier Four Tank
var Fighter TierFourTank

// OctoTankFlank Tier Four Tank Derives form Flank Tank
var OctoTankFlank TierFourTank

// Smasher Tier Three Tank only if the Basic Gun tank didn't upgrade to tier tow.
var Smasher = TierThreeSpecial{
	tierFourTanks: []TierFourTank{LandMine, AutoSmasher, Spike},
}

// Auto5Quad Tier Four Tank Derrives from Quad
var Auto5Quad TierFourTank

// TripleTwinFlank Tier Four Tank
var TripleTwinFlank TierFourTank

// BattleshipFlank TierFourTank
var BattleshipFlank TierFourTank

// Auto5Auto Tier Four Tank
var Auto5Auto TierFourTank

// AutoGunnerAuto Tier Four Tank
var AutoGunnerAuto TierFourTank

// TripleShot Tier Three Tank
var TripleShot = TierThreeTank{
	tierFourTanks: []TierFourTank{Triplet, PentaShot, SpreadShot},
}

// QuadTankTwin Tier Three Tank
var QuadTankTwin = TierThreeTank{
	tierFourTanks: []TierFourTank{OctoTankQuad, Auto5},
}

// TwinFlankTwin Tier Three Tank
var TwinFlankTwin = TierThreeTank{
	tierFourTanks: []TierFourTank{TripleTwinTwin, BattleshipTwin},
}

// Assassin Tier Three Tank
var Assassin = TierThreeTank{
	tierFourTanks: []TierFourTank{Ranger, Stalker},
}

// Overseer Tier Three Tank
var Overseer = TierThreeTank{
	tierFourTanks: []TierFourTank{Overlord, Necromancer, Manager, OverTrapperSeer, BattleshipSniper, Factory},
}

// Hunter Tier Three Tank
var Hunter = TierThreeTank{
	tierFourTanks: []TierFourTank{TriTrapper, StreamlinerHunter},
}

// Trapper Tier Three Tank
var Trapper = TierThreeTank{
	tierFourTanks: []TierFourTank{TriTrapper, GunnerTrapperSniper, OverTrapperSeer, MegaTrapper, AutoTrapper},
}

// Destroyer Tier Three Tank
var Destroyer = TierThreeTank{
	tierFourTanks: []TierFourTank{Hibrid, Annihilator, Skimmer},
}

// Gunner Tier Three Tank
var Gunner = TierThreeTank{
	tierFourTanks: []TierFourTank{AutoGunnerGunner, GunnerTrapperGunner, StreamlinerGunner},
}

// TriAngle Tier Three Tank
var TriAngle = TierThreeTank{
	tierFourTanks: []TierFourTank{Booster, Fighter},
}

// QuadTankGuard Tier Three Tank
var QuadTankGuard = TierThreeTank{
	tierFourTanks: []TierFourTank{OctoTankFlank, Auto5Quad},
}

// TwinFlankGuard Tier Three Tank
var TwinFlankGuard = TierThreeTank{
	tierFourTanks: []TierFourTank{TripleTwinFlank, BattleshipFlank},
}

// Auto3 Tier Three Tank
var Auto3 = TierThreeTank{
	tierFourTanks: []TierFourTank{Auto5Auto, AutoGunnerAuto},
}

// Twin Tier Two Tank
var Twin = TierTwoTank{
	tierThreeTanks: []TierThreeTank{TripleShot, QuadTankGuard, TwinFlankTwin},
}

// Sniper Tier Two Tank
var Sniper = TierTwoTank{
	tierThreeTanks: []TierThreeTank{Assassin, Overseer, Hunter, Trapper},
}

// MachineGun Tier Two Tank Spayer only become available if the machine gun chooses no Tier Three Tank
var MachineGun = TierTwoTank{
	tierThreeTanks:  []TierThreeTank{Destroyer, Gunner},
	tierFourSpecial: Sprayer,
}

// FlankGuard Tiew Two Tank
var FlankGuard = TierTwoTank{
	tierThreeTanks: []TierThreeTank{TriAngle, QuadTankGuard, TwinFlankGuard, Auto3},
}

// BaseTank Tier one tank or Base Tank, if the base tank doe not choose a tier 2 tank then the Smasher becomes available.
var BaseTank = TierOneTank{
	tierTwoTanks:     []TierTwoTank{Twin, Sniper, MachineGun, FlankGuard},
	tierThreeSpecial: Smasher,
}

func tankNames() {
	LandMine.Name = "Land Mine"
	AutoSmasher.Name = "Auto Smasher"
	Spike.Name = "Spike"
	Triplet.Name = "Triplet"
	PentaShot.Name = "Penta Shot"
	OctoTankQuad.Name = "Octo Tank Quad"
	Smasher.Name = "Smasher"
	Auto5.Name = "Auto 5"
	TripleTwinTwin.Name = "Triple Twin"
	BattleshipTwin.Name = "Battleship Twin"
	Ranger.Name = "Ranger"
	Stalker.Name = "Stalker"
	Overlord.Name = "Overlord"
	Necromancer.Name = "Necromancer"
	Manager.Name = "Manager"
	OverTrapperSeer.Name = "Overtrapper Seer"
	BattleshipSniper.Name = "Overtrapper Seer"
	Factory.Name = "Factory"
	Predator.Name = "Predator"
	StreamlinerHunter.Name = "Streamliner Hunter"
	Hibrid.Name = "Hibrid"
	Annihilator.Name = "Annihilator"
	Sprayer.Name = "Spayer"
	Skimmer.Name = "Skimmer"
	AutoGunnerGunner.Name = "Auto Gunner Gunner"
	GunnerTrapperGunner.Name = "Gunner Trapper Gunner"
	StreamlinerGunner.Name = "Streamliner Gunner"
	TriTrapper.Name = "Tri-Trapper"
	GunnerTrapperSniper.Name = "Gunner Trapper Sniper"
	OverTrapperSniper.Name = "Overtrapper Sniper"
	MegaTrapper.Name = "Mega Trapper"
	AutoTrapper.Name = "Auto Trapper"
	Booster.Name = "Boster"
	Fighter.Name = "Fighter"
	OctoTankFlank.Name = "Octo Tank Flank"
	Auto5Quad.Name = "Auto 5 Quad"
	TripleTwinFlank.Name = "Triple Twin Flank"
	BattleshipFlank.Name = "Battleship Flank"
	Auto5Auto.Name = "Auto 5 Auto"
	AutoGunnerAuto.Name = "Auto Gunner Auto"
	TripleShot.Name = "Triple Shot"
	QuadTankTwin.Name = "Quad Tank Twin"
	TwinFlankTwin.Name = "Twin Flank Twin"
	Assassin.Name = "Assassin"
	Overseer.Name = "Overseer"
	Hunter.Name = "Hunter"
	Trapper.Name = "Trapper"
	Destroyer.Name = "Destroyer"
	Gunner.Name = "Gunner"
	TriAngle.Name = "Tri-Angle"
	QuadTankGuard.Name = "Quad Tank Guard"
	TwinFlankGuard.Name = "Twin Flank Guard"
	Auto3.Name = "Auto 3"
	Twin.Name = "Twin"
	Sniper.Name = "Sniper"
	MachineGun.Name = "Machine Gun"
	FlankGuard.Name = "Flank Gaurd"
	BaseTank.Name = "Base Tank"
}

// type searchTank int

// func (s searchTank) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	err := req.ParseForm()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	tpl.ExecuteTemplate(w, "index.gohtml", BaseTank)

// 	// fmt.Println("%v", req.Form)
// 	// io.WriteString(w, "hello from a docker container tanktool")

// 	// if eerr != nil {
// 	// 	log.Fatalln(eerr)
// 	// }
// }

// Index handler
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", BaseTank)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {

	tankNames()
	// var t searchTank
	router := httprouter.New()
	router.GET("/", Index)

	// build := []string{"Sprayer;", "net_replace_color 2 0x66FF00;", "net_replace_color 1 0xFFFF00;", "game_stats_build 654786547865478874658746547475656"}

	// fmt.Println("%v", baseTank.tierTwoTanks[1].name)

	// http.HandleFunc("/", searchTank)
	http.ListenAndServe(":8008", router)

}
