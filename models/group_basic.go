package models

import (
	"fmt"
	"gorm.io/gorm"
	"mychat/clients"
)

//群信息
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}

//创建群
func CreateGroup(group GroupBasic) (int, string) {
	tx := clients.Mysql.Begin()
	//事务一旦开始，不论什么异常最终都会 Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if len(group.Name) == 0 {
		return -1, "群名称不能为空"
	}
	if group.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := clients.Mysql.Create(&group).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, "建群失败"
	}
	contact := UserContact{}
	contact.OwnerId = group.OwnerId
	contact.TargetId = group.ID
	contact.Type = 2 //群关系
	if err := clients.Mysql.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加群关系失败"
	}

	tx.Commit()
	return 0, "建群成功"

}

//加载群列表
func LoadGroup(ownerId uint) ([]*GroupBasic, string) {
	contacts := make([]UserContact, 0)
	objIds := make([]uint64, 0)
	clients.Mysql.Where("owner_id = ? and type=2", ownerId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}

	data := make([]*GroupBasic, 10)
	clients.Mysql.Where("id in ?", objIds).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	//clients.Mysql.Where()
	return data, "查询成功"
}
