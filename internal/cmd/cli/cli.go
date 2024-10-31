package cmd

import (
	"log"
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
					&cli.IntFlag{
						Name:  "max",
						Usage: "Máximo de resultados para processar",
						Value: 100,
					},
					&cli.StringFlag{
						Name:  "hour",
						Usage: "Defina um horário da inconsistência",
					},
					&cli.StringFlag{
						Name:  "category",
						Usage: "Defina o tipo da inconsistência",
					},
					&cli.BoolFlag{
						Name:  "filter",
						Usage: "Filtra pelo tipo de inconsistência primeiro",
						Value: false,
					},
					&cli.IntFlag{
						Name:  "batch",
						Usage: "Define o tamanho do lote para processamento",
						Value: 10,
					},
				},
				Action: func(ctx *cli.Context) error {
					config := config.Set(
						ctx.String("hour"),
						ctx.String("category"),
						ctx.Bool("filter"),
						ctx.Int("max"),
						ctx.Int("batch"),
					)

					if err := process.Execute(config); err != nil {
						log.Fatalf("error while trying to start the bot: %v", err)
					}
					return nil
				},
			},
		},
	}
	return app.Run(os.Args)
}
