package exporter

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kieranajp/langoustine/pkg/database/query"
	"github.com/kieranajp/langoustine/pkg/domain"
	"github.com/leotaku/mobi"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

type KindleExporter struct {
	db     *sqlx.DB
	writer io.Writer
}

func NewKindleExporter(db *sqlx.DB, writer io.Writer) *KindleExporter {
	return &KindleExporter{
		db:     db,
		writer: writer,
	}
}

func (e *KindleExporter) Export(recipeUUID string) error {
	q := query.GetFullRecipe{}
	recipe, err := q.New(e.db).Execute(recipeUUID)
	if err != nil {
		return errors.Wrap(err, "Failed to fetch recipe")
	}

	cover := NewCoverImage()
	err = cover.Generate(recipe)
	if err != nil {
		return errors.Wrap(err, "Failed to generate cover image")
	}

	book := e.generateMobi(recipe, cover)
	db := book.Realize()
	err = db.Write(e.writer)
	if err != nil {
		return errors.Wrap(err, "Failed to write mobi file")
	}

	return nil
}

func (e *KindleExporter) generateMobi(recipe *domain.Recipe, cover *CoverImage) mobi.Book {
	chapters := []mobi.Chapter{}
	chapters = append(chapters, e.generateDescriptionChapter(recipe))
	chapters = append(chapters, e.generateIngredientsChapter(recipe.Ingredients))
	chapters = append(chapters, e.generateStepsChapter(recipe.Steps))

	return mobi.Book{
		Title:       recipe.Name,
		Authors:     []string{"Kieran Patel"},
		CreatedDate: time.Now(),
		Language:    language.BritishEnglish,
		Chapters:    chapters,
		UniqueID:    rand.Uint32(),
		CoverImage:  cover.Image,
	}
}

func (e *KindleExporter) generateDescriptionChapter(recipe *domain.Recipe) mobi.Chapter {
	desc := fmt.Sprintf("<h1>%s</h1>", recipe.Name)
	desc += fmt.Sprintf("<p>Serves %d, takes %s</p>", recipe.Servings, recipe.Timing)
	desc += fmt.Sprintf("<p>%s</p>", recipe.Description)

	return mobi.Chapter{
		Title:  recipe.Name,
		Chunks: mobi.Chunks(recipe.Description),
	}
}

func (e *KindleExporter) generateIngredientsChapter(ingredients []*domain.Ingredient) mobi.Chapter {
	ingredientsHtml := "<h1>Ingredients</h1><ul>"
	for _, ingredient := range ingredients {
		ingredientsHtml += fmt.Sprintf("<li>%s</li>", ingredient.String())
	}
	ingredientsHtml += "</ul>"

	return mobi.Chapter{
		Title:  "Ingredients",
		Chunks: mobi.Chunks(ingredientsHtml),
	}
}

func (e *KindleExporter) generateStepsChapter(steps []*domain.Step) mobi.Chapter {
	stepsHtml := "<h1>Steps</h1><ol>"
	for _, step := range steps {
		stepsHtml += fmt.Sprintf("<li>%s</li>", step.Instruction)
	}
	stepsHtml += "</ol>"

	return mobi.Chapter{
		Title:  "Steps",
		Chunks: mobi.Chunks(stepsHtml),
	}
}
