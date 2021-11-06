package handler

import (
	"encoding/json"
	"net/http"
)

func (h *BaseHandler) ListRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.recipeRepository.FindAll()
	if err != nil {
		h.failOnError(w, err, "error listing recipes")
		return
	}

	json, err := json.Marshal(recipes)
	if err != nil {
		h.failOnError(w, err, "error marshalling recipes")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
