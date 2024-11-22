package cmd

import (
	"fmt"
	"os"

	"github.com/isaacgraper/spotfix.git/internal/bot"
	"github.com/isaacgraper/spotfix.git/internal/common/config"
	"github.com/urfave/cli/v2"
)

func Run() error {
	process := bot.NewProcess()

	app := &cli.App{
		Name:  "Bot",
		Usage: "Automação de inconsistências de ponto - RPA",
		Commands: []*cli.Command{
			{
				Name:  "exec",
				Usage: "Executar o processamento do bot",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "notRegistered",
						Usage: "Processamento por filtro nas inconsistências como \"Não Registrado\"",
						Value: false,
					},
					&cli.BoolFlag{
						Name:  "workSchedule",
						Usage: "Processamento por filtro para Erros de Escala",
						Value: false,
					},
				},
				Action: func(ctx *cli.Context) error {
					config := config.Set(
						ctx.Bool("notRegistered"),
						ctx.Bool("workSchedule"),
					)

					if err := process.Execute(config); err != nil {
						return fmt.Errorf("[cli] error while trying to execute the bot: %w", err)
					}
					return nil
				},
			},
		},
	}
	return app.Run(os.Args)
}
