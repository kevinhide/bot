package handlers

import (
	"bot/constants"
	"bot/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Category : ""
func (h *Handler) Category(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	vars := mux.Vars(r)

	isParent := vars["isParent"]

	category, err := h.Service.Category(isParent)
	if err != nil {
		log.Println("err=>", err.Error())
		switch err.Error() {
		case constants.NOTFOUND:
			response.With404mV2(w, "Bank not found", platform)
			return
		case constants.INTERNALSERVERERROR:
			response.With500mV2(w, constants.INTERNALSERVERERROR, platform)
			return
		default:
			response.With500mV2(w, "Failed - "+err.Error(), platform)
			return
		}
	}
	m := make(map[string]interface{})
	m["category"] = category

	response.With200V2(w, "Category details", m, platform)
}
