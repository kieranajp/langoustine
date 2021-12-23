package domain

import "fmt"

type Recipe struct {
	ID          string        `json:"id" db:"uuid"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Servings    int           `json:"servings" db:"serving_size"`
	Timing      string        `json:"timing"`
	Steps       []*Step       `json:"steps,omitempty"`
	Ingredients []*Ingredient `json:"ingredients,omitempty"`
}

type Step struct {
	ID          string `json:"id" db:"uuid"`
	Index       int    `json:"index"`
	Instruction string `json:"instruction"`
}

type Ingredient struct {
	ID   string        `json:"id" db:"uuid"`
	Name string        `json:"name"`
	Unit UnitOfMeasure `json:"unit" db:"unit"`
}

func (i *Ingredient) String() string {
	return fmt.Sprintf("%g %s %s", i.Unit.Quantity, i.Unit.Abbr, i.Name)
}

type UnitOfMeasure struct {
	ID       string  `json:"id" db:"uuid"`
	Quantity float64 `json:"quantity" db:"quantity"`
	Unit     string  `json:"unit" db:"name"`
	Abbr     string  `json:"abbr" db:"abbreviation"`
}
