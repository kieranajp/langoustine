package query

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type GetFullRecipe struct {
	db *sqlx.DB
}

func (q *GetFullRecipe) New(db *sqlx.DB) *GetFullRecipe {
	return &GetFullRecipe{
		db: db,
	}
}

func (q *GetFullRecipe) Execute(recipeUUID string) (*domain.Recipe, error) {
	getRecipe := GetRecipe{}
	recipe, err := getRecipe.New(q.db).Execute(recipeUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to find recipe: %s", err)
	}

	getIngredients := GetRecipeIngredients{}
	recipe.Ingredients, err = getIngredients.New(q.db).Execute(recipe)
	if err != nil {
		return nil, fmt.Errorf("failed to find ingredients: %s", err)
	}

	getSteps := GetRecipeSteps{}
	recipe.Steps, err = getSteps.New(q.db).Execute(recipe)
	if err != nil {
		return nil, fmt.Errorf("failed to find recipe steps: %s", err)
	}
	return recipe, nil
}
