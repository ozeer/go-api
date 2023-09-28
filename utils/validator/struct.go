package validator

import "log"

type User struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=10,lte=13"`
	Phone   string   `v:"required,e164"`
	Email   string   `v:"required,email"`
	FColor  string   `v:"iscolor"`
	FColor1 string   `v:"hexcolor|rgb|rgba|hsl|hsla"`
	Address *Address `v:"required"`
	// 联系人至少填一个
	ContactUser []*ContactUser `v:"required,gte=1,dive"`
	// 兴趣爱好
	Hobby []string `v:"required,gte=2,dive,required,gte=2,alphaunicode"`
}

type ContactUser struct {
	Name string `v:"required,alphaunicode"`
	// 三选一必填
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"`
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"`
	Address *Address `v:"required_without_all=Email Phone"`
}

type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}

func StructValidate() {
	address := Address{
		Province: "江苏省",
		City:     "南京市",
	}

	cu1 := &ContactUser{
		Name:    "张三",
		Phone:   "+8618701347621",
		Email:   "",
		Address: nil,
	}

	cu2 := &ContactUser{
		Name:    "李四",
		Phone:   "",
		Email:   "zs@gmail.com",
		Address: nil,
	}

	u := &User{
		Name:        "nick",
		Age:         12,
		Phone:       "+8618701588471",
		Email:       "nick@gmail.com",
		FColor:      "#ffff",
		FColor1:     "rgb(255,255,255)",
		Address:     &address,
		ContactUser: []*ContactUser{cu1, cu2},
		Hobby:       []string{"旅游", "游泳"},
	}

	v := validate
	err := v.Struct(u)

	if err != nil {
		log.Println("校验结构体发生错误:", err)
	}
}
