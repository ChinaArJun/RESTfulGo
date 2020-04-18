package model

import (
	"RESTfulGo/pkg/auth"
	validator "gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"username" gorm:"column:password;not null" binding:"required" validate:"min=5,max=32"`
}

// 设置用户表名称
func (u *UserModel) TableName() string {
	return "tb_users"
}

// 新建新用户
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func DeleteUser(userId uint64) error {
	db := DB.Self.Where("id = ?", userId).Update("state", "0")
	return db.Error
}

func (u *UserModel)UpdateUser() error {
	return DB.Self.Where("id = ?", u.Id).Update(u).Error
}

func GetUserList(pagesize int, pagenum int) ([]*UserInfo, int64, error)  {
	var userList = make([]*UserInfo, 0)
	var count int64

	db := DB.Self.Where("*").Limit(pagesize).Offset(pagenum).Model(&userList).Count(&count)
	if db.Error != nil {
		return nil, 0, db.Error
	}
	return userList, count, nil
}

func GetUser(userId string) (*UserModel, error) {
	model := &UserModel{}
	db := DB.Self.Where("id = ?", userId).First(&model)
	return model, db.Error
}

// 加密密码
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return err
}

/*
 效验用户参数
*/
func (u *UserModel) Validate() error {
	validator := validator.New()
	return validator.Struct(u)
}
