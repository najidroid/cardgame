package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/najidroid/cardGame/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.SetUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	fmt.Println(uid)
	if uid != "" {
		user := models.GetUserStruct(uid)
		//		if err != nil {
		//			u.Data["json"] = err
		//		} else {
		u.Data["json"] = user
		//		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /createvj/:uimei [get]
func (u *UserController) CreateVj() {
	uimei := u.GetString(":uimei")
	if uimei != "" {
		user, err := models.CreateVj(uimei)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /match/:uimei [get]
func (u *UserController) GetCardAmount() {
	uimei := u.GetString(":uimei")
	if uimei != "" {
		user := models.GetCardsAmount(uimei)

		u.Data["json"] = user

	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /match/start/:uimei [post]
func (u *UserController) StartMatchByMap() {
	uimei := u.GetString(":uimei")
	var ob *[][]string
	//fmt.Print(" %s ", string(u.Ctx.Input.RequestBody))
	json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	fmt.Println(uimei)

}

/*
// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /delete/:uimei [post]
func (u *UserController) DeleteCards() {
	uimei := u.GetString(":uimei")
	var Ob models.MatchResultStruct

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	fmt.Println(uimei)
	result, ok := models.Match(Ob)

	if ok {
		u.Data["json"] = result
		fmt.Println(result)
	} else {
		u.Data["json"] = nil
	}
	u.ServeJson()
}*/

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /buycards/:uimei [post]
func (u *UserController) BuyCards() {
	uimei := u.GetString(":uimei")
	var Ob models.BuyCard

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(uimei)
	ok := models.BuyingCard(Ob, uimei)
	if ok {
		u.Data["json"] = Ob
		fmt.Println(Ob)
	} else {
		u.Data["json"] = nil
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /forooshcards/:uimei [post]
func (u *UserController) ForooshCards() {
	uimei := u.GetString(":uimei")
	var Ob models.ForooshCardStr

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(uimei)
	ok := models.ForooshCard(Ob, uimei)
	if ok {
		u.Data["json"] = Ob
		fmt.Println(Ob)
	} else {
		u.Data["json"] = nil
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /updatemissions/:uimei [post]
func (u *UserController) UpdateMissions() {
	uimei := u.GetString(":uimei")
	var Ob models.UserStruct
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	ok := models.UpdateMissions(uimei, Ob)
	if ok {
		u.Data["json"] = Ob
		fmt.Println(Ob)
	} else {
		u.Data["json"] = nil
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getmissionprize/:uimei [post]
func (u *UserController) GetMissionPrize() {
	uimei := u.GetString(":uimei")
	var Ob models.MissionPrize
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	ok := models.GetMissionPrize(uimei, Ob)
	if ok {
		u.Data["json"] = Ob
		fmt.Println(Ob)
	} else {
		u.Data["json"] = nil
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /addwincolor/:uimei [post]
func (u *UserController) AddWinColor() {
	uimei := u.GetString(":uimei")
	var Ob models.Color

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	_, color := models.AddColor(Ob, uimei)
	fmt.Println(color)
	//if ok {

	u.Data["json"] = color
	//fmt.Println(Ob)
	/*} else {
		u.Data["json"] = nil
	}*/
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /removeloosecolor/:uimei [post]
func (u *UserController) RemoveLooseColor() {
	uimei := u.GetString(":uimei")
	var Ob models.Color

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	//fmt.Println(Ob)
	_, color := models.RemoveColor(Ob, uimei)
	//if ok {
	u.Data["json"] = color
	fmt.Println(color)
	/*} else {
		u.Data["json"] = nil
	}*/
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /changeteamstate/:uimei [post]
func (u *UserController) ChangeTeamState() {
	uimei := u.GetString(":uimei")
	var Ob models.TeamState

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	ok := models.ChangeTeamState(Ob, uimei)
	if ok {
		u.Data["json"] = Ob
		fmt.Println(Ob)
	} else {
		u.Data["json"] = nil
	}
	u.ServeJSON()
}

// @Title Get
// @Description get random oponent
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getoponent/:uimei [get]
func (u *UserController) GetOponent() {
	//fmt.Println("hi")
	uimei := u.GetString(":uimei")
	oponent := models.GetOponent(uimei)

	u.Data["json"] = oponent

	u.ServeJSON()
}

// @Title Get
// @Description get random oponent
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getrank/:uimei [get]
func (u *UserController) GetRankList() {
	//fmt.Println("hi")
	uimei := u.GetString(":uimei")
	rankList := models.GetRankList(uimei)

	u.Data["json"] = rankList

	u.ServeJSON()
}

// @Title Get
// @Description get played match by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /getmatch/:uid [post]
func (u *UserController) GetPlayedMatch() {
	uid := u.GetString(":uid")

	var Ob models.Imeis

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	match := models.GetMatch(uid, Ob)

	u.Data["json"] = match

	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /startmatch/:uimei [post]
func (u *UserController) StartMatch() {
	//uimei := u.GetString(":uimei")
	var Ob models.Imeis

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	//fmt.Println(Ob)
	//fmt.Println(uimei)
	matchResult, _ := models.StartMatch(Ob, false)
	/*if ok {
		u.Data["json"] = Ob
		fmt.Println(Ob)
	} else {
		u.Data["json"] = nil
	}*/
	u.Data["json"] = matchResult
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /messagehistory/:messageid [post]
func (u *UserController) GetMessageHistory() {
	messageId := u.GetString(":messageid")

	var Ob models.Imeis

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	//fmt.Println(Ob)
	//fmt.Println(messageId)
	history := models.GetMessageHistory(messageId, Ob)
	u.Data["json"] = history
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /sendmessage/:messageid [post]
func (u *UserController) SendMessage() {
	messageId := u.GetString(":messageid")

	var Ob models.SentMessage

	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	//fmt.Println(Ob)
	//fmt.Println(messageId)
	history := models.SendMessage(messageId, Ob)
	u.Data["json"] = history
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /sendchatmessage/:imei [post]
func (u *UserController) SendChatMsg() {
	uimei := u.GetString(":imei")

	var Ob models.ChatMessage
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	models.SendChatMessage(Ob)
	fmt.Println(uimei)
	fmt.Println(Ob)
}

// @Title Get
// @Description gchange isMail to true
// @Param	imei		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :imei is empty
// @router /changeismail/:imei [get]
func (u *UserController) ChangeIsMail() {
	//fmt.Println("hi")
	imei := u.GetString(":imei")
	models.ChangeIsMail(imei, false)
	//u.Data["json"] = match

	//u.ServeJson()
}

// @Title Get
// @Description get random oponent
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /changeavatar/:uimei [post]
func (u *UserController) ChangeAvatar() {
	//fmt.Println("hi")
	uimei := u.GetString(":uimei")
	var Ob models.Avatar
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	avatar := models.ChangeAvatar(uimei, Ob)

	u.Data["json"] = avatar

	u.ServeJSON()
}

// @Title Get
// @Description get random oponent
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /changename/:uimei [post]
func (u *UserController) ChangeName() {
	//fmt.Println("hi")
	uimei := u.GetString(":uimei")
	var Ob models.Name
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	name := models.ChangeName(uimei, Ob)

	u.Data["json"] = name

	u.ServeJSON()
}

// @Title Get
// @Description get random oponent
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /chooseteam/:uimei [post]
func (u *UserController) ChooseTeame() {
	//fmt.Println("hi")
	uimei := u.GetString(":uimei")
	var Ob models.BuyCards
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	cards := models.ChoosingTeam(uimei, Ob)

	u.Data["json"] = cards

	u.ServeJSON()
}

// @Title Get
// @Description get random oponent
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /promotesafebox/:uimei [get]
func (u *UserController) PromoteSafeBox() {
	//fmt.Println("hi")
	uimei := u.GetString(":uimei")
	decreaseMoney := models.PromoteSafeBox(uimei)

	u.Data["json"] = decreaseMoney

	u.ServeJSON()
}
