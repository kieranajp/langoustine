package query

import (
	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type GetAllRecipes struct {
	query string
	db    *sqlx.DB
}

func (q *GetAllRecipes) New(db *sqlx.DB) *GetAllRecipes {
	const query = `
		SELECT
			uuid,
			name,
			description,
			timing,
			serving_size
		FROM recipes
	`
	return &GetAllRecipes{
		db:    db,
		query: query,
	}
}

func (q *GetAllRecipes) Execute() ([]*domain.Recipe, error) {
	var recipes []*domain.Recipe

	rows, err := q.db.Queryx(q.query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var recipe domain.Recipe
		err = rows.StructScan(&recipe)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}
