package main

import (
	"log"
	"os"

	"github.com/ozeer/go-api/cli/user"
	"github.com/urfave/cli/v2"
)

type Config struct {
	Commands []*cli.Command
}

func main() {
	// 创建配置并添加命令
	config := &Config{
		Commands: []*cli.Command{
			// 在这里添加其他命令，根据需要
			user.CliAnalysisAgeGroup,
		},
	}

	app := &cli.App{
		Commands: config.Commands,
		Usage:    "Go-Api的命令行工具",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
