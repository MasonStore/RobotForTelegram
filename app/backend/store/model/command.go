package model

import "gorm.io/gorm"

type Command struct {
	*Model
	RobotId         uint   `json:"robotId" gorm:"column:robot_id;type:integer;not null"`
	Command         string `json:"command" gorm:"column:command;type:text;not null"`
	ReplyId         string `json:"replyId" gorm:"column:reply_id;type:integer;not null"`
	ReplyName       string `json:"replyName" gorm:"column:reply_name;type:text;not null"`
	ReplyIsTemplate string `json:"replyIsTemplate" gorm:"column:reply_is_template;type:text;not null"`
	HookUrl         string `json:"hookUrl" gorm:"column:hook_url;type:text;not null"`
}

func (c *Command) TableName() string {
	return "command"
}

func SaveCommands(tx *gorm.DB, robotId uint, commands []Command) {
	oldCommands := LoadCommands(tx, robotId)
	deleteCommands := make([]Command, 0)
	//如果oldCommands中的命令不在commands中，表示需要删除
	for _, oldCommand := range oldCommands {
		found := false
		for _, command := range commands {
			if oldCommand.ID == command.ID {
				found = true
				break
			}
		}
		if !found {
			deleteCommands = append(deleteCommands, oldCommand)
		}
	}
	_ = tx.Transaction(func(tx *gorm.DB) error {
		for _, command := range commands {
			tx.Save(&command)
		}
		for _, command := range deleteCommands {
			tx.Unscoped().Delete(&command)
		}
		return nil
	})
	return
}

func LoadCommands(tx *gorm.DB, robotId uint) []Command {
	var commands []Command
	tx.Where("robot_id = ?", robotId).Order("id ASC").Find(&commands)
	return commands
}
