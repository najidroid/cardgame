package models

import (
	//	"encoding/json"
	//	"fmt"
	//	"math/rand"
	//	"strconv"
	//	"strings"
	//	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	//	"os"

	//	mqtt "github.com/eclipse/paho.mqtt.golang"

	//	"io/ioutil"
)

type User struct {
	Id           int `orm:"pk"`
	Imei         string
	Name         string
	AvatarCode   int
	Credit       int
	MaxVj        int
	Level        int
	CreateVjTime string
	Coordinate   float64
	Cards        string `orm:"size(1500)"` // تو بدترین حالت ۲۱ کارت ذخیره میکنه
	//Title string `orm:"size(600)"`
	//string: default size is carchar(255)
	MessageIds string `orm:"size(600)"`
	Events     string `orm:"size(600)"`
	Flags      int
	Xp         int
	MaxXp      int
	IsMail     bool
	Diamond    int
	CupAlley   int
	CupCity    int
	Missions   string `orm:"size(1000)"`
}

type UserStruct struct {
	Id           int
	Imei         string
	Name         string
	AvatarCode   int
	Credit       int
	MaxVj        int
	Level        int
	CreateVjTime string
	Coordinate   float64
	//cards : [[number,winning colors,loosing colors,loosed games,team state],...]
	Cards []Card
	//messages : [["messageId","name","date","is message read"],...] , the messageId relates to the id if the message in DB
	Messages [][]string
	Event    []Event
	Flags    int
	Xp       int
	MaxXp    int
	IsMail   bool
	Diamond  int
	CupAlley int
	CupTown  int
	Missions []int
}

type Card struct {
	Id            int
	Number        string
	WinningColors string
	LoosingColors string
	LoosedGames   string
	TeamState     string
}

type Event struct {
	//	Body []string
	Id      int
	Name    string
	Time    string
	Imei    string
	State   int
	MatchId int
}

type MatchCards struct {
	Id        int
	MyCards   string `orm:"size(1000)"`
	OpCards   string `orm:"size(1000)"`
	MyLoss    string
	OpLoss    string
	MyImei    string
	OpImei    string
	MyVjPrize int
	OpVjPrize int
}

type Message struct {
	MessageId    string `orm:"pk"`
	Imei1        string
	Imei2        string
	MessagesBody string `orm:"size(1000)"`
}

type TownCup struct {
	CupId       int    `orm:"pk"`
	FirstRound  string `orm:"size(1000)"`
	SecondRound string `orm:"size(600)"`
	FinalRound  string
}

type Cup struct {
	CupId       int    `orm:"pk"`
	FirstRound  string `orm:"size(1000)"`
	SecondRound string `orm:"size(600)"`
	FinalRound  string
}

type TownMarket struct {
	MarketId    int `orm:"pk"`
	Level       int
	TownLevel   int
	Name        string
	Time        string
	MoneyAmount int
	State       string
}

type AlleyMarket struct {
	MarketId    int `orm:"pk"`
	Level       int
	Name        string
	Time        string
	MoneyAmount int
	State       string
	Imei        string
}

type Home struct {
	HomeId     int `orm:"pk"`
	Imei       string
	Name       string `orm:"size(300)"`
	Level      int
	AvatarCode int
}

func init() {
	orm.RegisterModel(new(User), new(MatchCards), new(Message), new(Home), new(AlleyMarket),
		new(TownMarket), new(Cup), new(TownCup))

	//orm.RegisterModel(new(MatchResult))
}
