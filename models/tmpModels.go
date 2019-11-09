package models

type vjValues struct {
	LevelsValues []levelValues
	//1 to 2, 2 to 3, 3 to 4, 4 to 5,5 to 6,6 to 7, 7 to 8
	WinColorImprovementPrice   []int
	winColorImprovementDiamond []int
	DiamondPrice               int
}

type levelValues struct {
	//safe box capacity,improved safe box capacity
	// safe box capacity = base card price
	SafeBox   []int
	VjProduce int
	//2 to 1, 3 to 2, 4 to 3, 5 to 4
	ReduceLooseColorePrice         []int
	SafeBoxImprovementPriceVj      int
	SafeBoxImprovementPriceDiamond int
}

type xpJson struct {
	Levels []level
}

type level struct {
	Level int
	// win 1,2,3,4,5,6,7,8
	WinOthersPrize []int
}

type MessageStruct struct {
	MessageId    string
	Imei1        string
	Imei2        string
	MessagesBody [][]string
	//MessageBody : [["message","number of the owner"],...]
	// 1: lowerImei
}

type SentMessage struct {
	Imei1       string
	Imei2       string
	MessageBody string
	Owner       string
}

type MatchResult struct {
	MyState string
	MyCards []Card
	OpCards []Card
	//loss:
	//[["personNumber","if loosedGame","if colorAdded"],...]
	MyLoss            [][]string
	OponentLoss       [][]string
	MyImei            string
	MatchEvent        Event
	IncreasedXp       int
	MyIncreasedVj     int
	OpIncreasedVj     int
	OponentAvatarCode int
	MyAvatarCode      int
}

type MyMatchResultStruct struct {
	MyLoss     [][]int
	MyEarnedVJ int
	Event      Event
	Level      float32
}

type BuyCards struct {
	Cards []BuyCard
}

type BuyCard struct {
	DecreaseVJ      int
	BoughtCard      Card
	DecreaseDiamond int
}

type DecreaseMoney struct {
	DecreaseVJ      int
	DecreaseDiamond int
}

type ForooshCardStr struct {
	IncreaseVJ int
	Card       Card
}

type Rank struct {
	TopRank []RankMember
	MyRank  int
}

type RankMember struct {
	Name       string
	AvatarCode int
	CardAmount int
}

type RemainingTime struct {
	RemainingTime   int
	IncreasingMoney int
}

type Color struct {
	Color        string
	PersonNumber string
	Price        int
}

type TeamState struct {
	PersonNumber1 string
	PersonNumber2 string
}

type Imeis struct {
	MyImei string
	OpImei string
}

type ChatMessage struct {
	Message      string
	Imei         string
	PushMessage  string
	Subscription string
}

type Avatar struct {
	AvatarCode int
	ReduceVj   int
}

type Name struct {
	Name     string
	ReduceVj int
}
