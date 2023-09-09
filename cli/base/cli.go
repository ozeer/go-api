package base

import "github.com/urfave/cli/v2"

// 创建CLI命令项
// name: 命令名
// usage: 使用方式说明
// f: 执行函数
func CreateCommand(name, usage string, f func(firstArg string)) *cli.Command {
	return &cli.Command{
		Name:  name,
		Usage: usage,
		Action: func(cCtx *cli.Context) error {
			f(cCtx.Args().First())
			return nil
		},
	}
}
