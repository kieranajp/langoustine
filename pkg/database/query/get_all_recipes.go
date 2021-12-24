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
			r.uuid,
			r.name,
			r.description,
			p.url AS photo_url,
			r.timing,
			r.serving_size
		FROM recipes r
		LEFT JOIN photos p ON r.uuid = p.recipe_id
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
