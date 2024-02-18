package store

import (
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
)

var Instance = NewStore()

type Store struct {
	DB *gorm.DB
}

func NewStore() *Store {
	// 获取可执行文件的路径 在同级目录创建数据库文件
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	dbPath := filepath.Join(exPath, "store.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
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
