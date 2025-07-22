package initialize

import (
	"github.com/Js41313/Futuer-2/internal/svc"
)

func StartInitSystemConfig(svc *svc.ServiceContext) {
	Migrate(svc)
	Site(svc)
	Node(svc)
	Email(svc)
	Invite(svc)
	Verify(svc)
	Subscribe(svc)
	Register(svc)
	Mobile(svc)
	TrafficDataToRedis(svc)
	if !svc.Config.Debug {
		Telegram(svc)
	}

}
