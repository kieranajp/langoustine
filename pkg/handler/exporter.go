package handler

import (
	"fmt"

	"github.com/bmaupin/go-epub"
	"github.com/kieranajp/langoustine/pkg/domain"
)

type Exporter struct {
	*BaseHandler
}

func NewExporter(h *BaseHandler) *Exporter {
	return &Exporter{
		BaseHandler: h,
	}
}

func (e *Exporter) ExportToEpub(recipeUUID string) error {
	recipe, err := e.GetFullRecipe(recipeUUID)
	if err != nil {
		return err
	}

	book := e.generateEpub(recipe)
	err = book.Write(recipe.Name + ".epub")
	if err != nil {
		return err
	}

	return nil
}

func (e *Exporter) generateEpub(recipe *domain.Recipe) *epub.Epub {
	book := epub.NewEpub(recipe.Name)
	book.SetDescription(recipe.Description)

	titlePage := fmt.Sprintf(`<h1>%s</h1><p>%s</p>`, recipe.Name, recipe.Description)
	titlePage += fmt.Sprintf("<pre>%s, serves %d</pre>", recipe.Timing, recipe.Servings)

	ingredientsPage := "<h1>Ingredients</h1><ul>"
	for _, ingredient := range recipe.Ingredients {
		ingredientsPage += fmt.Sprintf("<li>%s</li>", ingredient.String())
	}
	ingredientsPage += "</ul>"

	stepsPage := "<h1>Steps</h1><ol>"
	for _, step := range recipe.Steps {
		stepsPage += fmt.Sprintf("<li>%s</li>", step.Instruction)
	}
	stepsPage += "</ol>"

	css, _ := book.AddCSS("https://raw.githubusercontent.com/mattharrison/epub-css-starter-kit/master/css/base.css", "")

	book.AddSection(titlePage, "Title", "", css)
	book.AddSection(ingredientsPage, "Ingredients", "", css)
	book.AddSection(stepsPage, "Steps", "", css)
	return book
}
