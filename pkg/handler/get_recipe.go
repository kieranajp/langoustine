package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *BaseHandler) GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeID := chi.URLParam(r, "recipeID")
	recipe, err := h.GetFullRecipe(recipeID)
	h.failOnError(w, err, "failed to get recipe")
	h.respondWithJSON(w, recipe)
}
