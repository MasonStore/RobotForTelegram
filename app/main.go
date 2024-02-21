package main

import (
	"embed"
	"github.com/MasonStore/RobotForTelegram/app/backend/views"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:     "Robot For TG",
		Width:     900,
		Height:    600,
		MinWidth:  900,
		MinHeight: 600,
		//MaxWidth:          1200,
		//MaxHeight:         800,
		DisableResize:     false,
		Debug:             options.Debug{OpenInspectorOnStartup: false},
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 30, G: 30, B: 30, A: 0},
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			app,
			&views.AccountsView{},
			&views.HomeView{},
			&views.AgreementView{},
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 true,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			WebviewGpuIsDisabled:              true,
			Theme:                             windows.Dark,
			CustomTheme: &windows.ThemeSettings{
				// Theme to use when window is active
				DarkModeTitleBar:   windows.RGB(50, 50, 50),
				DarkModeTitleText:  windows.RGB(224, 224, 224),
				DarkModeBorder:     windows.RGB(50, 50, 50),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
				// Theme to use when window is inactive
				DarkModeTitleBarInactive:   windows.RGB(50, 50, 50),
				DarkModeTitleTextInactive:  windows.RGB(224, 224, 224),
				DarkModeBorderInactive:     windows.RGB(50, 50, 50),
				LightModeTitleBarInactive:  windows.RGB(100, 100, 100),
				LightModeTitleTextInactive: windows.RGB(10, 10, 10),
				LightModeBorderInactive:    windows.RGB(100, 100, 100),
			},
		},
		// Mac platform specific options
		// Mac平台特定选项
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            true,
				UseToolbar:                 true,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Robot For TG",
				Message: "A UI management tool for TG bots.",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
