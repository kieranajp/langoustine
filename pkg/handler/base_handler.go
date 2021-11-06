package handler

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kieranajp/langoustine/pkg/repository"
)

type BaseHandler struct {
	recipeRepository repository.RecipeRepository
}

func NewBaseHandler(db *pgxpool.Pool) *BaseHandler {
	return &BaseHandler{
		recipeRepository: repository.NewRecipeRepository(db),
	}
}
