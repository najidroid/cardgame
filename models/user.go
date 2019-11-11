package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"io/ioutil"
)

var (
	xpPrize     xpJson
	allVjValues vjValues
)

func init() {
	readJsonFiles()
}

func readJsonFiles() {
	readVjJsonFiles()
	readXpJsonFiles()
}

func readVjJsonFiles() {
	file, e := ioutil.ReadFile("./vjvalues.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	json.Unmarshal(file, &allVjValues)
	//fmt.Printf("Vj Results: %v\n", allVjValues)
}

func readXpJsonFiles() {
	file, e := ioutil.ReadFile("./xp.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	json.Unmarshal(file, &xpPrize)
	//fmt.Printf("Results: %v\n", xpPrize)
}

func SetUsers() []*User {
	var data []*User

	orm.NewOrm().QueryTable(new(User)).All(&data)

	//	for i := 0; i < len(data); i++ {
	//		if data[i].HomeId == 2 {
	//			data[i].Level = 0
	//			orm.NewOrm().Update(data[i])
	//			break
	//		}
	//	}

	return data
}

func SortRankImeiListByImei(uimei string) {
	//i := Index(uimei, RankImeiList)
}

//func GetAllUsers() map[string]*UserStruct {
//	return UserList
//}

func GetUserStruct(uimei string) *UserStruct {
	o := orm.NewOrm()
	var usr User
	o.QueryTable("User").Filter("Imei", uimei).RelatedSel().One(&usr)
	if usr.Id != 0 {
		//		fmt.Println("user:", usr)
		return convertUserToUserStruct(&usr)
	}
	return convertUserToUserStruct(AddUser(uimei))
}

func AddUser(uimei string) (u *User) {
	size, _ := orm.NewOrm().QueryTable("user").Count()
	coor := float64(size/64) + float64((size%64)/8)/10 + float64(size%8)/100
	//id starts from zero, it causes some problem when saving in DB
	id := int(size) + 1

	cardsJson, _ := json.Marshal([]string{})
	messages := []string{} // an empty list
	//messages = append(messages, Message{[]string{"تبریک", "25/4/94", "به برنامه ی کارد گیم خوش اومدی\nلحظات خوبی رو برای شما آرزومندیم\nگروه نرم افزاری ناجی", "0", "url"}})
	messagesJson, _ := json.Marshal(messages)
	//var events []Event
	//events = append(events, Event{[]string{"وحید", "25/4/94", "matchId", "0"}})
	events := []string{}
	//	events = append(events, Event{[]string{"مری", "25/4/94", "matchId", "1"}})
	eventsJson, _ := json.Marshal(events)
	missions := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	missionsJson, _ := json.Marshal(missions)

	avatarCode := rand.Intn(30) + 1
	uu := User{id, uimei, "بدون نام" + strconv.Itoa(id), avatarCode, 20, 60, 0, "Sat Jul 16 15:36:16 UTC 2015", coor,
		string(cardsJson), string(messagesJson), string(eventsJson), 0, 0, getMaxXp(1), false, 0, 0, 0, string(missionsJson)}

	_, err := Save(&uu)
	fmt.Println(err)
	home := Home{id, uimei, "بدون نام" + strconv.Itoa(id), 0, avatarCode}
	home.AddHome()
	//	HomeList[uimei] = &home
	//	RankImeiList = append(RankImeiList, uimei)
	return &uu
}

func getMaxXp(level int) int {
	return 200 + 100*level
}

func convertUserToUserStruct(uu *User) (u *UserStruct) {
	var crds []Card
	var msgs [][]string
	var evnts []Event
	var missions []int
	json.Unmarshal([]byte(uu.Cards), &crds)
	json.Unmarshal([]byte(uu.MessageIds), &msgs)
	json.Unmarshal([]byte(uu.Events), &evnts)
	json.Unmarshal([]byte(uu.Missions), &missions)
	ur := UserStruct{uu.Id, uu.Imei, uu.Name, uu.AvatarCode, uu.Credit, uu.MaxVj, uu.Level,
		uu.CreateVjTime, uu.Coordinate, crds, msgs, evnts, uu.Flags, uu.Xp, uu.MaxXp,
		uu.IsMail, uu.Diamond, uu.CupAlley, uu.CupCity, missions}
	//	UserList[uu.Imei] = &ur
	return &ur
}

func Save(this *User) (int64, error) {
	num, err := orm.NewOrm().Insert(this)
	if err != nil {
		//g.CatalogCacheDel("ids")
		fmt.Printf("save err... %s", err)
	}

	return num, err
}

func Update(this *User) (int64, error) {
	num, err := orm.NewOrm().Update(this)
	if err != nil {
		//g.CatalogCacheDel("ids")
		fmt.Printf("save err... %s", err)
	}

	return num, err
}

func CreateVj(uimei string) (r RemainingTime, err error) {
	user := GetUserStruct(uimei)
	var remainTime RemainingTime
	remainingTime := getRemainingTime(user.CreateVjTime, user.Level)

	if remainingTime == 0 {
		vj := allVjValues.LevelsValues[user.Level].VjProduce
		_, addedVj := AddVJ(uimei, vj, false)
		if addedVj != 0 {
			user.CreateVjTime = time.Now().UTC().Format(time.UnixDate)
		} else {
			remainingTime = 5
		}
		Update(convertUserStructToUser(user))
		AddVJ(uimei, vj, true)
		//		orm.NewOrm().Update(convertUserStructToUser(user))
		//		fmt.Println(user.CreateVjTime)
		remainTime = RemainingTime{remainingTime, addedVj}
	} else {
		remainTime = RemainingTime{remainingTime, 0}
	}

	return remainTime, nil
}

func PromoteSafeBox(uimei string) DecreaseMoney {
	user := GetUserStruct(uimei)
	diamondPrice := allVjValues.LevelsValues[user.Level].SafeBoxImprovementPriceDiamond
	vjPrice := allVjValues.LevelsValues[user.Level].SafeBoxImprovementPriceVj
	if user.Diamond >= diamondPrice && user.Credit >= vjPrice {
		user.Credit -= vjPrice
		user.Diamond -= diamondPrice
		user.MaxVj = allVjValues.LevelsValues[user.Level].SafeBox[1]
		fmt.Println("user in promotion:", user)
		_, err := Update(convertUserStructToUser(user))
		if err == nil {
			return DecreaseMoney{vjPrice, diamondPrice}
		} else {
			return DecreaseMoney{0, 0}
		}
	} else {
		return DecreaseMoney{0, 0}
	}
}

func getRemainingTime(userTimeString string, level int) (delta int) {
	//	fmt.Println("100*(level+1)*(level+1)", 100*(level+1)*(level+1))
	userTime, _ := time.Parse(time.UnixDate, userTimeString)
	now, _ := time.Parse(time.UnixDate,
		time.Now().UTC().Format(time.UnixDate))
	//	fmt.Println("now(time):", now)
	//	fmt.Println("userTime(time):", userTime)
	//	fmt.Println("userTime(string):", userTimeString)
	d := now.Sub(userTime)
	dd := d.Seconds()
	//	fmt.Println("dd", dd)
	//	fmt.Println(d)
	//	fmt.Println(userTimeString)
	if int(dd) >= getVjProduceTime(level) {
		return 0
	}
	return getVjProduceTime(level) - int(dd)
}

func getVjProduceTime(level int) int {
	return 100*(level+1)*(level+1) + 500
}

func convertUserStructToUser(u *UserStruct) (user *User) {
	//fmt.Println(u)
	messagesJson, _ := json.Marshal(u.Messages)
	cardsJson, _ := json.Marshal(u.Cards)
	eventsJson, _ := json.Marshal(u.Event)
	missionsJson, _ := json.Marshal(u.Missions)
	uu := User{u.Id, u.Imei, u.Name, u.AvatarCode, u.Credit, u.MaxVj, u.Level, u.CreateVjTime,
		u.Coordinate, string(cardsJson), string(messagesJson), string(eventsJson),
		u.Flags, u.Xp, u.MaxXp, u.IsMail, u.Diamond, u.CupAlley, u.CupTown, string(missionsJson)}
	return &uu
}

func GetCardsAmount(uimei string) (u int) {
	var amount int = 0
	/*user, _ := UserList[uimei]
	for i := 0; i < int(user.Level+1)*10; i++ {
		amount += user.Cards[i]
	}*/
	return amount
}

func GetOponent(uimei string) (user *UserStruct) {
	SetUsers()
	//	u := GetUserStruct(uimei)

	cnt, _ := orm.NewOrm().QueryTable("user").Count()
	randomNum := rand.Intn(int(cnt)) + 1
	usr := User{Id: randomNum}
	orm.NewOrm().Read(&usr)
	//	fmt.Println("user:", usr)
	if usr.Imei != uimei && len(user.Cards) > 0 {
		return convertUserToUserStruct(&usr)
	} else {
		return GetOponent(uimei)
	}

}

func ChangeTeamState(teamState TeamState, uimei string) bool {
	user := GetUserStruct(uimei)

	cards := user.Cards
	for i := 0; i < len(cards); i++ {
		crd := cards[i]
		if crd.Number == teamState.PersonNumber1 {
			if crd.TeamState == "0" {
				cards[i].TeamState = "1"
			} else {
				cards[i].TeamState = "0"
			}
			break
		}
	}
	for i := 0; i < len(cards); i++ {
		crd := cards[i]
		if crd.Number == teamState.PersonNumber2 {
			if crd.TeamState == "0" {
				cards[i].TeamState = "1"
			} else {
				cards[i].TeamState = "0"
			}
			break
		}
	}
	user.Cards = cards
	//	fmt.Println(user)
	Update(convertUserStructToUser(user))
	return true
}

func AddColor(addingColor Color, uimei string) (bool, Color) {
	user := GetUserStruct(uimei)
	//fmt.Println(user)

	cards := user.Cards
	var winningSize int
	var myCard Card
	for i := 0; i < len(cards); i++ {
		crd := cards[i]
		if crd.Number == addingColor.PersonNumber {
			winningSize = len(crd.WinningColors)
			myCard = crd
			//			crd[1] = crd[1] + addingColor.Color
			break
		}
		//user.Cards[crd[0]] += int(crd[1])
	}
	//user.Cards = append(user.Cards, card)
	//	fmt.Println(user)

	price := allVjValues.WinColorImprovementPrice[winningSize-1] * (-1)
	if ok, _ := AddVJ(uimei, price, false); ok {
		//user.Credit += price
		user.Cards[indexOf(user.Cards, myCard)].WinningColors = myCard.WinningColors + addingColor.Color

		Update(convertUserStructToUser(user))
		//adding vj agein, because updating ruins above addVj
		AddVJ(uimei, price, true)
		return true, Color{addingColor.Color, addingColor.PersonNumber, price}
	} else {
		return false, Color{"", addingColor.PersonNumber, 0}
	}
}

func RemoveColor(color Color, uimei string) (bool, Color) {

	SetUsers()
	user := GetUserStruct(uimei)
	//fmt.Println(user)

	cards := user.Cards
	var card Card
	for i := 0; i < len(cards); i++ {
		crd := cards[i]
		if crd.Number == color.PersonNumber {
			card = crd
			break
		}
		//user.Cards[crd[0]] += int(crd[1])
	}
	i := strings.Index(card.LoosingColors, color.Color)

	price := allVjValues.LevelsValues[user.Level].ReduceLooseColorePrice[len(card.LoosingColors)-2] * (-1)
	//fmt.Println(allVjValues.LevelsValues[user.Level].ReduceLooseColorePrice)
	//fmt.Println("price: ", price)

	//user.Cards = append(user.Cards, card)
	//fmt.Println(user)
	if ok, _ := AddVJ(uimei, price, false); ok {
		//user.Credit += price
		card.LoosingColors = card.LoosingColors[:i] + card.LoosingColors[i+1:]
		Update(convertUserStructToUser(user))
		AddVJ(uimei, price, true)
		return true, Color{color.Color, color.PersonNumber, price}
	} else {
		return false, Color{"", color.PersonNumber, 0}
	}

}

func ChoosingTeam(imei string, cards BuyCards) BuyCards {
	for _, card := range cards.Cards {
		BuyingCard(card, imei)
	}
	fmt.Println("cards: ", cards)
	return cards
}

func BuyingCard(buyCard BuyCard, uimei string) bool {
	user := GetUserStruct(uimei)
	if user.Credit+buyCard.DecreaseVJ < 0 {
		return false
	}
	if user.Diamond+buyCard.DecreaseDiamond < 0 {
		return false
	}
	user.Credit += buyCard.DecreaseVJ
	user.Diamond += buyCard.DecreaseDiamond
	user.Cards = append(user.Cards, buyCard.BoughtCard)

	Update(convertUserStructToUser(user))
	return true
}

func AddVJ(uimei string, vjAmount int, isAdd bool) (bool, int) {
	user := GetUserStruct(uimei)

	maxVjAmount := user.MaxVj
	//fmt.Println(user)
	//	fmt.Println("user credit: ", user.Credit)
	//	fmt.Println("vj amount: ", vjAmount)
	if user.Credit+vjAmount >= 0 {
		increasedVj := 0
		if user.Credit+vjAmount <= maxVjAmount {
			increasedVj = vjAmount
		} else {
			increasedVj = maxVjAmount - user.Credit
		}
		if isAdd {
			if user.Credit+vjAmount <= maxVjAmount {
				user.Credit += vjAmount
			} else {
				user.Credit = maxVjAmount
			}
			Update(convertUserStructToUser(user))
		}
		//		fmt.Println("addVj : true")
		if increasedVj < 0 {
			increasedVj = 0
		}
		return true, increasedVj
	} else {
		fmt.Println("addVj : false")
		return false, 0
	}
}

func ForooshCard(fCard ForooshCardStr, uimei string) bool {
	SetUsers()
	//	fmt.Println("foroosh card:", fCard)
	user := GetUserStruct(uimei)
	myCards := user.Cards
	//	fmt.Println("my cads befor foroosh:", myCards)
	increaseVj := getCardValue(user.Level, len(fCard.Card.WinningColors), len(fCard.Card.LoosingColors)) * 8 / 10
	for i, card := range myCards {
		if fCard.Card.Number == card.Number {
			user.Cards = append(myCards[:i], myCards[i+1:]...)
			break
		}
	}
	Update(convertUserStructToUser(user))
	//	fmt.Println("my cads after foroosh:", myCards)
	AddVJ(uimei, increaseVj, true)
	return true
}

//func AddCard(card Card, uimei string) {
//	user := GetUserStruct(uimei)
//	//	fmt.Println(user)
//	/*for i := 0; i < len(cards); i++ {
//		crd := cards[i]
//		user.Cards[crd[0]] += int(crd[1])
//	}*/
//	user.Cards = append(user.Cards, card)
//	//	fmt.Println("adding card")
//	Update(convertUserStructToUser(user))
//}

func StartMatch(imeis Imeis, isCup bool) (MatchResult, int) {
	myUimei := imeis.MyImei
	opUimei := imeis.OpImei

	me := GetUserStruct(myUimei)
	op := GetUserStruct(opUimei)

	myCards := getCards(me.Cards)
	opCards := getCards(op.Cards)

	//	fmt.Println(myCards)
	//	fmt.Println(opCards)

	mFinalCards, oFinalCards := DoMAtch(myCards, opCards)

	//	fmt.Println(mCards)
	//	fmt.Println(oCards)

	mResult, winner := analyzeData(me, op, myCards, opCards, mFinalCards, oFinalCards, isCup)

	return mResult, winner
}

func analyzeData(me *UserStruct, op *UserStruct, myCards []Card, opCards []Card, mCards []Card,
	oCards []Card, isCup bool) (MatchResult, int) {

	//analyze winer and looser, increase "loosedGames" of loss cards, save the result
	var winner int

	//**declaring loss cards for oponent: mCards**
	for _, card := range myCards {
		i := indexOf(mCards, card)
		if i >= 0 {
			//deleting card from mCards
			mCards = append(mCards[:i], mCards[i+1:]...)
		}
	}
	//**declaring loss cards for me: oCards**
	for _, card := range opCards {
		i := indexOf(oCards, card)
		if i >= 0 {
			//deleting card from oCards
			oCards = append(oCards[:i], oCards[i+1:]...)
		}
	}

	//fmt.Println(mCards)
	//fmt.Println(oCards)

	// *winner = 1 means i won the match, -1 means oponent won, 0 means equal*
	//	fmt.Println(len(mCards))
	//	fmt.Println(len(oCards))
	if len(oCards) > len(mCards) {
		winner = -1
		if !isCup {
			op.Flags += 1
		}
	} else if len(oCards) < len(mCards) {
		winner = 1
		if !isCup {
			me.Flags += 1
		}
	} else {
		winner = 0
	}

	//adding xpPrize to winner
	myXp := 0
	opXp := 0
	if winner == 1 {
		myXp = getXpPrize(me.Level, op.Level)
		if !isCup {
			me.Xp += myXp
			//			fmt.Println("me.XP", me.Xp)
			//			fmt.Println("max Xp", getMaxXp(me.Level))
			if me.Xp > getMaxXp(me.Level) {
				me.Level++
				me.Xp = 0
				me.MaxXp = getMaxXp(me.Level)
				me.MaxVj = allVjValues.LevelsValues[me.Level].SafeBox[0]
				setLevelForHome(me.Imei, me.Level)
			}
		}
		//		fmt.Println("my increasing xp:", getXpPrize(me.Level, op.Level))
	} else if winner == -1 {
		//		opXp = getXpPrize(op.Level, me.Level)
		//		if !isCup {
		//			op.Xp += opXp
		//		}
		//		fmt.Println("op xp: ", getXpPrize(op.Level, me.Level))
	}

	//fmt.Println(winner)

	//	fmt.Println("oCards(my loss): ", oCards)
	//	fmt.Println("mCards(op loss): ", mCards)

	//****increasing "loosedGames" for loss cards***
	//loss: [["personNumber","if loosedGame","if colorAdded"],...]

	myLoss := [][]string{}
	opLoss := [][]string{}
	for _, card := range mCards {
		if !isCup {
			indexOfCard := indexOf(op.Cards, card)
			i, _ := strconv.ParseInt(card.LoosedGames, 0, 64)
			if i+1 < 10 {
				card.LoosedGames = strconv.Itoa(int(i) + 1)
				op.Cards[indexOfCard].LoosedGames = card.LoosedGames
				opLoss = append(opLoss, []string{card.Number, card.LoosedGames, "", card.WinningColors, card.LoosingColors})
			} else {
				nlc := addNewLooseColor(card)
				me.Cards[indexOfCard].LoosedGames = "0"
				me.Cards[indexOfCard].LoosingColors += nlc
				opLoss = append(opLoss, []string{card.Number, "0", nlc, card.WinningColors, card.LoosingColors})
			}
		} else {
			opLoss = append(opLoss, []string{card.Number, "0", "", card.WinningColors, card.LoosingColors})
		}
	}

	for _, card := range oCards {
		i, _ := strconv.ParseInt(card.LoosedGames, 0, 64)
		if !isCup {
			indexOfCard := indexOf(me.Cards, card)
			if i+1 < 10 {
				card.LoosedGames = strconv.Itoa(int(i) + 1)
				me.Cards[indexOfCard].LoosedGames = card.LoosedGames
				myLoss = append(myLoss, []string{card.Number, card.LoosedGames, "", card.WinningColors, card.LoosingColors})
			} else {
				nlc := addNewLooseColor(card)
				me.Cards[indexOfCard].LoosedGames = "0"
				me.Cards[indexOfCard].LoosingColors += nlc
				myLoss = append(myLoss, []string{card.Number, "0", nlc, card.WinningColors, card.LoosingColors})
			}
		} else {
			myLoss = append(myLoss, []string{card.Number, "0", "", card.WinningColors, card.LoosingColors})
		}
	}
	//	fmt.Println("my loss: ", myLoss)
	//	fmt.Println("op loss: ", opLoss)

	//fmt.Println(mCards)
	//fmt.Println(oCards)

	//	//	//createEvent(*me, *op, myCards, opCards, winner)
	//***creating new event for this match***
	myState := winner
	var opState int
	if myState == 0 {
		opState = 0
	} else if myState == 1 {
		opState = -1
	} else {
		opState = 1
	}

	if !isCup {
		if len(me.Event) > 4 {
			me.Event = append(me.Event[:4], me.Event[5:]...)
		}
		if len(op.Event) > 4 {
			op.Event = append(op.Event[:4], op.Event[5:]...)
		}
	}

	//calculating vjPrize
	myVjPrize := 0
	opVjPrize := 0
	if myState == 1 {
		myVjPrize = getVjPrize(me.Level, mCards)
		_, myVjPrize = AddVJ(me.Imei, myVjPrize, false)
	} else if myState == -1 {
		opVjPrize = getVjPrize(op.Level, oCards)
		_, opVjPrize = AddVJ(op.Imei, opVjPrize, false)
	}

	//	fmt.Println("myVjPrize:", myVjPrize)
	//	fmt.Println("opVjPrize:", opVjPrize)

	//updating data
	matchId := SaveMatchCard(myCards, opCards, myLoss, opLoss, me.Imei, op.Imei, myVjPrize, opVjPrize)

	mEvent := Event{Name: op.Name, Time: time.Now().UTC().Format(time.UnixDate), Imei: op.Imei, State: myState, MatchId: matchId}
	oEvent := Event{Name: me.Name, Time: time.Now().UTC().Format(time.UnixDate), Imei: me.Imei, State: opState, MatchId: matchId}
	//fmt.Println(mEvent)

	//cCards = append([][]string{card}, cCards...)
	if !isCup {
		me.Event = append([]Event{mEvent}, me.Event...)
		op.Event = append([]Event{oEvent}, op.Event...)
	}

	if !isCup {
		Update(convertUserStructToUser(me))
		Update(convertUserStructToUser(op))

		AddVJ(me.Imei, myVjPrize, true)
		AddVJ(op.Imei, opVjPrize, true)
	}

	matchResult := MatchResult{strconv.Itoa(myState), myCards, opCards, myLoss, opLoss,
		me.Imei, mEvent, myXp, myVjPrize, opVjPrize, op.AvatarCode, me.AvatarCode}

	//push match result to oponent
	//match-xpPrize-vjPrize-name-date-oponentImei-eventId-state
	if !isCup {
		push(op.Imei, "match-"+strconv.Itoa(opXp)+"-"+
			strconv.Itoa(opVjPrize)+"-"+me.Name+"-"+time.Now().UTC().Format(time.UnixDate)+
			"-"+me.Imei+"-"+strconv.Itoa(matchId)+"-"+strconv.Itoa(opState))

		ChangeIsMail(op.Imei, true)
	}

	return matchResult, winner
}

func getVjPrize(level int, cards []Card) int {
	vj := 0
	for _, card := range cards {
		//1:winning colors, 2: loosing colors
		vj += getCardValue(level, len(card.WinningColors), len(card.LoosingColors)) / 10
	}
	vj = vj/5 + 1
	return vj
}

func getCardValue(level int, winColorNumber int, looseColorNumber int) int {
	vj := 0
	vj += allVjValues.LevelsValues[level].SafeBox[1]
	//	fmt.Println("base value", allVjValues.LevelsValues[level].SafeBox[0])
	for i := 0; i < winColorNumber-1; i++ {
		vj += allVjValues.WinColorImprovementPrice[i]
		//		fmt.Println("win price", allVjValues.WinColorImprovementPrice[i])
	}
	for i := 0; i < looseColorNumber-1; i++ {
		vj -= allVjValues.LevelsValues[level].ReduceLooseColorePrice[i]
		//		fmt.Println("loose price", allVjValues.LevelsValues[level].ReduceLooseColorePrice[i])
	}
	if vj < 0 {
		vj = 0
	}
	//	fmt.Println("card value:", vj)
	return vj
}

func getXpPrize(level1 int, level2 int) int {
	if xpPrize.Levels == nil {
		readXpJsonFiles()
	}
	//	fmt.Println(level1)
	return xpPrize.Levels[level1].WinOthersPrize[level2]
}

func SaveMatchCard(myCards []Card, opCards []Card, myLoss [][]string,
	opLoss [][]string, myImei string, opImei string, myVjPrize int, opVjPrize int) int {
	mCards, _ := json.Marshal(myCards)
	oCards, _ := json.Marshal(opCards)
	mLoss, _ := json.Marshal(myLoss)
	oLoss, _ := json.Marshal(opLoss)

	matchCards := new(MatchCards)
	matchCards.MyCards = string(mCards)
	matchCards.OpCards = string(oCards)
	matchCards.MyLoss = string(mLoss)
	matchCards.OpLoss = string(oLoss)
	matchCards.MyImei = myImei
	matchCards.OpImei = opImei
	matchCards.MyVjPrize = myVjPrize
	matchCards.OpVjPrize = opVjPrize
	_, err := orm.NewOrm().Insert(matchCards)
	if err != nil {
		//g.CatalogCacheDel("ids")
		fmt.Printf("save err... %s", err)
	}

	//	id := strconv.Itoa(matchCards.Id)

	return matchCards.Id
}

func addNewLooseColor(card Card) string {
	loosColors := card.LoosingColors
	colors := "abcdefghij"
	var mColors string
	for i := 0; i < 10; i++ {
		if strings.Index(loosColors, string(colors[i])) < 0 {
			mColors = mColors + string(colors[i])
		}
	}
	r := rand.Intn(len(mColors))
	newLooseColor := string(mColors[r])
	card.LoosingColors = card.LoosingColors + newLooseColor
	return newLooseColor
}

func indexOf(cards []Card, card Card) int {
	for i, c := range cards {
		if c.Number == card.Number {
			return i
		}
	}
	return -1
}

func DoMAtch(myCards []Card, opCards []Card) ([]Card, []Card) {
	mCards := append([]Card{}, myCards...)
	oCards := append([]Card{}, opCards...)
	var cCards []Card
	var turn int

	turn = 0
	max := (len(mCards) + len(oCards)) * 2
	//fmt.Println("max: ", max)

	for counter := 0; counter < max; counter++ {
		mLength := len(mCards)
		oLength := len(oCards)
		cLength := len(cCards)

		//		fmt.Println("mLength:", mLength)
		//		fmt.Println("oLength:", oLength)
		//		fmt.Println("cLength:", cLength)

		if turn == 0 {
			if mLength > 0 {
				turn = 1
				mCards, oCards, cCards, turn = moveMyCards(mCards, oCards, cCards, turn)
			} else {
				if cLength > 0 {
					mCards, oCards, cCards, turn = sendCardsBack(mCards, oCards, cCards, turn)
				} else {
					return mCards, oCards
				}
			}
		} else {
			if oLength > 0 {
				turn = 0
				mCards, oCards, cCards, turn = moveOpCards(mCards, oCards, cCards, turn)
			} else {
				if cLength > 0 {
					mCards, oCards, cCards, turn = sendCardsBack(mCards, oCards, cCards, turn)
				} else {
					return mCards, oCards
				}
			}
		}
	}

	mCards, oCards, cCards, turn = sendCardsBack(mCards, oCards, cCards, turn)

	return mCards, oCards
}

func sendCardsBack(mCards []Card, oCards []Card, cCards []Card, trn int) ([]Card, []Card, []Card, int) {
	//	fmt.Println("sending cards back")
	if trn == 0 {
		mCards, oCards, cCards = sendCardsBackOpFirst(mCards, oCards, cCards)
	} else {
		mCards, oCards, cCards = sendCardsBackMeFirst(mCards, oCards, cCards)
	}
	return mCards, oCards, cCards, trn
}

func sendCardsBackMeFirst(mCards []Card, oCards []Card, cCards []Card) ([]Card, []Card, []Card) {
	turnn := 0
	for i := 0; i < len(cCards); i++ {
		card := cCards[i]
		if turnn == 0 {
			mCards = append(mCards, card)
			turnn = 1
		} else {
			oCards = append(oCards, card)
			turnn = 0
		}
	}
	cCards = cCards[:0]
	return mCards, oCards, cCards
}

func sendCardsBackOpFirst(mCards []Card, oCards []Card, cCards []Card) ([]Card, []Card, []Card) {
	turnm := 0
	for i := 0; i < len(cCards); i++ {
		card := cCards[i]
		if turnm == 0 {
			oCards = append(oCards, card)
			turnm = 1
		} else {
			mCards = append(mCards, card)
			turnm = 0
		}
	}
	cCards = cCards[:0]
	return mCards, oCards, cCards
}

func moveMyCards(mCards []Card, oCards []Card, cCards []Card, turn int) ([]Card, []Card, []Card, int) {

	card := mCards[0]
	//fmt.Println("my cards(befor moving): ", mCards)
	mCards = append(mCards[:0], mCards[1:]...)
	cCards = append([]Card{card}, cCards...)
	//fmt.Println("my cards(after moving): ", mCards)
	//fmt.Println("centre cards(after moving): ", cCards)
	if len(cCards) > 1 {
		if check(cCards[0].WinningColors, cCards[1].LoosingColors) {
			mCards, cCards = sendCardsToMe(mCards, cCards)
			turn = 0
			//fmt.Println("*** I won so I play again")
		}
	}
	return mCards, oCards, cCards, turn
}

func moveOpCards(mCards []Card, oCards []Card, cCards []Card, turn int) ([]Card, []Card, []Card, int) {

	card := oCards[0]
	//fmt.Println("op cards(befor moving): ", oCards)
	oCards = append(oCards[:0], oCards[1:]...)
	cCards = append([]Card{card}, cCards...)
	//fmt.Println("op cards(after moving): ", oCards)
	//fmt.Println("centre cards(after moving): ", cCards)
	if len(cCards) > 1 {
		if check(cCards[0].WinningColors, cCards[1].LoosingColors) {
			oCards, cCards = sendCardsToOp(oCards, cCards)
			turn = 1
			//fmt.Println("*** op won, so op plays again")
		}
	}
	return mCards, oCards, cCards, turn
}

func sendCardsToMe(mCards []Card, cCards []Card) ([]Card, []Card) {
	mCards = append(mCards, cCards...)
	cCards = cCards[:0]
	return mCards, cCards
}

func sendCardsToOp(oCards []Card, cCards []Card) ([]Card, []Card) {
	oCards = append(oCards, cCards...)
	cCards = cCards[:0]
	return oCards, cCards
}

func check(wColors string, lColors string) bool {
	for i := 0; i < len(wColors); i++ {
		c := string(wColors[i])
		//fmt.Println(strings.Index(lColors, c))

		if strings.Index(lColors, c) >= 0 {
			return true
		}
	}
	return false
}

func getCards(cardsStr []Card) []Card {
	var cards []Card

	for i := 0; i < len(cardsStr); i++ {
		crd := cardsStr[i]
		if crd.TeamState == "1" {
			//card := CardJson{crd[0], crd[1], crd[2]}
			cards = append(cards, crd)
		}
	}

	//mCards = append(mCards[:0], mCards[1:]...)
	//cCards = append([][]string{card}, cCards...)

	//fmt.Println("cards befor randomizing", cards)

	//randomizing cards
	size := len(cards)
	if size > 2 {
		for i := 0; i < size; i++ {
			j := rand.Intn(size - 1)
			cards[i], cards[j] = cards[j], cards[i]
		}
	}

	//fmt.Println("cards after randomizing", cards)

	return cards
}

func GetMatch(uid string, imeis Imeis) MatchResult {
	SetUsers()

	id, _ := strconv.ParseInt(uid, 0, 64)
	var m MatchCards
	orm.NewOrm().QueryTable(new(MatchCards)).Filter("Id", id).One(&m)

	me := GetUserStruct(imeis.MyImei)
	op := GetUserStruct(imeis.OpImei)

	oponentCards := []Card{}
	myCards := []Card{}
	opLoss := [][]string{}
	myLoss := [][]string{}

	json.Unmarshal([]byte(m.MyLoss), &myLoss)
	json.Unmarshal([]byte(m.OpLoss), &opLoss)
	json.Unmarshal([]byte(m.OpCards), &oponentCards)
	json.Unmarshal([]byte(m.MyCards), &myCards)

	if myLoss == nil {
		myLoss = [][]string{}
	}
	if opLoss == nil {
		opLoss = [][]string{}
	}

	mResult := MatchResult{"", myCards, oponentCards, myLoss, opLoss, m.MyImei, Event{}, 0, m.MyVjPrize,
		m.OpVjPrize, op.AvatarCode, me.AvatarCode}

	//	fmt.Println("match result: ", mResult)

	return mResult
}

func GetRankList(uimei string) Rank {
	//	var rankList []RankMember
	//	for i := 0; i < len(RankImeiList); i++ {
	//		u := GetUserStruct(RankImeiList[i])
	//		rankList = append(rankList, RankMember{u.Name, u.AvatarCode, GetCardsAmount(u.Imei)})
	//	}
	return Rank{nil, 0}
}

func GetMessageHistory(messageId string, imeis Imeis) MessageStruct {

	var m Message
	orm.NewOrm().QueryTable(new(Message)).Filter("MessageId", messageId).One(&m)

	//changing isRead state
	messageNum := 0
	user := GetUserStruct(imeis.MyImei)

	for num, message := range user.Messages {
		if message[0] == messageId {
			messageNum = num
			break
		}
	}
	if messageNum > 0 {
		user.Messages[messageNum][3] = "0"
	}

	//	fmt.Println(user)

	Update(convertUserStructToUser(user))

	return convertMessageToMessageStruct(m)

}

func convertMessageToMessageStruct(message Message) MessageStruct {
	var messageStr MessageStruct
	var body [][]string

	json.Unmarshal([]byte(message.MessagesBody), &body)

	messageStr = MessageStruct{message.MessageId, message.Imei1, message.Imei2, body}

	return messageStr
}

func SendMessage(messageId string, message SentMessage) SentMessage {
	var m Message
	orm.NewOrm().QueryTable(new(Message)).Filter("MessageId", messageId).One(&m)

	if m.MessageId == "" {
		body := [][]string{}
		body = append(body, []string{message.MessageBody, message.Owner})

		bodyJson, _ := json.Marshal(body)
		mm := Message{messageId, message.Imei1, message.Imei2, string(bodyJson)}
		//		fmt.Println(mm)
		orm.NewOrm().Insert(&mm)
		addMessageToUsers(messageId, message.Imei1, message.Imei2, message.Owner)
	} else {
		mMessage := convertMessageToMessageStruct(m)
		body := mMessage.MessagesBody
		//putting message in the 0-th place
		body = append([][]string{[]string{message.MessageBody, message.Owner}}, body...)

		bodyJson, _ := json.Marshal(body)
		mm := Message{messageId, message.Imei1, message.Imei2, string(bodyJson)}
		//		fmt.Println(mm)
		orm.NewOrm().Update(&mm)
		bringMessageToFront(messageId, message.Imei1, message.Imei2, message.Owner)
	}

	user1 := GetUserStruct(message.Imei1)
	user2 := GetUserStruct(message.Imei2)
	if message.Owner == "1" {
		push(message.Imei2, "msg-"+message.MessageBody+"-"+messageId+"-"+user1.Name)
		ChangeIsMail(message.Imei2, true)
	} else {
		push(message.Imei1, "msg-"+message.MessageBody+"-"+messageId+"-"+user2.Name)
		ChangeIsMail(message.Imei1, true)
	}

	return message
}

func bringMessageToFront(messageId string, imei1 string, imei2 string, owner string) {
	user1 := GetUserStruct(imei1)
	user2 := GetUserStruct(imei2)

	time := time.Now().UTC().Format(time.UnixDate)

	read1 := "1"
	read2 := "1"

	if owner == "1" {
		read1 = "0"
	} else {
		read2 = "0"
	}

	var messageNum1 int
	var messageNum2 int

	for num, message := range user1.Messages {
		if message[0] == messageId {
			messageNum1 = num
			break
		}
	}

	for num, message := range user2.Messages {
		if message[0] == messageId {
			messageNum2 = num
			break
		}
	}

	//deleting the old message
	user1.Messages = append(user1.Messages[:messageNum1], user1.Messages[messageNum1+1:]...)
	user2.Messages = append(user2.Messages[:messageNum2], user2.Messages[messageNum2+1:]...)

	//putting new message in the 0-th place
	user1.Messages = append([][]string{[]string{messageId, user2.Name, time, read1}}, user1.Messages...)
	user2.Messages = append([][]string{[]string{messageId, user1.Name, time, read2}}, user2.Messages...)

	Update(convertUserStructToUser(user1))
	Update(convertUserStructToUser(user2))
}

func addMessageToUsers(messageId string, imei1 string, imei2 string, owner string) {
	user1 := GetUserStruct(imei1)
	user2 := GetUserStruct(imei2)
	read1 := "1"
	read2 := "1"

	if owner == "1" {
		read1 = "0"
	} else {
		read2 = "0"
	}

	time := time.Now().UTC().Format(time.UnixDate)

	//putting new message in the 0-th place
	user1.Messages = append([][]string{[]string{messageId, user2.Name, time, read1}}, user1.Messages...)
	user2.Messages = append([][]string{[]string{messageId, user1.Name, time, read2}}, user2.Messages...)

	Update(convertUserStructToUser(user1))
	Update(convertUserStructToUser(user2))

}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func push(subscription string, message string) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://soldier.cloudmqtt.com:12466")
	opts.SetClientID("card game-server")
	opts.SetPassword("1ZG2oHDsATYa")
	opts.SetUsername("hnlbufmc")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		//		panic(token.Error())
		fmt.Println("connecting error : ", token.Error())
	}
	//	fmt.Println("subscription:", subscription, "\n message:", message)

	c.Publish(subscription, 0, false, message)
	//	c.Publish("update/id-2", 0, false, "message: johney")

	//time.Sleep(3 * time.Second)

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe(subscription); token.Wait() && token.Error() != nil {
		fmt.Println("unsubscribing error: ", token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}

func SendChatMessage(chatMessage ChatMessage) {
	push(chatMessage.Subscription, chatMessage.PushMessage)
}

func ChangeIsMail(imei string, isMail bool) {
	user := GetUserStruct(imei)
	user.IsMail = isMail
	Update(convertUserStructToUser(user))
}

func ChangeAvatar(imei string, avatar Avatar) Avatar {
	user := GetUserStruct(imei)
	newAvatarCode := avatar.AvatarCode
	if ok, _ := AddVJ(imei, 0, false); ok {
		user.AvatarCode = newAvatarCode
		Update(convertUserStructToUser(user))
		AddVJ(imei, -0, true)
		changeAvatarInMap(imei, newAvatarCode)
		return Avatar{newAvatarCode, 0}
	} else {
		return Avatar{user.AvatarCode, 0}
	}
}

func ChangeName(imei string, name Name) Name {
	SetUsers()
	user := GetUserStruct(imei)
	newName := name.Name
	if ok, _ := AddVJ(imei, -0, false); ok {
		user.Name = name.Name
		//		fmt.Println("name: ", name.Name)
		Update(convertUserStructToUser(user))
		AddVJ(imei, -0, true)
		changeNameInMap(imei, name.Name)
		return Name{newName, -0}
	} else {
		return Name{user.Name, 0}
	}
}

func (*User) TableEngine() string {
	return engine()
}
func (*Message) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}
