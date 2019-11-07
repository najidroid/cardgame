package routers

//"github.com/astaxie/beego"

func init() {

	/*
		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"GetAll",
				"/",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"Get",
				"/:uid",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"CreateVj",
				"/createvj/:uimei",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"GetCardAmount",
				"/match/:uimei",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"StartMatchByMap",
				"/match/start/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"BuyCards",
				"/buycards/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"ForooshCards",
				"/forooshcards/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"AddWinColor",
				"/addwincolor/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"RemoveLooseColor",
				"/removeloosecolor/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"ChangeTeamState",
				"/changeteamstate/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"GetOponent",
				"/getoponent/:uimei",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"GetRankList",
				"/getrank/:uimei",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"GetPlayedMatch",
				"/getmatch/:uid",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"StartMatch",
				"/startmatch/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"GetMessageHistory",
				"/messagehistory/:messageid",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"SendMessage",
				"/sendmessage/:messageid",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"SendChatMsg",
				"/sendchatmessage/:imei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"ChangeIsMail",
				"/changeismail/:imei",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"ChangeAvatar",
				"/changeavatar/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"ChangeName",
				"/changename/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["cardGame/controllers:UserController"],
			beego.ControllerComments{
				"ChooseTeame",
				"/chooseteam/:uimei",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetAlley",
				"/alley/:alleynum",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetTown",
				"/town/:townnum",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetVjAlley",
				"/getvjalley/:uaddress",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetVjTown",
				"/getvjtown/:uaddress",
				[]string{"post"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetMarketAlleyData",
				"/getmarketalleydata/:marketid",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetMarketTownData",
				"/getmarkettowndata/:marketid",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"MatchAlley",
				"/matchalley/:marketid",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"MatchTown",
				"/matchtown/:marketid",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetAlleyCup",
				"/getalleycup/:alleynum",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MapController"],
			beego.ControllerComments{
				"GetTownCup",
				"/gettowncup/:townnum",
				[]string{"get"},
				nil})

		beego.GlobalControllerRouter["cardGame/controllers:MatchController"] = append(beego.GlobalControllerRouter["cardGame/controllers:MatchController"],
			beego.ControllerComments{
				"MapMatch",
				"/match/:uaddress",
				[]string{"get"},
				nil})
	*/
}
