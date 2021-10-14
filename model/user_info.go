package model

import (
	"time"
)

var Tables []interface{} = []interface{}{new(UserInfo)}

type UserInfo struct {
	UUID     string `json:"uuid" gorm:"column:uuid;primaryKey"`
	UserId   string `json:"userId" gorm:"column:userId;uniqueIndex:userId_idx"`
	UserName string `json:"userName" gorm:"column:userName"`
	Sex      string `json:"sex" gorm:"column:sex"`
	Memo     string `json:"memo" gorm:"column:memo"`

	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime"`
	LastTime   time.Time `json:"lastTime" gorm:"column:lastTime;autoUpdateTime"`
}

func (u *UserInfo) TableName() string {
	return "iwallet_userInfo"
}
