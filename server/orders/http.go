package orders

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pavlovskyive/architecture-lab-3/server/tools"
)

// HTTPHandlerFunc handles orders HTTP
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of orders HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListOrders(store, rw)
		} else if r.Method == "POST" {
			handleCreateOrder(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListOrders(store *Store, rw http.ResponseWriter) {
	res, err := store.ListOrders()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJSONInternalError(rw)
		return
	}
	tools.WriteJSONOk(rw, res)
}

func handleCreateOrder(r *http.Request, rw http.ResponseWriter, store *Store) {
	var o Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		log.Printf("Error decoding menu item input: %s", err)
		tools.WriteJSONBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateOrder(o.MenuItemID, o.TableNumber)
	if err == nil {
		tools.WriteJSONOk(rw, &o)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJSONInternalError(rw)
	}
}
