package handler

import (
	"net/http"

	"github.com/kieranajp/langoustine/pkg/database/query"
)

func (h *BaseHandler) ListRecipesHandler(w http.ResponseWriter, r *http.Request) {
	q := query.GetAllRecipes{}
	recipes, err := q.New(h.db).Execute()
	if err != nil {
		h.failOnError(w, err, "error listing recipes")
		return
	}

	h.respondWithJSON(w, recipes)
}
