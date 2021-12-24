package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kieranajp/langoustine/pkg/database/query"
)

func (h *BaseHandler) GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeID := chi.URLParam(r, "recipeID")
	q := query.GetFullRecipe{}
	recipe, err := q.New(h.db).Execute(recipeID)
	h.failOnError(w, err, "failed to get recipe")
	h.respondWithJSON(w, recipe)
}
