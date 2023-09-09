package user

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ozeer/go-api/cache"
	"github.com/ozeer/go-api/cli/base"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/model/user"
	"github.com/ozeer/go-api/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	CliAnalysisAgeGroup = base.CreateCommand("AgeGroup", "用户年龄分组统计脚本", AnalysisAgeGroup)
	CliDemo             = base.CreateCommand("Demo", "用户年龄分组统计脚本", Demo)
	Test                = base.CreateCommand("Test", "用户年龄分组统计脚本", OK)
)

const PAGE_LIMIT = 5000

// 统计年龄分布脚本
// run: go run enter.go AgeGroup 0 50000、go run enter.go AgeGroup 50000 100000
func AnalysisAgeGroup(args cli.Args) {
	begin := time.Now()

	startId := args.Get(0)
	endId := args.Get(1)

	if startId == "" || endId == "" {
		log.Fatalln("请输入要清洗数据的起始点和结束点")
	}

	fmt.Println(args.Get(0), args.Get(1))

	iStartId := utils.StringToInt(startId)
	iEndId := utils.StringToInt(endId)

	if iStartId < 0 || iEndId <= 0 || iStartId > iEndId {
		log.Fatalln("输入参数有误")
		return
	}
	// 查询
	var users []user.UserExtend

	iCursorId := iStartId
	iQueryCount := 0
	id := 1
	counter := 1

	for iCursorId < iEndId {
		if counter == iEndId {
			break
		}

		iStartId = iCursorId
		iCursorId = iStartId + PAGE_LIMIT

		if iCursorId > iEndId {
			iCursorId = iEndId
		}

		sql := "select * from user_extend where id > %d and id <= %d order by id"
		sqlData := fmt.Sprintf(sql, iStartId, iCursorId)
		global.DB.Debug().Raw(sqlData).Scan(&users)
		// global.DB.Raw(sqlData).Scan(&users)

		for _, user := range users {
			counter++
			iQueryCount++
			id++
			birthday := user.Birthday

			if birthday == "" {
				continue
			}

			if iQueryCount >= PAGE_LIMIT {
				fmt.Println(birthday)
				// 查累了，休息1s
				time.Sleep(1 * time.Second)
				iQueryCount = 0
			}
		}
	}

	fmt.Println("计数器: ", counter-1)
	fmt.Printf("\ncost: %dms\n", time.Since(begin).Milliseconds())
	return
}

func Demo(_ cli.Args) {
	begin := time.Now()

	// 查询
	var users []user.UserExtend
	// 获取所有用户
	var total int64
	err := global.DB.Model(&user.UserExtend{}).Find(&users).Count(&total).Error

	if err != nil {
		return
	}

	ageGroups := map[string]int{
		"under_20": 0,
		"20-30":    0,
		"30-40":    0,
		"40-50":    0,
		"above_50": 0,
	}

	today := time.Now()

	for _, user := range users {
		year := strings.Split(user.Birthday, "-")[0]
		intYear := utils.StringToInt(year)
		age := today.Year() - intYear

		switch {
		case age < 20:
			ageGroups["under_20"]++
		case age >= 20 && age < 30:
			ageGroups["20-30"]++
		case age >= 30 && age < 40:
			ageGroups["30-40"]++
		case age >= 40 && age < 50:
			ageGroups["40-50"]++
		default:
			ageGroups["above_50"]++
		}
	}

	for k := range cache.RdsData {
		cache.RdsData[k] = utils.IntToString(ageGroups[k])
	}

	fmt.Println("总人数:", total)
	fmt.Println(cache.RdsData)

	cache.SetAnalysisAgeGroup(cache.RdsData)

	fmt.Printf("\ncost: %dms\n", time.Since(begin).Milliseconds())
}

// go run enter.go Test
func OK(_ cli.Args) {
	data, err := global.REDIS.HGetAll(context.Background(), cache.USER_ANALYSIS_AGE).Result()

	if err != nil {
		global.LOG.Error("获取用户分心数据失败", zap.Error(err))
	}

	log.Println(data)
}
