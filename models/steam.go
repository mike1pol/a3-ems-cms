package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PlayerSummaries struct {
	SteamId     string `json:"steamid"`
	PersonaName string `json:"personaname"`
	ProfileUrl  string `json:"profileurl"`
	Avatar      string `json:"avatar"`
	RealName    string `json:"realname"`
}

func GetPlayerSummaries(steamId, apiKey string) (*PlayerSummaries, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", apiKey, steamId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	type Result struct {
		Response struct {
			Players []PlayerSummaries `json:"players"`
		} `json:"response"`
	}
	var data Result
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	if len(data.Response.Players) == 0 {
		return nil, errors.New("error parse data")
	}
	return &data.Response.Players[0], err
}
