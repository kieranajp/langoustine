package query

import (
	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type GetRecipeSteps struct {
	query string
	db    *sqlx.DB
}

func (q *GetRecipeSteps) New(db *sqlx.DB) *GetRecipeSteps {
	const query = `
		SELECT 
			uuid,
			index,
			instruction
		FROM steps
		WHERE recipe_id = $1
		ORDER BY index ASC
	`
	return &GetRecipeSteps{
		db:    db,
		query: query,
	}
}

func (q *GetRecipeSteps) Execute(recipe *domain.Recipe) ([]*domain.Step, error) {
	var steps []*domain.Step

	rows, err := q.db.Queryx(q.query, recipe.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var step domain.Step
		err = rows.StructScan(&step)
		if err != nil {
			return nil, err
		}
		steps = append(steps, &step)
	}

	return steps, nil
}
