package views

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

type AgreementView struct {
}

func (a *AgreementView) Quit() {
	if ctx != nil {
		runtime.Quit(ctx)
	} else {
		os.Exit(0)
	}
}
