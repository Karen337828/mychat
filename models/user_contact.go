package models

import (
	"gorm.io/gorm"
	"mychat/clients"
)

//人员关系
type UserContact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁 /群 ID
	Type     int  //对应的类型  1好友  2群  3xx
	Desc     string
}

func (table *UserContact) TableName() string {
	return "user_contact"
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]UserContact, 0)
	objIds := make([]uint64, 0)
	clients.Mysql.Where("owner_id = ? and type=1", userId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	clients.Mysql.Where("id in ?", objIds).Find(&users)
	return users
}

//添加好友   自己的ID  ， 好友的ID
func AddFriend(userId uint, targetName string) (int, string) {

	if targetName != "" {
		targetUser := FindUserByName(targetName)
		//fmt.Println(targetUser, " userId        ", )
		if targetUser.Salt != "" {
			if targetUser.ID == userId {
				return -1, "不能添加自己为好友"
			}
			contact0 := UserContact{}
			clients.Mysql.Where("owner_id =?  and target_id =? and type=1", userId, targetUser.ID).Find(&contact0)
			if contact0.ID != 0 {
				return -1, "不能重复添加"
			}
			tx := clients.Mysql.Begin()
			//事务一旦开始，不论什么异常最终都会 Rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()
			contact := UserContact{}
			contact.OwnerId = userId
			contact.TargetId = targetUser.ID
			contact.Type = 1
			if err := clients.Mysql.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			contact1 := UserContact{}
			contact1.OwnerId = targetUser.ID
			contact1.TargetId = userId
			contact1.Type = 1
			if err := clients.Mysql.Create(&contact1).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			tx.Commit()
			return 0, "添加好友成功"
		}
		return -1, "没有找到此用户"
	}
	return -1, "好友ID不能为空"
}

func SearchUserByGroupId(communityId uint) []uint {
	contacts := make([]UserContact, 0)
	objIds := make([]uint, 0)
	clients.Mysql.Where("target_id = ? and type=2", communityId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, v.OwnerId)
	}
	return objIds
}
