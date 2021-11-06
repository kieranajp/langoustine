package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/repository"
	"github.com/rs/zerolog/log"
)

type BaseHandler struct {
	recipeRepository repository.RecipeRepository
}

func (h *BaseHandler) failOnError(w http.ResponseWriter, err error, msg string) {
	if err != nil {
		log.Error().Err(err).Msg(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

func NewBaseHandler(db *sqlx.DB) *BaseHandler {
	return &BaseHandler{
		recipeRepository: repository.NewRecipeRepository(db),
	}
}
