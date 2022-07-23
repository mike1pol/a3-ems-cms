package models

type User struct {
	Id         int64
	Name       string
	SteamId    string
	ProfileUrl string
	Avatar     string
	IsUser     bool
	IsAdmin    bool
}

type UserSession struct {
	Id     int64
	User   User
	UserId int64
	Token  string
}
