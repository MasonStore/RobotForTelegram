package store

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path/filepath"
	goruntime "runtime"
)

func GetDbPath(ctx context.Context) (dbPath string) {
	var goos = goruntime.GOOS
	if goos == "windows" {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		dbPath = filepath.Join(exPath, "store.db")
	} else if goos == "darwin" {
		// 获取HOME目录
		home, err := os.UserHomeDir()
		if err != nil {
			_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
				Type:          runtime.ErrorDialog,
				Title:         "Get User Home Error",
				Message:       err.Error(),
				Buttons:       nil,
				DefaultButton: "",
				CancelButton:  "",
				Icon:          nil,
			})
			runtime.Quit(ctx)
			panic(err)
		}
		dbPath = filepath.Join(home, "Library", "Application Support", "com.wails.RobotForTelegram", "store.db")
		dbDir := filepath.Dir(dbPath)
		if _, err := os.Stat(dbDir); os.IsNotExist(err) {
			err = os.MkdirAll(dbDir, 0755)
			if err != nil {
				_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
					Type:          runtime.ErrorDialog,
					Title:         "Create Directory Error",
					Message:       err.Error(),
					Buttons:       nil,
					DefaultButton: "",
					CancelButton:  "",
					Icon:          nil,
				})
				runtime.Quit(ctx)
				panic(err)
			}
		}
	} else if goos == "linux" {
		// 获取HOME目录
		home, err := os.UserHomeDir()
		if err != nil {
			_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
				Type:          runtime.ErrorDialog,
				Title:         "Get User Home Error",
				Message:       err.Error(),
				Buttons:       nil,
				DefaultButton: "",
				CancelButton:  "",
				Icon:          nil,
			})
			runtime.Quit(ctx)
			panic(err)
		}
		dbPath = filepath.Join(home, ".local", "share", "com.wails.RobotForTelegram", "store.db")
		dbDir := filepath.Dir(dbPath)
		if _, err := os.Stat(dbDir); os.IsNotExist(err) {
			err = os.MkdirAll(dbDir, 0755)
			if err != nil {
				_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
					Type:          runtime.ErrorDialog,
					Title:         "Create Directory Error",
					Message:       err.Error(),
					Buttons:       nil,
					DefaultButton: "",
					CancelButton:  "",
					Icon:          nil,
				})
				runtime.Quit(ctx)
				panic(err)
			}
		}
	} else {
		_, _ = runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Unsupported OS",
			Message:       "Unsupported OS",
			Buttons:       nil,
			DefaultButton: "",
			CancelButton:  "",
			Icon:          nil,
		})
		runtime.Quit(ctx)
		panic("Unsupported OS")
	}
	return
}
