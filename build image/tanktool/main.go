package main

import (
	"html/template"
	"log"
	"net/http"
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

type TierFourTank struct {
	Name string
}
type TierThreeTank struct {
	Name          string
	tierFourTanks []TierFourTank
}

type TierTwoTank struct {
	Name           string
	tierThreeTanks []TierThreeTank
}

type TierOneTank struct {
	Name         string
	TierTwoTanks []TierTwoTank
}

var LandMine = TierFourTank{
	Name: "Land Mine",
}

var AutoSmasher = TierFourTank{
	Name: "Auto Smasher",
}

var Spike = TierFourTank{
	Name: "Spike",
}

var Triplet = TierFourTank{
	Name: "Triplet",
}

var PentaShot = TierFourTank{
	Name: "Penta Shot",
}

var SpreadShot = TierFourTank{
	Name: "Spread Shot",
}

var OctoTankQuad = TierFourTank{
	Name: "Octo Tank Quad",
}

var Auto5 = TierFourTank{
	Name: "Auto 5",
}

var TripleTwinTwin = TierFourTank{
	Name: "Triple Twin",
}

var BattleshipTwin = TierFourTank{
	Name: "Battleship Twin",
}

var Ranger = TierFourTank{
	Name: "Ranger",
}

var Stalker = TierFourTank{
	Name: "Stalker",
}

var Overlord = TierFourTank{
	Name: "Overlord",
}

var Necromancer = TierFourTank{
	Name: "Necromancer",
}

var Manager = TierFourTank{
	Name: "Manager",
}

var OverTrapperSeer = TierFourTank{
	Name: "Overtrapper Seer",
}

var BattleshipSniper = TierFourTank{
	Name: "Battleship Sniper",
}

var Factory = TierFourTank{
	Name: "Factory",
}

var Predator = TierFourTank{
	Name: "Predator",
}

var StreamlinerHunter = TierFourTank{
	Name: "Streamliner Hunter",
}

var Hibrid = TierFourTank{
	Name: "Hibrid",
}

var Annihilator = TierFourTank{
	Name: "Annihilator",
}

var Skimmer = TierFourTank{
	Name: "Skimmer",
}

var AutoGunnerGunner = TierFourTank{
	Name: "Auto Gunner Gunner",
}

var GunnerTrapperGunner = TierFourTank{
	Name: "Gunner Trapper Gunner",
}

var StreamlinerGunner = TierFourTank{
	Name: "Streamliner Gunner",
}

var TriTrapper = TierFourTank{
	Name: "Tri-Trapper",
}

var GunnerTrapperSniper = TierFourTank{
	Name: "Gunner Trapper Sniper",
}

var OverTrapperSniper = TierFourTank{
	Name: "Overtrapper Sniper",
}

var MegaTrapper = TierFourTank{
	Name: "Mega Trapper",
}

var AutoTrapper = TierFourTank{
	Name: "Auto Trapper",
}

var Booster = TierFourTank{
	Name: "Boster",
}

var Fighter = TierFourTank{
	Name: "Fighter",
}

var OctoTankFlank = TierFourTank{
	Name: "Octo Tank Flank",
}

var Auto5Quad = TierFourTank{
	Name: "Auto 5 Quad",
}

var TripleTwinFlank = TierFourTank{
	Name: "Triple Twin Flank",
}

var BattleshipFlank = TierFourTank{
	Name: "Battleship Flank",
}

var Auto5Auto = TierFourTank{
	Name: "Auto 5 Auto",
}

var AutoGunnerAuto = TierFourTank{
	Name: "Auto Gunner Auto",
}

var Smasher = TierThreeTank{
	Name:          "Smasher",
	tierFourTanks: []TierFourTank{LandMine, AutoSmasher, Spike},
}

var TripleShot = TierThreeTank{
	Name:          "Triple Shot",
	tierFourTanks: []TierFourTank{Triplet, PentaShot, SpreadShot},
}

var QuadTankTwin = TierThreeTank{
	Name:          "Quad Tank Twin",
	tierFourTanks: []TierFourTank{OctoTankQuad, Auto5},
}

var TwinFlankTwin = TierThreeTank{
	Name:          "Twin Flank Twin",
	tierFourTanks: []TierFourTank{TripleTwinTwin, BattleshipTwin},
}

var Assassin = TierThreeTank{
	Name:          "Assassin",
	tierFourTanks: []TierFourTank{Ranger, Stalker},
}

var Overseer = TierThreeTank{
	Name:          "Overseer",
	tierFourTanks: []TierFourTank{Overlord, Necromancer, Manager, OverTrapperSeer, BattleshipSniper, Factory},
}

var Hunter = TierThreeTank{
	Name:          "Hunter",
	tierFourTanks: []TierFourTank{TriTrapper, StreamlinerHunter},
}

var Trapper = TierThreeTank{
	Name:          "Trapper",
	tierFourTanks: []TierFourTank{TriTrapper, GunnerTrapperSniper, OverTrapperSeer, MegaTrapper, AutoTrapper},
}

var Destroyer = TierThreeTank{
	Name:          "Destroyer",
	tierFourTanks: []TierFourTank{Hibrid, Annihilator, Skimmer},
}

var Gunner = TierThreeTank{
	Name:          "Gunner",
	tierFourTanks: []TierFourTank{AutoGunnerGunner, GunnerTrapperGunner, StreamlinerGunner},
}

var Sprayer = TierThreeTank{
	Name:          "Spayer",
	tierFourTanks: []TierFourTank{AutoGunnerGunner, GunnerTrapperGunner, StreamlinerGunner},
}

var TriAngle = TierThreeTank{
	Name:          "Tri-Angle",
	tierFourTanks: []TierFourTank{Booster, Fighter},
}

var QuadTankGuard = TierThreeTank{
	Name:          "Quad Tank Guard",
	tierFourTanks: []TierFourTank{OctoTankFlank, Auto5Quad},
}

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

type searchTank int

func (s searchTank) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("%v", req.Form)
	// io.WriteString(w, "hello from a docker container tanktool")
	tpl.ExecuteTemplate(w, "index.gohtml", BaseTank)
	// if eerr != nil {
	// 	log.Fatalln(eerr)
	// }
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	var t searchTank
	// build := []string{"Sprayer;", "net_replace_color 2 0x66FF00;", "net_replace_color 1 0xFFFF00;", "game_stats_build 654786547865478874658746547475656"}

	// fmt.Println("%v", baseTank.tierTwoTanks[1].name)

	// http.HandleFunc("/", searchTank)
	http.ListenAndServe(":8080", t)

}
