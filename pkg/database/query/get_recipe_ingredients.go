package query

import (
	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type GetRecipeIngredients struct {
	query string
	db    *sqlx.DB
}

func (q *GetRecipeIngredients) New(db *sqlx.DB) *GetRecipeIngredients {
	const query = `
		SELECT
			i.uuid,
			i.name,
			u.uuid "unit.uuid",
			u.name "unit.name",
			u.abbreviation "unit.abbreviation",
			ri.quantity "unit.quantity"
		FROM ingredients i
		INNER JOIN recipe_ingredient ri
			ON i.uuid = ri.ingredient_id
		INNER JOIN units u
			ON ri.unit_id = u.uuid
		WHERE ri.recipe_id = $1
	`
	return &GetRecipeIngredients{
		db:    db,
		query: query,
	}
}

func (q *GetRecipeIngredients) Execute(recipe *domain.Recipe) ([]*domain.Ingredient, error) {
	var ingredients []*domain.Ingredient

	rows, err := q.db.Queryx(q.query, recipe.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ingredient domain.Ingredient
		err = rows.StructScan(&ingredient)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, &ingredient)
	}

	return ingredients, nil
}
