package functions

import "github.com/shima-park/agollo"

type AppInfo struct {
	AppName string
}

const (
	FlagAppName = "app.name"
)

func GetAppInfo() AppInfo {
	appName := agollo.Get(FlagAppName)
	return AppInfo{AppName: appName}
}