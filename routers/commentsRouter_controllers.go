package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetAlley",
            Router: `/alley/:alleynum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetAlleyCup",
            Router: `/getalleycup/:alleynum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetMarketAlleyData",
            Router: `/getmarketalleydata/:marketid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetMarketTownData",
            Router: `/getmarkettowndata/:marketid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetTownCup",
            Router: `/gettowncup/:townnum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetVjAlley",
            Router: `/getvjalley/:uaddress`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetVjTown",
            Router: `/getvjtown/:uaddress`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "MatchAlley",
            Router: `/matchalley/:marketid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "MatchTown",
            Router: `/matchtown/:marketid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MapController"],
        beego.ControllerComments{
            Method: "GetTown",
            Router: `/town/:townnum`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MatchController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:MatchController"],
        beego.ControllerComments{
            Method: "MapMatch",
            Router: `/match/:uaddress`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddWinColor",
            Router: `/addwincolor/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "BuyCards",
            Router: `/buycards/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChangeAvatar",
            Router: `/changeavatar/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChangeIsMail",
            Router: `/changeismail/:imei`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChangeName",
            Router: `/changename/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChangeTeamState",
            Router: `/changeteamstate/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChooseTeame",
            Router: `/chooseteam/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "CreateVj",
            Router: `/createvj/:uimei`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "ForooshCards",
            Router: `/forooshcards/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetPlayedMatch",
            Router: `/getmatch/:uid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOponent",
            Router: `/getoponent/:uimei`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetRankList",
            Router: `/getrank/:uimei`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetCardAmount",
            Router: `/match/:uimei`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "StartMatchByMap",
            Router: `/match/start/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetMessageHistory",
            Router: `/messagehistory/:messageid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "PromoteSafeBox",
            Router: `/promotesafebox/:uimei`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "RemoveLooseColor",
            Router: `/removeloosecolor/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendChatMsg",
            Router: `/sendchatmessage/:imei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendMessage",
            Router: `/sendmessage/:messageid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/cardGame/controllers:UserController"],
        beego.ControllerComments{
            Method: "StartMatch",
            Router: `/startmatch/:uimei`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
