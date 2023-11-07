package global

import (
	"sync"

	"github.com/ozeer/go-api/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/sony/sonyflake"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/ozeer/go-api/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	DBList              map[string]*gorm.DB
	REDIS               *redis.Client
	CONFIG              config.Server
	VP                  *viper.Viper
	LOG                 *zap.Logger
	Timer               timer.Timer = timer.NewTimerTask()
	Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex

	SF *sonyflake.Sonyflake
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
