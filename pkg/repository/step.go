package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type StepRepository interface {
	FindByRecipe(*domain.Recipe) ([]*domain.Step, error)
}

type Step struct {
	db *sqlx.DB
}

func NewStepRepository(db *sqlx.DB) StepRepository {
	return &Step{db: db}
}

func (r *Step) FindByRecipe(recipe *domain.Recipe) ([]*domain.Step, error) {
	var steps []*domain.Step

	rows, err := r.db.Queryx(`
		SELECT uuid, index, instruction
		FROM steps
		WHERE recipe_id = $1
		ORDER BY index ASC`,
		recipe.ID,
	)
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
