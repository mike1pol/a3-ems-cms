package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go-pg/pg"
	"github.com/gobuffalo/envy"

	h "github.com/mike1pol/rms/helpers"
	. "github.com/mike1pol/rms/models"
)

func create(db *pg.DB, n string, s int64) (nw Online, err error) {
	nw = Online{
		Name:     n,
		ServerId: s,
	}
	errIns := db.Insert(&nw)
	if errIns != nil {
		log.Println("Error insert new online: ", errIns)
		err = errIns
		return
	}
	return
}

func cuData(db *pg.DB, server Server, info ServerInfo) {
	db.Exec("delete from onlines where server_id = ?", server.Id)
	t := time.Now().Add(time.Hour * 3)
	if info.Status {
		server.Name = info.Name
	}
	server.Online = info.Raw.Players
	server.MaxPlayers = info.MaxPlayers
	server.Status = info.Status
	server.LastUpdate = time.Now().Add(time.Hour * 3)
	errU := db.Update(&server)
	if errU != nil {
		log.Println("error update server:", errU)
	}
	var re = regexp.MustCompile(`(?m)(\w+)(\s\[.+\])`)
	for _, p := range info.Players {
		name := strings.TrimSpace(p.Name)
		if len(name) > 0 {
			name := re.ReplaceAllString(name, "$1")
			syncTime(db, name, t)
			create(db, name, server.Id)
		}
	}
}

type ServerInfoRaw struct {
	Players int `json:"numplayers"`
}

type ServerInfoPlayers struct {
	Name string `json:"name"`
}

type ServerInfo struct {
	Name       string              `json:"name"`
	Status     bool                `json:"status"`
	Raw        ServerInfoRaw       `json:"raw"`
	MaxPlayers int                 `json:"maxplayers"`
	Players    []ServerInfoPlayers `json:"players"`
}

func getOnline(server Server) (info ServerInfo, e error) {
	qs, err := envy.MustGet("QSERVER")
	if err != nil {
		return info, err
	}

	var url = fmt.Sprintf("%s?host=%s&port=%d", qs, server.Ip, server.Port)
	resp, e := http.Get(url)
	if e != nil {
		return info, e
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return info, err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return info, readErr
	}

	jsonErr := json.Unmarshal(body, &info)
	if jsonErr != nil {
		return info, jsonErr
	}
	return info, nil
}

func ServerUpdate(s Server) error {
	var db = h.GetDB()
	defer db.Close()
	info, err := getOnline(s)
	if err != nil {
		log.Printf("Error getting server info %d %s:%d - %s", s.Id, s.Ip, s.Port, err)
		return err
	} else {
		cuData(db, s, info)
	}

	return nil
}

func ServersUpdate() error {
	var db = h.GetDB()
	defer db.Close()
	var servers []Server
	errS := db.Model(&servers).Select()
	if errS != nil {
		log.Println("Error getting servers ", errS)
		return errS
	}

	var sId []int64

	for _, s := range servers {
		sId = append(sId, s.Id)
		info, err := getOnline(s)
		if err != nil {
			log.Printf("Error getting server info %d %s:%d - %s", s.Id, s.Ip, s.Port, err)
		} else {
			cuData(db, s, info)
		}
	}

	syncEMSDuty(db, sId)
	return nil
}
