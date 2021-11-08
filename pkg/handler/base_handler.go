package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/repository"
	"github.com/rs/zerolog/log"
)

type BaseHandler struct {
	recipeRepository repository.RecipeRepository
}

func (h *BaseHandler) respondWithJSON(w http.ResponseWriter, payload interface{}) {
	wrapped := map[string]interface{}{
		"data": payload,
	}
	j, err := json.Marshal(wrapped)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (h *BaseHandler) failOnError(w http.ResponseWriter, err error, msg string) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + msg + `"}`))
		log.Error().Err(err).Msg(msg)
	}
}

func NewBaseHandler(db *sqlx.DB) *BaseHandler {
	return &BaseHandler{
		recipeRepository: repository.NewRecipeRepository(db),
	}
}
