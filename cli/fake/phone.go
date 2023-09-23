package fake

import (
	"context"
	"log"

	"github.com/ozeer/go-api/cli/base"
	"github.com/ozeer/go-api/cli/utils"
	"github.com/ozeer/go-api/global"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	CliGenFakePhone  = base.CreateCommand("GenFakePhone", "生成手机号Fake数据", GenFakePhone)
	CliGetFakePhone  = base.CreateCommand("GetFakePhone", "获取手机号Fake数据", GetFakePhone)
	cache_fake_phone = "phone_pool"
)

// run: go run enter.go GenFakePhone 10000
func GenFakePhone(args cli.Args) {
	count := args.Get(0)
	iCount := utils.StringToInt(count)

	for i := 0; i < iCount; i++ {
		phones := utils.FixNumberPhones(50)

		rows, err := global.REDIS.SAdd(context.Background(), cache_fake_phone, phones).Result()

		if err != nil {
			log.Println("错误:", err)
		}

		log.Println("写入数据条数:", rows)
	}
}

// run: go run enter.go GetFakePhone
func GetFakePhone(_ cli.Args) {
	phone, err := global.REDIS.SPop(context.Background(), cache_fake_phone).Result()

	if err != nil {
		global.LOG.Error("获取手机号fake数据失败", zap.Error(err))
	}

	log.Println("手机号:", phone)
}
