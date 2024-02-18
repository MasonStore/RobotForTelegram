package model

import "gorm.io/gorm"

type Account struct {
	*gorm.Model
	Token     string `gorm:"column:token;type:text;not null;" json:"token"`
	FirstName string `gorm:"column:first_name;type:text;not null;default:'';" json:"first_name"`
	LastName  string `gorm:"column:last_name;type:text;not null;default:'';" json:"last_name"`
	UserName  string `gorm:"column:username;type:text;not null;default:'';" json:"username"`
}

func (m *Account) TableName() string {
	return "accounts"
}

func GetAccountById(db *gorm.DB, id uint) (*Account, bool) {
	var account Account
	db.Where("id = ?", id).First(&account)
	if account.ID == 0 {
		return nil, false
	}
	return &account, true
}
