package models

import "time"

type Server struct {
	Id         int64
	Ip         string
	Port       int64
	Name       string
	Status     bool
	MaxPlayers int
	Online     int
	Players    []Online
	LastUpdate time.Time
}

func (s Server) GetLastUpdate() string {
	return s.LastUpdate.Format("02.01.2006 15:04:05")
}

type ServerOnline struct {
	Id         int64
	Name       string
	Status     bool
	MaxPlayers int
	Online     int
	OnlineMed  int
	Players    []OnlineMed
	LastUpdate time.Time
}

func (s ServerOnline) GetLastUpdate() string {
	return s.LastUpdate.Format("02.01.2006 15:04:05")
}

type Online struct {
	Id       int64
	Name     string
	Server   Server
	ServerId int64
}

type OnlineMed struct {
	Online
	Personal PersonalJson
}
