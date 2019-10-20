package tablets

import (
	"database/sql"
)

type Tablet struct {
	Name 	string `json:"name"`
}

type State struct {
	Battery 			string	`json:"battery"`
	CurrentVideo 	string	`json:"currentVideo"`
	DeviceTime 		string	`json:"deviceTime"`
	ServerTime 		string	`json:"serverTime"`
}

// type Responce struct {
// 	Id						int64	 				`json:"id"`
// 	Name 					string 				`json:"name"`
// 	Telemetry			[]*State			`json:"telemetry`
// }

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) getData(tablet *Tablet) ([]*State, error) {
	// id, err := s.Db.Query("SELECT id FROM TabletsList WHERE name='class1-tablet2'")
	rows, err := s.Db.Query("SELECT battery, currentVideo FROM tabletsState WHERE tabletid=(SELECT id FROM TabletsList WHERE name='" + tablet.Name + "')");
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*State
	for rows.Next() {
		var state State
		if err := rows.Scan(&state.Battery, &state.CurrentVideo, &state.DeviceTime, &state.ServerTime); err != nil {
			return nil, err
		}
		res = append(res, &state)
	}
	if res == nil {
		res = make([]*State, 0)
	}
	// var responce *Responce
	// id.Scan(&responce.Id)
	// responce.Name = "class1-tablet2"
	// responce.Telemetry = res;
	return res, nil
}

func (s *Store) setData(state *State) error {
	_, err := s.Db.Exec("")
	return err
}
