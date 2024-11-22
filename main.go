package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	cmd "github.com/isaacgraper/spotfix.git/internal/cmd/cli"
)

func init() {
	fileName := "bot.log"
	filePath := filepath.Join("Z:\\", "RobôCOP", "Relatórios", "Execução", fileName)

	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("error creating directory: %v\n", err)
			return
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error opening log file: %v\n", err)
		return
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[INFO] ")

	log.Println("Bot is initializing")
}

func main() {
	if err := cmd.Run(); err != nil {
		log.Printf("[cli] error while trying to run bot: %v", err)
	}
}
