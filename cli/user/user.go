package user

import (
	"fmt"

	"github.com/ozeer/go-api/cli/base"
)

var (
	CliAnalysisAgeGroup = base.CreateCommand("AgeGroup", "用户年龄分组统计脚本", TestAnalysisAgeGroup)
)

// 统计年龄分布脚本
func TestAnalysisAgeGroup(firstArg string) {
	// 在这里使用 UserCli，TaskCli 和 TestCli，如果需要的话
	fmt.Println(firstArg)
}
