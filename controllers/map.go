package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/najidroid/cardGame/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type MapController struct {
	beego.Controller
}

/*
// @Title Get
// @Description get town by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /town/:uaddress [get]
func (u *MapController) GetTown() {
	models.SetHomes()
	uaddress := u.GetString(":uaddress")
	if uaddress != "" {
		town, err := models.GetTown(uaddress)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = town
		}
	}
	u.ServeJson()
}

// @Title Get
// @Description get town by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /city/:uaddress [get]
func (u *MapController) GetCity() {
	models.SetHomes()
	uaddress := u.GetString(":uaddress")
	if uaddress != "" {
		city, err := models.GetCity(uaddress)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = city
		}
	}
	u.ServeJson()
}*/

// @Title Get
// @Description get alley by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /alley/:alleynum [get]
func (u *MapController) GetAlley() {
	models.SetHomes()
	alleynum := u.GetString(":alleynum")
	if alleynum != "" {
		alley, err := models.GetAlley(alleynum)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = alley
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get alley by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /town/:townnum [get]
func (u *MapController) GetTown() {
	models.SetHomes()
	townNum := u.GetString(":townnum")
	if townNum != "" {
		alley, err := models.GetTown(townNum)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = alley
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /getvjalley/:uaddress [post]
func (u *MapController) GetVjAlley() {
	models.SetHomes()
	//uaddress := u.GetString(":uaddress")

	var Ob models.GetVjAlley
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	vjAdded := models.AddVjAlley(Ob)
	u.Data["json"] = vjAdded
	u.ServeJSON()
}

// @Title Get
// @Description get user by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /getvjtown/:uaddress [post]
func (u *MapController) GetVjTown() {
	models.SetHomes()
	//uaddress := u.GetString(":uaddress")

	var Ob models.GetVjAlley
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)

	vjAdded := models.AddVjTown(Ob)
	u.Data["json"] = vjAdded
	u.ServeJSON()
}

// @Title Get
// @Description get alley by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /getmarketalleydata/:marketid [get]
func (u *MapController) GetMarketAlleyData() {
	models.SetHomes()
	marketId := u.GetString(":marketid")
	if marketId != "" {
		marketData, err := models.GetMarketAlleyData(marketId)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = marketData
			fmt.Println(marketData)
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get town by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /getmarkettowndata/:marketid [get]
func (u *MapController) GetMarketTownData() {
	models.SetHomes()
	marketId := u.GetString(":marketid")
	if marketId != "" {
		marketData, err := models.GetMarketTownData(marketId)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = marketData
			fmt.Println(marketData)
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get alley by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /matchalley/:marketid [get]
func (u *MapController) MatchAlley() {
	models.StartAlleyCup()
}

// @Title Get
// @Description get alley by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /matchtown/:marketid [get]
func (u *MapController) MatchTown() {
	models.StartTownCup()
}

// @Title Get
// @Description get alley by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /getalleycup/:alleynum [get]
func (u *MapController) GetAlleyCup() {
	alleyNum := u.GetString(":alleynum")
	u.Data["json"] = models.GetAlleyCup(alleyNum)

	u.ServeJSON()
}

// @Title Get
// @Description get alley by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Map
// @Failure 403 :uaddress is empty
// @router /gettowncup/:townnum [get]
func (u *MapController) GetTownCup() {
	townNum := u.GetString(":townnum")
	u.Data["json"] = models.GetTownCup(townNum)

	u.ServeJSON()
}
