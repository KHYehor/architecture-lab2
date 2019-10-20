package tablets

import (
	"encoding/json"
	"github.com/KHYehor/architecture-lab2/server/tools"
	"log"
	"net/http"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			//handleListChannels(store, rw)
		} else if r.Method == "POST" {
			handleGetData(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleSendData(r *http.Request, rw http.ResponseWriter, store *Store) {
	var state State
	if err := json.NewDecoder(r.Body).Decode(&state); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.setData(&state)
	if err != nil {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	} else {
		tools.WriteJsonOk(rw, "ok")
	}
}

func handleGetData(r *http.Request, rw http.ResponseWriter, store *Store) {
	var tablet Tablet
	if err := json.NewDecoder(r.Body).Decode(&tablet); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	res, err := store.getData(&tablet)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
