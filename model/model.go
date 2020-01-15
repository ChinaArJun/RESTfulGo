package model

import "time"

type BaseModel struct {
	//  gorm模型定义  column:数据库名称  "-" json忽略这个字段
	Id int `gorm:"primary_key;auto_increment;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`

	//CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	//UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	//DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserInfo struct {
	Id   int `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Token struct {
	Token string `json:"token"`
}