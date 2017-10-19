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

// TierFourTank Tier Frou Tank Struct
type TierFourTank struct {
	Name string
}

// TierThreeTank Tier Three Tank Struct
type TierThreeTank struct {
	Name          string
	tierFourTanks []TierFourTank
}

// TierTwoTank Tier Two Tank Struct
type TierTwoTank struct {
	Name           string
	tierThreeTanks []TierThreeTank
}

// TierOneTank Tier One Tank Struct
type TierOneTank struct {
	Name         string
	TierTwoTanks []TierTwoTank
}

// LandMine Tier Four Tank
var LandMine = TierFourTank{
	Name: "Land Mine",
}

// AutoSmasher Tier Four Tank
var AutoSmasher = TierFourTank{
	Name: "Auto Smasher",
}

// Spike Tier Four Tank
var Spike = TierFourTank{
	Name: "Spike",
}

// Triplet Tier Four Tank
var Triplet = TierFourTank{
	Name: "Triplet",
}

// PentaShot Tier Four Tank
var PentaShot = TierFourTank{
	Name: "Penta Shot",
}

// SpreadShot Tier Four Tank
var SpreadShot = TierFourTank{
	Name: "Spread Shot",
}

// OctoTankQuad Tier Four Tank Quad Derivitve
var OctoTankQuad = TierFourTank{
	Name: "Octo Tank Quad",
}

// Auto5 Tier Four Tank
var Auto5 = TierFourTank{
	Name: "Auto 5",
}

// TripleTwinTwin Tier Four Tank Twin Derivitive
var TripleTwinTwin = TierFourTank{
	Name: "Triple Twin",
}

// BattleshipTwin Tier Four Tank Twin Derivitive
var BattleshipTwin = TierFourTank{
	Name: "Battleship Twin",
}

// Ranger Tier Four Tank
var Ranger = TierFourTank{
	Name: "Ranger",
}

// Stalker Tier Four Tank
var Stalker = TierFourTank{
	Name: "Stalker",
}

// Overlord Tier Four Tank
var Overlord = TierFourTank{
	Name: "Overlord",
}

// Necromancer Tier Four Tank
var Necromancer = TierFourTank{
	Name: "Necromancer",
}

// Manager Tier Four Tank
var Manager = TierFourTank{
	Name: "Manager",
}

// OverTrapperSeer Tier Four Tank Overseer Derivitive
var OverTrapperSeer = TierFourTank{
	Name: "Overtrapper Seer",
}

// BattleshipSniper Tier Four Tank Derives from Sniper
var BattleshipSniper = TierFourTank{
	Name: "Battleship Sniper",
}

// Factory Tier Four Tank
var Factory = TierFourTank{
	Name: "Factory",
}

// Predator Tier Four Tank
var Predator = TierFourTank{
	Name: "Predator",
}

// StreamlinerHunter Tier Four Tank Derives from Hunter
var StreamlinerHunter = TierFourTank{
	Name: "Streamliner Hunter",
}

// Hibrid Tier Four Tank
var Hibrid = TierFourTank{
	Name: "Hibrid",
}

// Annihilator Tier Four Tank
var Annihilator = TierFourTank{
	Name: "Annihilator",
}

// Skimmer Tier Four Tank
var Skimmer = TierFourTank{
	Name: "Skimmer",
}

// AutoGunnerGunner Tier Four Tank Derives from Gunner Tank
var AutoGunnerGunner = TierFourTank{
	Name: "Auto Gunner Gunner",
}

// GunnerTrapperGunner Tier Four Tank Derives from Gunner Tank
var GunnerTrapperGunner = TierFourTank{
	Name: "Gunner Trapper Gunner",
}

// StreamlinerGunner Tier Four Tank Derives from Gunner Tank
var StreamlinerGunner = TierFourTank{
	Name: "Streamliner Gunner",
}

// TriTrapper Tier Four Tank
var TriTrapper = TierFourTank{
	Name: "Tri-Trapper",
}

// GunnerTrapperSniper Tier Four Tank Derives from Sniper Tank
var GunnerTrapperSniper = TierFourTank{
	Name: "Gunner Trapper Sniper",
}

// OverTrapperSniper Tier Four Tank Derives from Sniper Tank
var OverTrapperSniper = TierFourTank{
	Name: "Overtrapper Sniper",
}

// MegaTrapper Tier Four Tank
var MegaTrapper = TierFourTank{
	Name: "Mega Trapper",
}

// AutoTrapper Tier Four Tank
var AutoTrapper = TierFourTank{
	Name: "Auto Trapper",
}

// Booster Tier Four Tank
var Booster = TierFourTank{
	Name: "Boster",
}

// Fighter Tier Four Tank
var Fighter = TierFourTank{
	Name: "Fighter",
}

// OctoTankFlank Tier Four Tank Derives form Flank Tank
var OctoTankFlank = TierFourTank{
	Name: "Octo Tank Flank",
}

// Auto5Quad Tier Four Tank Derrives from Quad
var Auto5Quad = TierFourTank{
	Name: "Auto 5 Quad",
}

// TripleTwinFlank Tier Four Tank
var TripleTwinFlank = TierFourTank{
	Name: "Triple Twin Flank",
}

// BattleshipFlank TierFourTank
var BattleshipFlank = TierFourTank{
	Name: "Battleship Flank",
}

// Auto5Auto Tier Four Tank
var Auto5Auto = TierFourTank{
	Name: "Auto 5 Auto",
}

// AutoGunnerAuto Tier Four Tank
var AutoGunnerAuto = TierFourTank{
	Name: "Auto Gunner Auto",
}

// Smasher Tier Three Tank
var Smasher = TierThreeTank{
	Name:          "Smasher",
	tierFourTanks: []TierFourTank{LandMine, AutoSmasher, Spike},
}

// TripleShot Tier Three Tank
var TripleShot = TierThreeTank{
	Name:          "Triple Shot",
	tierFourTanks: []TierFourTank{Triplet, PentaShot, SpreadShot},
}

// QuadTankTwin Tier Three Tank
var QuadTankTwin = TierThreeTank{
	Name:          "Quad Tank Twin",
	tierFourTanks: []TierFourTank{OctoTankQuad, Auto5},
}

// TwinFlankTwin Tier Three Tank
var TwinFlankTwin = TierThreeTank{
	Name:          "Twin Flank Twin",
	tierFourTanks: []TierFourTank{TripleTwinTwin, BattleshipTwin},
}

// Assassin Tier Three Tank
var Assassin = TierThreeTank{
	Name:          "Assassin",
	tierFourTanks: []TierFourTank{Ranger, Stalker},
}

// Overseer Tier Three Tank
var Overseer = TierThreeTank{
	Name:          "Overseer",
	tierFourTanks: []TierFourTank{Overlord, Necromancer, Manager, OverTrapperSeer, BattleshipSniper, Factory},
}

// Hunter Tier Three Tank
var Hunter = TierThreeTank{
	Name:          "Hunter",
	tierFourTanks: []TierFourTank{TriTrapper, StreamlinerHunter},
}

// Trapper Tier Three Tank
var Trapper = TierThreeTank{
	Name:          "Trapper",
	tierFourTanks: []TierFourTank{TriTrapper, GunnerTrapperSniper, OverTrapperSeer, MegaTrapper, AutoTrapper},
}

// Destroyer Tier Three Tank
var Destroyer = TierThreeTank{
	Name:          "Destroyer",
	tierFourTanks: []TierFourTank{Hibrid, Annihilator, Skimmer},
}

// Gunner Tier Three Tank
var Gunner = TierThreeTank{
	Name:          "Gunner",
	tierFourTanks: []TierFourTank{AutoGunnerGunner, GunnerTrapperGunner, StreamlinerGunner},
}

// Sprayer Tier Three Tank
var Sprayer = TierThreeTank{
	Name:          "Spayer",
	tierFourTanks: []TierFourTank{AutoGunnerGunner, GunnerTrapperGunner, StreamlinerGunner},
}

// TriAngle Tier Three Tank
var TriAngle = TierThreeTank{
	Name:          "Tri-Angle",
	tierFourTanks: []TierFourTank{Booster, Fighter},
}

// QuadTankGuard Tier Three Tank
var QuadTankGuard = TierThreeTank{
	Name:          "Quad Tank Guard",
	tierFourTanks: []TierFourTank{OctoTankFlank, Auto5Quad},
}

// TwinFlankGuard Tier Three Tank
var TwinFlankGuard = TierThreeTank{
	Name:          "Twin Flank Guard",
	tierFourTanks: []TierFourTank{TripleTwinFlank, BattleshipFlank},
}

// Auto3 Tier Three Tank
var Auto3 = TierThreeTank{
	Name:          "Auto 3",
	tierFourTanks: []TierFourTank{Auto5Auto, AutoGunnerAuto},
}

// Twin Tier Two Tank
var Twin = TierTwoTank{
	Name:           "Twin",
	tierThreeTanks: []TierThreeTank{TripleShot, QuadTankGuard, TwinFlankTwin},
}

// Sniper Tier Two Tank
var Sniper = TierTwoTank{
	Name:           "Sniper",
	tierThreeTanks: []TierThreeTank{Assassin, Overseer, Hunter, Trapper},
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
	TierTwoTanks: []TierTwoTank{Twin, Sniper, MachineGun, FlankGuard},
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
	// var t searchTank
	router := httprouter.New()
	router.GET("/", Index)
	// build := []string{"Sprayer;", "net_replace_color 2 0x66FF00;", "net_replace_color 1 0xFFFF00;", "game_stats_build 654786547865478874658746547475656"}

	// fmt.Println("%v", baseTank.tierTwoTanks[1].name)

	// http.HandleFunc("/", searchTank)
	http.ListenAndServe(":8080", router)

}
