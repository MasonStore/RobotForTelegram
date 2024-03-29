package store

import (
	"context"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Instance *Store

func InitInstance(ctx context.Context) {
	Instance = NewStore(ctx)
}

type Store struct {
	DB *gorm.DB
}

func NewStore(ctx context.Context) *Store {
	// 获取可执行文件的路径 在同级目录创建数据库文件
	dbPath := GetDbPath(ctx)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Database Error",
			Message:       err.Error(),
			Buttons:       nil,
			DefaultButton: "",
			CancelButton:  "",
			Icon:          nil,
		})
		runtime.Quit(ctx)
		panic(err)
	}
	db.AutoMigrate(
		&model.Account{},
		&model.KV{},
		&model.KeyboardGroup{},
		&model.Reply{},
		&model.Command{},
	)
	return &Store{DB: db}
}
