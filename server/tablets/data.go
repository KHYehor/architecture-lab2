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

type SendData struct {
	Name					string	`json:"name"`
	Battery				string	`json:"battery"`
	DeviceTime		string	`json:"deviceTime"`
	CurrentVideo	string	`json:"currentVideo"`
}

type Responce struct {
	Id						int64	 				`json:"id"`
	Name 					string 				`json:"name"`
	Telemetry			[]*State			`json:"telemetry`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) getData(tablet *Tablet) ([]*State, error) {
	// var ide int64
	// id, err := s.Db.Query(`SELECT id FROM TabletsList WHERE name=$1`, tablet.Name)
	// id.Next()
	// if err := id.Scan(&ide); err != nil {
	// 	return nil, err
	// }
	rows, err := s.Db.Query(`
		SELECT 
			battery, currentVideo, devicetime, servertime 
		FROM tabletsState 
			WHERE tabletid=(SELECT id FROM TabletsList WHERE name = $1)`, tablet.Name);
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//var responce *Responce
	var res []*State
	for rows.Next() {
		var state State
		if err := rows.Scan(&state.Battery, &state.CurrentVideo, &state.DeviceTime, &state.ServerTime); err != nil {
			return nil, err
		}
		res = append(res, &state)
		//responce.Telemetry = append(responce.Telemetry, &state)
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

func (s *Store) setData(sdata *SendData) error {
	_, err := s.Db.Exec(`
		INSERT INTO TabletsState 
			(battery, devicetime, servertime, currentvideo, tebletid) 
		VALUES 
			($1, $2, CURRENT_TIMESTAMP, $3, (SELECT id FROM TabletsList WHERE name = '$4'))`, 
		sdata.Battery, sdata.DeviceTime, sdata.CurrentVideo, sdata.Name)
	return err
}
