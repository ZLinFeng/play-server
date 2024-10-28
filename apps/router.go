package apps

import (
	"github.com/ZLinFeng/play-server/apps/schedule"
	sysRouter "github.com/ZLinFeng/play-server/apps/sys/router"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System   sysRouter.SystemRouterGroup
	Schedule schedule.RouterGroup
}
