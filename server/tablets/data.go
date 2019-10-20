package tablets

import (
	"database/sql"
	"time"
)

type State struct {
	name 					string
	battery 			string
	devicetime 		string
	currentVideo 	string
}

type Tablet struct {
	name 					string
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) getData(tablet *Tablet) (string, error) {
	rows, err := s.Db.Query("SELECT battery, devicetime, timestamp, currentVideo FROM State WHERE tabletid='" + "(SELECT id FROM TabletsList WHERE name='" + tablet.name + "')")
	if err != nil {
		return "af", err
	}
	defer rows.Close()
	// преобразовать rows в res
	return "wfd", nil
}

func (s *Store) setData(state *State) error {
	_, err := s.Db.Exec("INSERT INTO State (battery, devicetime, timestamp, currentVideo, tabletid) VALUES (" + state.battery + ", " + state.devicetime + ", " + time.Now().Format(time.RFC3339) + ", " + state.currentVideo + ", " + "(SELECT id FROM TabletsList WHERE name='" + state.name + "')" + ")")
	return err
}
