// Collection management handlers.

package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Create a collection.
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	var col string
	if !Require(w, r, "col", &col) {
		return
	}
	if err := HttpDB.CreateCol(col); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

// Alll Return all collection names.
func All(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
    cols := HttpDB.GetAllCol()
	resp, err := json.Marshal(cols)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

// Rename a collection.
func Rename(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	var oldName, newName string
	if !Require(w, r, "old", &oldName) {
		return
	}
	if !Require(w, r, "new", &newName) {
		return
	}
	if err := HttpDB.RenameCol(oldName, newName); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
	}
}

// Drop a collection.
func Drop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "must-revalidate")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	var col string
	if !Require(w, r, "col", &col) {
		return
	}
	HttpDB.RemoveCol(col)
}

