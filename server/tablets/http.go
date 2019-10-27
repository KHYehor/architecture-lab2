package tablets

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KHYehor/architecture-lab2/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/getData" {
			handleGetData(r, rw, store)
		} else if r.URL.String() == "/sendData" {
			handleSendData(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleSendData(r *http.Request, rw http.ResponseWriter, store *Store) {
	var sdata SendData
	if err := json.NewDecoder(r.Body).Decode(&sdata); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.setData(&sdata)
	if err != nil {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	} else {
		tools.WriteJsonOk(rw, "ok")
	}
	// Next request timeout
	time.Sleep(10 * time.Second)
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
