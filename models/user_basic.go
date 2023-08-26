package models

import (
	"fmt"
	"mychat/clients"
	"mychat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string //用户名
	PassWord      string //密码
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"` //手机
	Email         string `valid:"email"`                      //邮箱
	Avatar        string //头像
	Identity      string //身份标识
	ClientIp      string
	ClientPort    string
	Salt          string //MD5加密盐值
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	clients.Mysql.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	clients.Mysql.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	clients.Mysql.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	clients.Mysql.Where("name = ?", name).First(&user)
	return user
}
func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return clients.Mysql.Where("Phone = ?", phone).First(&user)
}
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return clients.Mysql.Where("email = ?", email).First(&user)
}
func CreateUser(user UserBasic) *gorm.DB {
	return clients.Mysql.Create(&user)
}
func DeleteUser(user UserBasic) *gorm.DB {
	return clients.Mysql.Delete(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return clients.Mysql.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email, Avatar: user.Avatar})
}

//查找某个用户
func FindByID(id uint) UserBasic {
	user := UserBasic{}
	clients.Mysql.Where("id = ?", id).First(&user)
	return user
}
