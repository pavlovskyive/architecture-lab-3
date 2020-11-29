package menu

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pavlovskyive/architecture-lab-3/server/tools"
)

// HTTPHandlerFunc handles menu HTTP
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of menu HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListMenu(store, rw)
		} else if r.Method == "POST" {
			handleMenuCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListMenu(store *Store, rw http.ResponseWriter) {
	res, err := store.ListMenuItems()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJSONInternalError(rw)
		return
	}
	tools.WriteJSONOk(rw, res)
}

func handleMenuCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var i Item
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		log.Printf("Error decoding menu item input: %s", err)
		tools.WriteJSONBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateMenuItem(i.Name, i.Price)
	if err == nil {
		tools.WriteJSONOk(rw, &i)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJSONInternalError(rw)
	}
}
