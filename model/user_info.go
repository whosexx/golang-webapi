package model

import "time"

type Model struct {
	UUID       string    `json:"uuid" gorm:"primaryKey"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;autoCreateTime:milli"`
	LastTime   time.Time `json:"lastTime" gorm:"column:lastTime;autoUpdateTime:milli"`
}

type UserInfo struct {
	Model
	UserId   string `json:"userId" gorm:"column:userId"`
	UserName string `json:"userName" gorm:"column:userName"`
	Sex      string `json:"sex" gorm:"column:sex"`
	Memo     string `json:"memo" gorm:"column:memo"`
}

func (u *UserInfo) TableName() string {
	return "iwallet_userInfo"
}
