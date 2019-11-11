package models

import (
	"fmt"

	//pkg gocron helps to run a function in specific time (good for running cups)
	//more information at : https://github.com/jasonlvhit/gocron
	//"github.com/claudiu/gocron"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
//	Homes []*Home
/*Alleys       []*Alley
Towns        []*Town
AlleyMarkets []*AlleyMarket

TownMarkets  []*TownMarket*/
)

func init() {

	//fmt.Println("hi")
	//gocron.Start()
	//s := gocron.NewScheduler()
	/*gocron.Every(5).Seconds().Do(task)
	gocron.Every(10).Seconds().Do(vijay)
	gocron.Every(1).Monday().Do(task)*/
	//gocron.Every(1).Thursday().At("18:30").Do(doTownCup)
	//gocron.Every(1).Friday().At("18:30").Do(doAlleyCup)

	//<-s.Start()
}

type Alley struct {
	Homes       []*Home
	AlleyMarket *AlleyMarket
	Name        string
	AlleyNumber int
}

type Town struct {
	AlleysInTown []AlleyInTown
	TownMarket   *TownMarket
	Name         string
}

type AlleyInTown struct {
	HomeLevels  []int
	AlleyName   string
	AlleyNumber int
}

type GetVjAlley struct {
	MarketId string
	Imei     string
}

type VjAlleyResponse struct {
	IncreasingMoney int
	NextTime        string
}

type MarketData struct {
	RemainingTime int
}

type CupStruct struct {
	FirstRound  []MatchCap
	SecondRound []MatchCap
	FinalRound  MatchCap
}

type MatchCap struct {
	Id          int
	Imei1       string
	Imei2       string
	Name1       string
	Name2       string
	AvatarCode1 int
	AvatarCode2 int
	MatchId     int
	//winner: 1 means imei1 won. -1 means imei2 won
	Winner int
}

func SetHomes() {
	//	Homes = Homes[:0]
	//var Homes []*Home
	//	orm.NewOrm().QueryTable(new(Home)).All(&Homes)
	//fmt.Println(err)

}
func (home *Home) AddHome() {
	//fmt.Println("home added?")
	fmt.Println(orm.NewOrm().Insert(home))

	//	Homes = append(Homes, home)

	if home.Id%8 == 1 {
		alleyMrkt := &AlleyMarket{home.Id/8 + 1, 1, "name", time.Now().UTC().Format(time.UnixDate), 5, "state", ""}
		orm.NewOrm().Insert(alleyMrkt)
	}

	if home.Id%64 == 1 {
		townMrkt := &TownMarket{home.Id/64 + 1, 1, 1, "name", time.Now().Add(time.Hour * (-12)).UTC().Format(time.UnixDate), 5, "state"}
		orm.NewOrm().Insert(townMrkt)
	}

}

func GetAlley(alleyNumberStr string) (Alley, error) {
	alleyNum, _ := strconv.ParseInt(alleyNumberStr, 0, 64)
	alleyNumber := int(alleyNum)
	var homes []*Home
	startHome := alleyNumber * 8
	max := 8
	//	if len(Homes) == 0 {
	//		SetHomes()
	//	}
	size, _ := orm.NewOrm().QueryTable("home").Count()
	homesSize := int(size)
	if startHome+7 > homesSize-1 {
		max = homesSize - startHome
	}

	fmt.Println(max)

	for i := 0; i < max; i++ {
		//		home := Homes[startHome+i]
		home := getHome(startHome + 1 + i)
		homes = append(homes, home)
		fmt.Println(startHome + i)
	}

	alleyMarket := GetAlleyMarket(alleyNumber + 1)
	alley := Alley{homes, &alleyMarket, "no name", alleyNumber}
	fmt.Println(alley)
	return alley, nil
}

func getHome(id int) *Home {
	var home Home
	orm.NewOrm().QueryTable("Home").Filter("Id", id).RelatedSel().One(&home)
	return &home

}

func GetAlleyMarket(alleyNumber int) AlleyMarket {
	var alleyMarket AlleyMarket
	orm.NewOrm().QueryTable(new(AlleyMarket)).Filter("MarketId", alleyNumber).One(&alleyMarket)
	return alleyMarket
}

func GetTownMarket(townNumber int) TownMarket {
	var townMarket TownMarket
	orm.NewOrm().QueryTable(new(TownMarket)).Filter("MarketId", townNumber).One(&townMarket)
	return townMarket
}

func GetTown(townNumberStr string) (Town, error) {
	townNum, _ := strconv.ParseInt(townNumberStr, 0, 64)
	townNumber := int(townNum)
	var alleys []AlleyInTown
	var homeLevels []int
	startNumber := townNumber * 64
	max := 64
	size, _ := orm.NewOrm().QueryTable("home").Count()
	homesSize := int(size)
	if startNumber+63 > homesSize-1 {
		max = homesSize - startNumber
	}
	for i := 0; i < max; i++ {
		//		homeLevels = append(homeLevels, Homes[startNumber+i].Level)
		homeLevels = append(homeLevels, getHome(startNumber+i).Level)
	}

	for i := 0; i < len(homeLevels)/8; i++ {
		alley := AlleyInTown{homeLevels[i*8 : i*8+8], GetAlleyMarket(i + townNumber + 1).Name, i + townNumber}
		alleys = append(alleys, alley)
		fmt.Println(alleys)
	}
	if (len(homeLevels) % 8) != 0 {
		alley := AlleyInTown{homeLevels[8*(len(homeLevels)/8):],
			GetAlleyMarket(townNumber + len(homeLevels)/8 + 1).Name,
			townNumber + len(homeLevels)/8}
		alleys = append(alleys, alley)
	}

	var townMarket TownMarket
	orm.NewOrm().QueryTable(new(TownMarket)).Filter("MarketId", townNumber+1).One(&townMarket)
	town := Town{alleys, &townMarket, "name"}

	return town, nil
}

func AddVjAlley(vjAlley GetVjAlley) VjAlleyResponse {
	var alleyMarket AlleyMarket
	orm.NewOrm().QueryTable(new(AlleyMarket)).Filter("MarketId", vjAlley.MarketId).One(&alleyMarket)

	vjAdded := 0

	now, _ := time.Parse(time.UnixDate,
		time.Now().UTC().Format(time.UnixDate))

	alleyTime, _ := time.Parse(time.UnixDate,
		alleyMarket.Time)
	deltaTime := now.Sub(alleyTime)

	if deltaTime.Seconds() > 100 {
		vjAdded = getVjAmountAlley(alleyMarket.Level)
		_, finalVjAdded := AddVJ(vjAlley.Imei, vjAdded, true)
		vjAdded = finalVjAdded
		alleyMarket.Time = time.Now().UTC().Format(time.UnixDate)
		updateAlleyMarket(&alleyMarket)
	}

	nextTime := time.Now().Add(time.Hour * 12)

	return VjAlleyResponse{vjAdded, nextTime.String()}
}

func AddVjTown(vjTown GetVjAlley) VjAlleyResponse {
	var townMarket TownMarket
	orm.NewOrm().QueryTable(new(TownMarket)).Filter("MarketId", vjTown.MarketId).One(&townMarket)

	fmt.Println("town market: ", townMarket)

	vjAdded := 0

	now, _ := time.Parse(time.UnixDate,
		time.Now().UTC().Format(time.UnixDate))

	townTime, _ := time.Parse(time.UnixDate,
		townMarket.Time)
	deltaTime := now.Sub(townTime)

	fmt.Println("delta time", deltaTime)

	if deltaTime.Seconds() > 100 {
		vjAdded = getVjAmountTown(townMarket.Level)
		_, finalVjAdded := AddVJ(vjTown.Imei, vjAdded, true)
		vjAdded = finalVjAdded
		townMarket.Time = time.Now().UTC().Format(time.UnixDate)
		updateTownMarket(&townMarket)
	}

	nextTime := time.Now().Add(time.Hour * 12)

	return VjAlleyResponse{vjAdded, nextTime.String()}
}

func getVjAmountTown(marketLevel int) int {
	return allVjValues.LevelsValues[marketLevel-1].VjProduce * 4
}

func getVjAmountAlley(marketLevel int) int {
	return allVjValues.LevelsValues[marketLevel-1].VjProduce * 2
}

func GetMarketAlleyData(marketId string) (MarketData, error) {
	var alleyMarket AlleyMarket
	orm.NewOrm().QueryTable(new(AlleyMarket)).Filter("MarketId", marketId).One(&alleyMarket)

	alleyTime, _ := time.Parse(time.UnixDate,
		alleyMarket.Time)
	now, _ := time.Parse(time.UnixDate,
		time.Now().UTC().Format(time.UnixDate))
	remainingTime := int(alleyTime.Add(time.Hour * 12).Sub(now).Seconds())
	if remainingTime < 0 {
		remainingTime = 0
	}

	return MarketData{remainingTime}, nil
}

func GetMarketTownData(marketId string) (MarketData, error) {
	var townMarket TownMarket
	orm.NewOrm().QueryTable(new(TownMarket)).Filter("MarketId", marketId).One(&townMarket)

	townTime, _ := time.Parse(time.UnixDate,
		townMarket.Time)
	fmt.Println("town time: ", townTime)
	now, _ := time.Parse(time.UnixDate,
		time.Now().UTC().Format(time.UnixDate))
	fmt.Println("now time: ", now)
	remainingTime := int(townTime.Add(time.Hour * 12).Sub(now).Seconds())
	fmt.Println("remaining time: ", remainingTime)
	if remainingTime < 0 {
		remainingTime = 0
	}

	return MarketData{remainingTime}, nil
}

func updateAlleyMarket(market *AlleyMarket) {
	fmt.Println(market)
	_, err := orm.NewOrm().Update(market)
	if err != nil {
		//g.CatalogCacheDel("ids")
		fmt.Printf("save err... %s", err)
	}
}

func updateTownMarket(market *TownMarket) {
	fmt.Println(market)
	_, err := orm.NewOrm().Update(market)
	if err != nil {
		//g.CatalogCacheDel("ids")
		fmt.Printf("save err... %s", err)
	}
}

func (home *Home) SetLevel(level int) {
	home.Level = level
	_, err := orm.NewOrm().Update(home)
	fmt.Println(err)
}

func changeAvatarInMap(imei string, avatarCode int) {
	o := orm.NewOrm()
	var home Home
	o.QueryTable("home").Filter("Imei", imei).RelatedSel().One(&home)
	fmt.Println(home)
	home.AvatarCode = avatarCode
	_, err := orm.NewOrm().Update(&home)
	fmt.Println(err)
}

func setLevelForHome(imei string, level int) {
	o := orm.NewOrm()
	var home Home
	o.QueryTable("home").Filter("Imei", imei).RelatedSel().One(&home)
	fmt.Println(home)
	home.Level = level
	_, err := orm.NewOrm().Update(&home)
	fmt.Println(err)
}

func changeNameInMap(imei string, name string) {
	o := orm.NewOrm()
	var home Home
	o.QueryTable("home").Filter("Imei", imei).RelatedSel().One(&home)
	fmt.Println(home)
	home.Name = name
	_, err := orm.NewOrm().Update(&home)
	fmt.Println(err)
}

func StartAlleyCup() {
	//doing alley cup
	SetHomes()
	SetUsers()

	alleysNumber := 0
	size, _ := orm.NewOrm().QueryTable("home").Count()
	homesSize := int(size)
	if homesSize%8 == 0 {
		alleysNumber = homesSize / 8
	} else {
		alleysNumber = homesSize/8 + 1
	}
	for alleyNumber := 0; alleyNumber < alleysNumber; alleyNumber++ {
		alley, _ := GetAlley(strconv.Itoa(alleyNumber))
		fmt.Println("alley: ", alley)
		doAlleyCup(alley, alleyNumber)
	}
}

func doAlleyCup(alley Alley, alleyNumber int) {

	homeImeis := getAlleyRandomizedHomes(alley)

	fmt.Println("home numbers: ", len(homeImeis))

	var fRoundMatches []MatchCap
	for i := 0; i < 4; i++ {
		if homeImeis[i*2] == "0" {
			matchCup := buildMatchCupOffOff()
			fmt.Println("match cup: ", matchCup)
			fRoundMatches = append(fRoundMatches, matchCup)
			continue
		}
		if homeImeis[i*2+1] == "0" {
			matchCup := buildMatchCupOff(homeImeis[i*2])
			fmt.Println("match cup: ", matchCup)
			fRoundMatches = append(fRoundMatches, matchCup)
			continue
		}
		imeis := Imeis{homeImeis[i*2], homeImeis[i*2+1]}
		matchResult, winner := StartMatch(imeis, true)
		matchCup := buildMatchCup(matchResult, imeis, winner)
		fmt.Println("match cup: ", matchCup)
		fRoundMatches = append(fRoundMatches, matchCup)
	}

	fmt.Println("first round matches: ", fRoundMatches)

	//var secondRound SecondRound
	var sRoundMatches []MatchCap
	for i := 0; i < 2; i++ {
		var imei1, imei2 string
		if fRoundMatches[i*2].Winner == 1 {
			imei1 = fRoundMatches[i*2].Imei1
		} else {
			imei1 = fRoundMatches[i*2].Imei2
		}
		if fRoundMatches[i*2+1].Winner == 1 {
			imei2 = fRoundMatches[i*2+1].Imei1
		} else {
			imei2 = fRoundMatches[i*2+1].Imei2
		}

		if imei1 == "0" {
			matchCup := buildMatchCupOffOff()
			fmt.Println("match cup second: ", matchCup)
			sRoundMatches = append(sRoundMatches, matchCup)
			continue
		}
		if imei2 == "0" {
			matchCup := buildMatchCupOff(imei1)
			fmt.Println("match cup second: ", matchCup)
			sRoundMatches = append(sRoundMatches, matchCup)
			continue
		}
		imeis := Imeis{imei1, imei2}
		matchResult, winner := StartMatch(imeis, true)
		matchCup := buildMatchCup(matchResult, imeis, winner)
		fmt.Println("match cup second: ", matchCup)
		sRoundMatches = append(sRoundMatches, matchCup)
	}
	fmt.Println("second round matches: ", sRoundMatches)
	//secondRound.Matches = sRoundMatches

	var imei1, imei2 string
	if sRoundMatches[0].Winner == 1 {
		imei1 = sRoundMatches[0].Imei1
	} else {
		imei1 = sRoundMatches[0].Imei2
	}
	if sRoundMatches[1].Winner == 1 {
		imei2 = sRoundMatches[1].Imei1
	} else {
		imei2 = sRoundMatches[1].Imei2
	}

	if imei1 == "0" {
		matchCup := buildMatchCupOffOff()
		fmt.Println("match cup final: ", matchCup)
	}
	if imei2 == "0" {
		matchCup := buildMatchCupOff(imei1)
		fmt.Println("match cup final: ", matchCup)
	}
	imeis := Imeis{imei1, imei2}
	matchResult, winner := StartMatch(imeis, true)
	finalMatch := buildMatchCup(matchResult, imeis, winner)
	fmt.Println("match cup final: ", finalMatch)

	var cup = new(CupStruct)
	cup.FirstRound = fRoundMatches
	cup.SecondRound = sRoundMatches
	cup.FinalRound = finalMatch

	saveAlleyCup(convertCupStructToCup(cup, alleyNumber+1), alleyNumber+1)

	user1 := GetUserStruct(finalMatch.Imei1)
	user2 := GetUserStruct(finalMatch.Imei2)
	alleyMarket := GetAlleyMarket(alleyNumber + 1)
	if finalMatch.Winner == 1 && user1 != nil {
		user1.CupAlley += 1
		Update(convertUserStructToUser(user1))
		alleyMarket.Name = user1.Name
		alleyMarket.Imei = user1.Imei
	} else if finalMatch.Winner == -1 && user2 != nil {
		user2.CupAlley += 1
		Update(convertUserStructToUser(user2))
		alley.Name = finalMatch.Name2
		alleyMarket.Name = user2.Name
		alleyMarket.Imei = user2.Imei
	}
	updateAlleyMarket(&alleyMarket)
	fmt.Println("alley market: ", alleyMarket)
	fmt.Println("cup :", cup)
}

func buildMatchCupOffOff() MatchCap {
	var matchCup MatchCap
	matchCup.Imei1 = "0"
	matchCup.Imei2 = "0"
	matchCup.Name1 = ""
	matchCup.Name2 = ""
	matchCup.AvatarCode1 = 0
	matchCup.AvatarCode2 = 0
	matchCup.MatchId = 0
	matchCup.Winner = 0

	return matchCup
}

func buildMatchCupOff(imei string) MatchCap {
	user := GetUserStruct(imei)

	var matchCup MatchCap
	matchCup.Imei1 = imei
	matchCup.Imei2 = "0"
	matchCup.Name1 = user.Name
	matchCup.Name2 = ""
	matchCup.AvatarCode1 = user.AvatarCode
	matchCup.AvatarCode2 = 0
	matchCup.MatchId = 0
	matchCup.Winner = 1

	return matchCup
}

func buildMatchCup(matchResult MatchResult, imeis Imeis, winner int) MatchCap {
	user1 := GetUserStruct(imeis.MyImei)
	user2 := GetUserStruct(imeis.OpImei)

	var matchCup MatchCap
	matchCup.Imei1 = imeis.MyImei
	matchCup.Imei2 = imeis.OpImei
	matchCup.Name1 = user1.Name
	matchCup.Name2 = user2.Name
	matchCup.AvatarCode1 = user1.AvatarCode
	matchCup.AvatarCode2 = user2.AvatarCode
	matchCup.MatchId = matchResult.MatchEvent.MatchId
	matchCup.Winner = winner

	return matchCup
}

func getAlleyRandomizedHomes(alley Alley) []string {
	size := len(alley.Homes)
	//var homes []*Home
	homes := alley.Homes
	for i := 0; i < size; i++ {
		j := rand.Intn(size - 1)
		homes[i], homes[j] = homes[j], homes[i]
	}
	var imeis = make([]string, 8)

	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			if i < size {
				imeis[i/2] = homes[i].Imei
			} else {
				imeis[i/2] = "0"
			}
		} else {
			if i < size {
				imeis[i/2+4] = homes[i].Imei
			} else {
				imeis[i/2+4] = "0"
			}
		}
	}

	return imeis
}

func convertCupStructToCup(c *CupStruct, cupId int) *Cup {
	firstRound, _ := json.Marshal(c.FirstRound)
	secondRound, _ := json.Marshal(c.SecondRound)
	finalRound, _ := json.Marshal(c.FinalRound)

	cup := Cup{cupId, string(firstRound), string(secondRound), string(finalRound)}

	return &cup
}

func saveAlleyCup(this *Cup, alleyNum int) (int64, error) {
	var cup Cup
	orm.NewOrm().QueryTable(new(Cup)).Filter("CupId", alleyNum).One(&cup)

	var num int64
	var err error

	//fmt.Println("saving cup: ", this)

	if cup.CupId == alleyNum {
		num, err = orm.NewOrm().Update(this)
		if err != nil {
			//g.CatalogCacheDel("ids")
			fmt.Printf("update err... %s", err)
		}
	} else {
		num, err = orm.NewOrm().Insert(this)
		if err != nil {
			//g.CatalogCacheDel("ids")
			fmt.Printf("save err... %s", err)
		}
	}

	return num, err
}

func GetAlleyCup(alleyNumber string) CupStruct {
	id, _ := strconv.ParseInt(alleyNumber, 0, 64)
	var cup Cup
	orm.NewOrm().QueryTable(new(Cup)).Filter("CupId", id+1).One(&cup)
	fmt.Println(cup)
	return getCupStruct(cup)
}

func GetTownCup(townNumber string) CupStruct {
	id, _ := strconv.ParseInt(townNumber, 0, 64)
	var cup TownCup
	orm.NewOrm().QueryTable(new(TownCup)).Filter("CupId", id+1).One(&cup)
	return getTownCupStruct(cup)
}

func getTownCupStruct(cup TownCup) CupStruct {
	var firstRound []MatchCap
	var secondRound []MatchCap
	var finalRound MatchCap
	json.Unmarshal([]byte(cup.FinalRound), &finalRound)
	json.Unmarshal([]byte(cup.FirstRound), &firstRound)
	json.Unmarshal([]byte(cup.SecondRound), &secondRound)

	cupStr := CupStruct{firstRound, secondRound, finalRound}

	fmt.Println("cup struct: ", cupStr)

	return cupStr
}

func getCupStruct(cup Cup) CupStruct {
	var firstRound []MatchCap
	var secondRound []MatchCap
	var finalRound MatchCap
	json.Unmarshal([]byte(cup.FinalRound), &finalRound)
	json.Unmarshal([]byte(cup.FirstRound), &firstRound)
	json.Unmarshal([]byte(cup.SecondRound), &secondRound)

	cupStr := CupStruct{firstRound, secondRound, finalRound}

	fmt.Println("cup struct: ", cupStr)

	return cupStr
}

func StartTownCup() {
	//doing alley cup
	//	SetHomes()
	SetUsers()

	townsNumber := 0
	size, _ := orm.NewOrm().QueryTable("home").Count()
	homesSize := int(size)
	if homesSize%64 == 0 {
		townsNumber = homesSize / 64
	} else {
		townsNumber = homesSize/64 + 1
	}
	for townNumber := 0; townNumber < townsNumber; townNumber++ {
		town, _ := GetTown(strconv.Itoa(townNumber))
		fmt.Println("town: ", town)
		doTownCup(town, townNumber)
	}
}

func doTownCup(town Town, townNumber int) {
	homeImeis := getTownRandomizedHomes(town)

	fmt.Println("home imeis: ", homeImeis)

	var fRoundMatches []MatchCap
	for i := 0; i < 4; i++ {
		if homeImeis[i*2] == "0" {
			matchCup := buildMatchCupOffOff()
			fmt.Println("match cup: ", matchCup)
			fRoundMatches = append(fRoundMatches, matchCup)
			continue
		}
		if homeImeis[i*2+1] == "0" {
			matchCup := buildMatchCupOff(homeImeis[i*2])
			fmt.Println("match cup: ", matchCup)
			fRoundMatches = append(fRoundMatches, matchCup)
			continue
		}
		imeis := Imeis{homeImeis[i*2], homeImeis[i*2+1]}
		matchResult, winner := StartMatch(imeis, true)
		matchCup := buildMatchCup(matchResult, imeis, winner)
		fmt.Println("match cup: ", matchCup)
		fRoundMatches = append(fRoundMatches, matchCup)
	}

	fmt.Println("first round matches: ", fRoundMatches)

	//var secondRound SecondRound
	var sRoundMatches []MatchCap
	for i := 0; i < 2; i++ {
		var imei1, imei2 string
		if fRoundMatches[i*2].Winner == 1 {
			imei1 = fRoundMatches[i*2].Imei1
		} else {
			imei1 = fRoundMatches[i*2].Imei2
		}
		if fRoundMatches[i*2+1].Winner == 1 {
			imei2 = fRoundMatches[i*2+1].Imei1
		} else {
			imei2 = fRoundMatches[i*2+1].Imei2
		}

		if imei1 == "0" {
			matchCup := buildMatchCupOffOff()
			fmt.Println("match cup second: ", matchCup)
			sRoundMatches = append(sRoundMatches, matchCup)
			continue
		}
		if imei2 == "0" {
			matchCup := buildMatchCupOff(imei1)
			fmt.Println("match cup second: ", matchCup)
			sRoundMatches = append(sRoundMatches, matchCup)
			continue
		}
		imeis := Imeis{imei1, imei2}
		matchResult, winner := StartMatch(imeis, true)
		matchCup := buildMatchCup(matchResult, imeis, winner)
		fmt.Println("match cup second: ", matchCup)
		sRoundMatches = append(sRoundMatches, matchCup)
	}
	fmt.Println("second round matches: ", sRoundMatches)
	//secondRound.Matches = sRoundMatches

	var imei1, imei2 string

	if sRoundMatches[0].Winner == 1 {
		imei1 = sRoundMatches[0].Imei1
	} else {
		imei1 = sRoundMatches[0].Imei2
	}
	if sRoundMatches[1].Winner == 1 {
		imei2 = sRoundMatches[1].Imei1
	} else {
		imei2 = sRoundMatches[1].Imei2
	}
	var finalMatch MatchCap
	if imei1 == "0" {
		matchCup := buildMatchCupOffOff()
		finalMatch = matchCup
		fmt.Println("match cup final: ", matchCup)
	} else if imei2 == "0" {
		matchCup := buildMatchCupOff(imei1)
		finalMatch = matchCup
		fmt.Println("match cup final: ", matchCup)
	} else {
		imeis := Imeis{imei1, imei2}
		matchResult, winner := StartMatch(imeis, true)
		finalMatch = buildMatchCup(matchResult, imeis, winner)
		fmt.Println("match cup final: ", finalMatch)
	}
	var cup = new(CupStruct)
	cup.FirstRound = fRoundMatches
	cup.SecondRound = sRoundMatches
	cup.FinalRound = finalMatch

	saveTownCup(convertCupStructToTownCup(cup, townNumber+1), townNumber+1)

	user1 := GetUserStruct(finalMatch.Imei1)
	user2 := GetUserStruct(finalMatch.Imei1)
	townMarket := GetTownMarket(townNumber + 1)
	fmt.Println("twon market: ", townMarket)
	if finalMatch.Winner == 1 && user1 != nil {
		user1.CupTown += 1
		Update(convertUserStructToUser(user1))
		townMarket.Name = user1.Name
	} else if finalMatch.Winner == -1 && user2 != nil {
		user2.CupTown += 1
		Update(convertUserStructToUser(user2))
		townMarket.Name = user2.Name
	}
	updateTownMarket(&townMarket)
	fmt.Println("town market: ", townMarket)
	fmt.Println("cup :", cup)
}

func getTownRandomizedHomes(town Town) []string {
	size := len(town.AlleysInTown)

	alleys := town.AlleysInTown

	if size > 1 {
		for i := 0; i < size; i++ {
			j := rand.Intn(size - 1)
			alleys[i], alleys[j] = alleys[j], alleys[i]
		}
	}
	var imeis = make([]string, 8)

	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			if i < size {
				//fmt.Println("imei:", GetAlleyMarket(alleys[i].AlleyNumber+1).Imei)
				imeis[i/2] = GetAlleyMarket(alleys[i].AlleyNumber + 1).Imei
			} else {
				imeis[i/2] = "0"
			}
		} else {
			if i < size {
				imeis[i/2+4] = GetAlleyMarket(alleys[i].AlleyNumber).Imei
			} else {
				imeis[i/2+4] = "0"
			}
		}
	}

	return imeis
}

func convertCupStructToTownCup(c *CupStruct, cupId int) *TownCup {
	firstRound, _ := json.Marshal(c.FirstRound)
	secondRound, _ := json.Marshal(c.SecondRound)
	finalRound, _ := json.Marshal(c.FinalRound)

	cup := TownCup{cupId, string(firstRound), string(secondRound), string(finalRound)}

	return &cup
}

func saveTownCup(this *TownCup, townNum int) (int64, error) {
	var cup TownCup
	orm.NewOrm().QueryTable(new(TownCup)).Filter("CupId", townNum).One(&cup)

	var num int64
	var err error

	//fmt.Println("saving cup: ", this)

	if cup.CupId == townNum {
		num, err = orm.NewOrm().Update(this)
		if err != nil {
			//g.CatalogCacheDel("ids")
			fmt.Printf("update err... %s", err)
		}
	} else {
		num, err = orm.NewOrm().Insert(this)
		if err != nil {
			//g.CatalogCacheDel("ids")
			fmt.Printf("save err... %s", err)
		}
	}

	return num, err
}
