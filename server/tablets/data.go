package tablets

import (
	"database/sql"
)

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// GET request
func (s *Store) getData(tablet *Tablet) (*Response, error) {
	// Init response
	var response Response
	response.Name = tablet.Name
	// Check if tablet exists
	row := s.Db.QueryRow(`SELECT id FROM TabletsList WHERE name=$1`, tablet.Name)
	err := row.Scan(&response.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		// Tablet not found
		return &response, nil
	}
	// Get all data with tablet
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

// Post request
func (s *Store) setData(sdata *SendData) error {
	// Check if tablet exists
	var id int64
	row := s.Db.QueryRow(`SELECT id FROM TabletsList WHERE name = $1`, sdata.Name)
	err := row.Scan(&id)
	// If it doesnt, so create it
	if err != sql.ErrNoRows && err != nil {
		return err
	} else if err == sql.ErrNoRows {
		_, err := s.Db.Exec(`INSERT INTO TabletsList ("name") VALUES ($1)`, sdata.Name)
		if err != nil {
			return err
		}
		row = s.Db.QueryRow(`SELECT id FROM TabletsList WHERE name = $1`, sdata.Name)
	}
	// Insert Data linked to table
	_, err = s.Db.Exec(`
  INSERT INTO TabletsState
   ("battery", "devicetime", "servertime", "currentvideo", "tabletid")
  VALUES
   ($1, $2, CURRENT_TIMESTAMP, $3, $4)`,
		sdata.Battery, sdata.DeviceTime, sdata.CurrentVideo, id)
	return err
}
