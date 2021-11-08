package cmd

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kieranajp/langoustine/pkg/database"
	"github.com/kieranajp/langoustine/pkg/handler"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const uuidRegex = `[\w]{8}(-[\w]{4}){3}-[\w]{12}`

func Serve(c *cli.Context) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := database.Connect(c.String("db-dsn"))
	if err != nil {
		return err
	}
	defer db.Close()

	h := handler.NewBaseHandler(db)

	r.Get("/", handler.WelcomeHandler)
	r.Get("/recipes", h.ListRecipesHandler)
	r.Get(`/recipes/{recipeID:`+uuidRegex+`}`, h.GetRecipeHandler)

	log.Info().Str("listen-addr", c.String("listen-addr")).Msg("Starting server")
	return http.ListenAndServe(c.String("listen-addr"), r)
}
