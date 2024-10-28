package apps

import (
	"github.com/ZLinFeng/play-server/apps/schedule"
	"github.com/ZLinFeng/play-server/apps/sys"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System   sys.SystemRouterGroup
	Schedule schedule.RouterGroup
}
