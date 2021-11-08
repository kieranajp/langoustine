package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *BaseHandler) GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeID := chi.URLParam(r, "recipeID")
	recipe, err := h.recipeRepository.FindByUUID(recipeID)
	if err != nil {
		h.failOnError(w, err, "failed to find recipe")
		return
	}

	h.respondWithJSON(w, recipe)
}
