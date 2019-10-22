package tablets

import (
	"database/sql"
)

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) getData(tablet *Tablet) (*Response, error) {
	var response Response
	response.Name = tablet.Name
	// First query
	row := s.Db.QueryRow(`SELECT id FROM TabletsList WHERE name=$1`, tablet.Name)
	if err := row.Scan(&response.Id); err != nil {
		return nil, err
	}
	// Second query
	rows, err := s.Db.Query(`
		SELECT 
			battery, currentVideo, devicetime, servertime 
		FROM tabletsState 
			WHERE tabletid=$1`, response.Id)
	if err != nil {
		return nil, err
	}
	// Proccesing states
	defer rows.Close()
	var res []*State
	for rows.Next() {
		var state State
		if err := rows.Scan(&state.Battery, &state.CurrentVideo, &state.DeviceTime, &state.ServerTime); err != nil {
			return nil, err
		}
		res = append(res, &state)
		response.Telemetry = append(response.Telemetry, &state)
	}
	if res == nil {
		res = make([]*State, 0)
	}
	// Sending response
	return &response, nil
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
