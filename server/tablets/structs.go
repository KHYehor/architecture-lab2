package tablets

import "database/sql"

type Store struct {
	Db *sql.DB
}

type Tablet struct {
	Name string `json:"name"`
}

type State struct {
	Battery      string `json:"battery"`
	CurrentVideo string `json:"currentVideo"`
	DeviceTime   string `json:"deviceTime"`
	ServerTime   string `json:"serverTime"`
}

type SendData struct {
	Name         string `json:"name"`
	Battery      string `json:"battery"`
	DeviceTime   string `json:"deviceTime"`
	CurrentVideo string `json:"currentVideo"`
}

type Response struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	Telemetry []*State `json:"telemetry`
}
