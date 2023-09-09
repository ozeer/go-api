package user

import (
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/model/common/response"
	"github.com/ozeer/go-api/model/user"
	"github.com/ozeer/go-api/model/user/request"
	"github.com/ozeer/go-api/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserApi struct{}

// 注册接口
func (u *UserApi) Register(c *gin.Context) {
	// 数据绑定
	var registerInfo request.RegisterInfo

	err := c.ShouldBindJSON(&registerInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 数据校验
	err = utils.Verify(registerInfo, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 邮箱可选参数，如果填了就校验
	if registerInfo.Email != "" {
		if err = utils.Verify(registerInfo.Email, utils.EmailVerify); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}

	// 出生日期可选
	if registerInfo.Birthday != "" {
		if err = utils.Verify(registerInfo.Birthday, utils.BirthdayVerify); err != nil {
			response.FailWithMessage("出生日期格式不对", c)
			return
		}
	}

	// 性别检查
	if registerInfo.Sex != user.UN_KNOW && registerInfo.Sex != user.MALE && registerInfo.Sex != user.FEMALE {
		response.FailWithMessage("性别设置有误", c)
		return
	}

	// 数据初始化
	var userInfo user.User
	var userExtendInfo user.UserExtend

	uid, err := global.SF.NextID()
	if err != nil {
		global.LOG.Error("snowflake err:", zap.Error(err))
		return
	}

	// 用户信息
	if registerInfo.UserName == "" {
		registerInfo.UserName = utils.GenNewNick(uid)
	}

	ip, err := utils.GetIP(c.Request)
	if err != nil {
		global.LOG.Error("获取IP地址失败!", zap.Error(err))
	}

	userInfo.Uid = uid                                          // 封装生成uid的方法
	userInfo.UserName = registerInfo.UserName                   // 生成随机用户名
	userInfo.Phone = registerInfo.Phone                         // 加密，防止用户敏感信息泄漏
	userInfo.Password = utils.BcryptHash(registerInfo.Password) // 加密后写入
	userInfo.Email = registerInfo.Email                         // 邮箱
	userInfo.Ip = ip                                            // 封装获取ip的方法

	// 用户扩展信息
	userExtendInfo.Uid = uid
	userExtendInfo.Birthday = registerInfo.Birthday
	userExtendInfo.Sex = registerInfo.Sex

	// 数据事物写入用户表、用户扩展表
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&userInfo).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&userExtendInfo).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

type Person struct {
	AgeGroup string `json:"age_group"`
	Num      int    `json:"num"`
}

// 用户分析接口: 按年龄段统计注册用户数。如20岁以下，20-30，30-40，40-50，50以上。按数量从大到小排序。
func (u *UserApi) Analysis(c *gin.Context) {
	data := analysisService.GetAnalysisUserByAge()

	var persons []Person
	for k, v := range data {
		persons = append(persons, Person{k, utils.StringToInt(v)})
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Num > persons[j].Num
	})

	response.OkWithDetailed(persons, "获取成功", c)
}
