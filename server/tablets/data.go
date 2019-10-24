package tablets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
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
	// IF tablet not exist -> crete it
	// timer data
	var id int64
	row := s.Db.QueryRow(`SELECT id FROM TabletsList WHERE name = $1`, sdata.Name)
	// fmt.Println(*row)
	if err := row.Scan(&id); err != nil {
		fmt.Scanln(id)
		fmt.Println("Her222")
		return err
	}
	if &id == nil {
		_, err := s.Db.Exec(`INSERT INTO TabletsList ("name") VALUES ($1)`, sdata.Name)
		if err != nil {
			return err
		}
		row = s.Db.QueryRow(`SELECT id FROM TabletsList WHERE name = $1`, sdata.Name)
	}
	if err := row.Scan(&id); err != nil {
		return err
	}
	_, err := s.Db.Exec(`
  INSERT INTO TabletsState
   ("battery", "devicetime", "servertime", "currentvideo", "tabletid")
  VALUES
   ($1, $2, CURRENT_TIMESTAMP, $3, $4)`,
		sdata.Battery, sdata.DeviceTime, sdata.CurrentVideo, id)
	// _, err := s.Db.Exec(`
	// 	INSERT INTO TabletsState
	// 		("battery", "devicetime", "servertime", "currentvideo", "tabletid")
	// 	VALUES
	// 		($1, $2, CURRENT_TIMESTAMP, $3, (SELECT id FROM TabletsList WHERE name = $4))`,
	// 	sdata.Battery, sdata.DeviceTime, sdata.CurrentVideo, sdata.Name)
	return err
}
