package domain

type Recipe struct {
	ID          string        `json:"id" db:"uuid"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
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

type UnitOfMeasure struct {
	ID       string `json:"id" db:"uuid"`
	Quantity int    `json:"quantity" db:"quantity"`
	Unit     string `json:"unit" db:"name"`
	Abbr     string `json:"abbr" db:"abbreviation"`
}
