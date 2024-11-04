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
						Name:  "processBatch",
						Usage: "Processamento por batch",
						Value: false,
					},
					&cli.IntFlag{
						Name:  "batch",
						Usage: "Define o tamanho do lote para processamento",
						Value: 10,
					},
					&cli.IntFlag{
						Name:  "max",
						Usage: "Máximo de resultados para processar",
						Value: 100,
					},
					&cli.StringFlag{
						Name:  "hour",
						Usage: "Define o horário da inconsistência",
					},
					&cli.StringFlag{
						Name:  "category",
						Usage: "Define o tipo da inconsistência",
					},
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
						ctx.Bool("processBatch"),
						ctx.Int("batch"),
						ctx.Int("max"),
						ctx.String("hour"),
						ctx.String("category"),
						ctx.Bool("notRegistered"),
						ctx.Bool("workSchedule"),
					)

					if err := process.Execute(config); err != nil {
						return fmt.Errorf("[cli] error while trying to start the bot: %w", err)
					}
					return nil
				},
			},
		},
	}
	return app.Run(os.Args)
}
