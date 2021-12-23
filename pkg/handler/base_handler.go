package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
	"github.com/kieranajp/langoustine/pkg/repository"
	"github.com/rs/zerolog/log"
)

type BaseHandler struct {
	recipeRepository     repository.RecipeRepository
	ingredientRepository repository.IngredientRepository
	stepRepository       repository.StepRepository
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

func (h *BaseHandler) GetFullRecipe(recipeUUID string) (*domain.Recipe, error) {
	recipe, err := h.recipeRepository.FindByUUID(recipeUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to find recipe: %s", err)
	}

	ingredients, err := h.ingredientRepository.FindByRecipe(recipe)
	if err != nil {
		return nil, fmt.Errorf("failed to find ingredients: %s", err)
	}
	recipe.Ingredients = ingredients

	steps, err := h.stepRepository.FindByRecipe(recipe)
	if err != nil {
		return nil, fmt.Errorf("failed to find recipe steps: %s", err)
	}
	recipe.Steps = steps
	return recipe, nil
}

func NewBaseHandler(db *sqlx.DB) *BaseHandler {
	return &BaseHandler{
		recipeRepository:     repository.NewRecipeRepository(db),
		ingredientRepository: repository.NewIngredientRepository(db),
		stepRepository:       repository.NewStepRepository(db),
	}
}
