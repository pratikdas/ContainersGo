package handlers

import (
	"net/http"
)

type Handler interface {
	handle(w http.ResponseWriter, r *http.Request)
}

/*type GetInventoryItemHandler GetInventoryItemHandler

func (handler GetInventoryItemHandler) handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}*/
