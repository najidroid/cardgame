package main

import (
	_ "github.com/najidroid/cardGame/docs"
	_ "github.com/najidroid/cardGame/routers"

	"github.com/astaxie/beego"

	"fmt"

	"github.com/astaxie/beego/orm"

	"log"
	"time"

	/*

		"flag"
		"log"
		"net/http"
		"text/template"
	*/)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.DefaultTimeLoc = time.UTC
	//	orm.RegisterDataBase("default", "mysql", "root:root@/cardgame?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "uongnacbhojt7sh8:335XIbI4t5NcdlUTC3la@tcp(bibrlater1qedqzgtdac-mysql.services.clever-cloud.com:3306)/bibrlater1qedqzgtdac?charset=utf8")

}

type jsonobject struct {
	Json string
}

func main() {

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := false

	// Error.
	err := orm.RunSyncdb(name, force, verbose)

	if err != nil {
		fmt.Println(err)

	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()

	/*this is for chat:
	flag.Parse()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err1 := http.ListenAndServe(*addr, nil)
	if err1 != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		fmt.Println(err)
		fmt.Println(msg)
	}
}

/*this is for chat
var addr = flag.String("addr", ":8080", "http service address")
var homeTemplate = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}
*/
