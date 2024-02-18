cd ..\app

wails build -platform windows/amd64 -clean -ldflags "-s -w" -nsis
