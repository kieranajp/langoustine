package query

import (
	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type GetRecipe struct {
	query string
	db    *sqlx.DB
}

func (q *GetRecipe) New(db *sqlx.DB) *GetRecipe {
	const query = `
		SELECT
			r.uuid,
			r.name,
			r.description,
			r.timing,
			r.serving_size
		FROM recipes r
		WHERE r.uuid = $1
	`
	return &GetRecipe{
		db:    db,
		query: query,
	}
}

func (q *GetRecipe) Execute(recipeUUID string) (*domain.Recipe, error) {
	var recipe domain.Recipe

	err := q.db.Get(&recipe, q.query, recipeUUID)
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}
