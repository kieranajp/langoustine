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
			p.url AS photo_url,
			r.description,
			r.timing,
			r.serving_size
		FROM recipes r
		LEFT JOIN photos p ON r.uuid = p.recipe_id
		WHERE r.uuid = $1
		LIMIT 1
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
