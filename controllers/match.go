package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type MatchController struct {
	beego.Controller
}

// @Title Get
// @Description get town by uaddress
// @Param	uaddress		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Match
// @Failure 403 :uaddress is empty
// @router /match/:uaddress [get]
func (u *MatchController) MapMatch() {
	/*models.SetHomes()
	uaddress := u.GetString(":uaddress")
	if uaddress != "" {
		town, err := models.GetTown(uaddress)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = town
		}
	}
	u.ServeJson()*/
}
