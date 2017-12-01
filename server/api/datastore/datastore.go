package datastore

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zcrawl/zcrawl/server/api/helpers"
	"github.com/zcrawl/zcrawl/server/api/models"
)

// Router wraps chi.Router.
type Router struct {
	chi.Router
}

/*
// StoreRecord wraps a dynamic data structure.
type StoreRecord struct {
	ID             bson.ObjectId `bson:"_id" json:"_id"`
	JobID          bson.ObjectId `bson:"job_id" json:"job_id"`
	CollectionName string        `bson:"collection_name" json:"collection_name"`
	Data           map[string]interface{}
}
*/
func (r *Router) storeData(w http.ResponseWriter, req *http.Request) {
	rawRecord, err := ioutil.ReadAll(req.Body)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	record := models.StoreRecord{}
	err = json.Unmarshal(rawRecord, &record)
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = record.Save()
	if err != nil {
		helpers.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	recordJSON, _ := json.Marshal(&record)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(recordJSON)
}

// New initializes a new router.
func New() *Router {
	r := &Router{chi.NewRouter()}
	r.Post("/", r.storeData)
	return r
}
