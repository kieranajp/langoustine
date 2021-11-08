package handler

import (
	"net/http"
)

func (h *BaseHandler) ListRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.recipeRepository.FindAll()
	if err != nil {
		h.failOnError(w, err, "error listing recipes")
		return
	}

	h.respondWithJSON(w, recipes)
}
