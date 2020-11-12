package plants

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jn-lp/se-lab3/server/tools"
)

// HTTPHandlerFunc is Plants HTTP handler.
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of plants HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListPlants(store, rw)
		} else if r.Method == "POST" {
			handlePlantCreate(r, rw, store)
		} else if r.Method == "PATCH" {
			handlePlantUpdate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handlePlantCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	plant := Plant{SoilDataTimestamp: time.Now().Format(time.RFC3339)}
	if err := json.NewDecoder(r.Body).Decode(&plant); err != nil {
		log.Printf("plant input is empty")
	}
	plant, err := store.CreatePlant(plant.SoilMoistureLevel)
	if err == nil {
		tools.WriteJsonOk(rw, &plant)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListPlants(store *Store, rw http.ResponseWriter) {
	res, err := store.ListPlants()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handlePlantUpdate(r *http.Request, rw http.ResponseWriter, store *Store) {
	plant := Plant{SoilDataTimestamp: time.Now().Format(time.RFC3339)}
	if err := json.NewDecoder(r.Body).Decode(&plant); err != nil {
		log.Printf("Error decoding plant input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.UpdatePlant(plant.ID, plant.SoilMoistureLevel)
	if err == nil {
		tools.WriteJsonOk(rw, &plant)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
