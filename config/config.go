package config

type Server struct {
	// 鉴权
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// 日志
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Redis
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	// 邮件
	Email Email `mapstructure:"email" json:"email" yaml:"email"`
	// 系统相关
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// 验证码
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	// excel
	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	// 定时器
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
