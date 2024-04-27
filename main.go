package main 

import (
	"database/sql"
	"net/http"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"bytes"
	"time"
)
var (
	db, _ = sql.Open("sqlite3", "stats.db")
	Config = LoadConfig()
)

func mainpage(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "discord.gg/nitroclaim")
}

func newclaim(w http.ResponseWriter, r *http.Request){
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	AddClaim(buf.String())
}

func Hook(w http.ResponseWriter, r *http.Request){
	enc, err := Encrypt(Config.PubWebhook)
	if err != nil {
		fmt.Fprint(w, "")
		return
	}

    fmt.Fprint(w, enc)
}

func DailyClaimsLoop() {
	for {
		DailyClaims = make(map[string]int)
		time.Sleep(time.Hour * 24)
	}
}


func main() {
	db.Exec(`
		CREATE TABLE IF NOT EXISTS claims (
			id TEXT PRIMARY KEY,
			count INT
		)
		
	`)
	db.Exec(`CREATE TABLE IF NOT EXISTS stats (
		id TEXT PRIMARY KEY,
		count INT
	)`)

	go StartBot()
	go DailyClaimsLoop()
	http.HandleFunc("/", mainpage)
	http.HandleFunc("/newclaim", newclaim)
	http.HandleFunc("/hook", Hook)

    log.Fatal(http.ListenAndServe(":1442", nil))
}