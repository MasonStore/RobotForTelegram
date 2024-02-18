package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

func T2Keys(keys [][]string) string {
	b, _ := json.Marshal(keys)
	return string(b)
}

func (k *KeyboardGroup) GetKeys() [][]string {
	var keys = make([][]string, 0)
	_ = json.Unmarshal([]byte(k.Keys), &keys)
	return keys
}

type KeyboardGroup struct {
	*Model
	RobotId uint   `json:"robotId" gorm:"column:robot_id;type:integer;not null"`
	Name    string `json:"name" gorm:"column:name;type:text;not null"`
	Keys    string `json:"keys" gorm:"column:keys;type:text;not null"`
}

func (k *KeyboardGroup) TableName() string {
	return "keyboard_group"
}

func SaveKeyboardGroups(tx *gorm.DB, robotId uint, keyboardGroups []KeyboardGroup) {
	oldKeyboardGroups := LoadKeyboardGroups(tx, robotId)
	deleteKeyboardGroups := make([]KeyboardGroup, 0)
	//如果oldKeyboardGroups中的命令不在keyboardGroups中，表示需要删除
	for _, oldKeyboardGroup := range oldKeyboardGroups {
		found := false
		for _, keyboardGroup := range keyboardGroups {
			if oldKeyboardGroup.ID == keyboardGroup.ID {
				found = true
				break
			}
		}
		if !found {
			deleteKeyboardGroups = append(deleteKeyboardGroups, oldKeyboardGroup)
		}
	}
	_ = tx.Transaction(func(tx *gorm.DB) error {
		for _, keyboardGroup := range keyboardGroups {
			tx.Save(&keyboardGroup)
		}
		for _, keyboardGroup := range deleteKeyboardGroups {
			tx.Unscoped().Delete(&keyboardGroup)
		}
		return nil
	})
	return
}

func LoadKeyboardGroups(tx *gorm.DB, robotId uint) []KeyboardGroup {
	var keyboardGroups []KeyboardGroup
	tx.Where("robot_id = ?", robotId).Order("id ASC").Find(&keyboardGroups)
	return keyboardGroups
}
