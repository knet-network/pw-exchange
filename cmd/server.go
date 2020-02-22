package cmd

import (
	"github.com/aerogo/aero"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start PW Exchange Serveer",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	app := aero.New()
	configure(app).Run()
	return nil
}

func configure(app *aero.Application) *aero.Application {
	app.Get("/", func(ctx aero.Context) error {
		return ctx.String("Hello World")
	})

	return app
}
