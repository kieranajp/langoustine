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

	ingredients, err := h.ingredientRepository.FindByRecipe(recipe)
	if err != nil {
		h.failOnError(w, err, "failed to find ingredients")
		return
	}
	recipe.Ingredients = ingredients

	steps, err := h.stepRepository.FindByRecipe(recipe)
	if err != nil {
		h.failOnError(w, err, "failed to find steps")
		return
	}
	recipe.Steps = steps

	h.respondWithJSON(w, recipe)
}
