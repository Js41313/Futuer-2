package handler

import (
	"github.com/Js41313/Futuer-2/internal/svc"
	countrylogic "github.com/Js41313/Futuer-2/queue/logic/country"
	orderLogic "github.com/Js41313/Futuer-2/queue/logic/order"
	smslogic "github.com/Js41313/Futuer-2/queue/logic/sms"
	"github.com/Js41313/Futuer-2/queue/logic/subscription"
	"github.com/Js41313/Futuer-2/queue/logic/traffic"
	"github.com/Js41313/Futuer-2/queue/types"
	"github.com/hibiken/asynq"

	emailLogic "github.com/Js41313/Futuer-2/queue/logic/email"
)

func RegisterHandlers(mux *asynq.ServeMux, serverCtx *svc.ServiceContext) {
	// get country task
	mux.Handle(types.ForthwithGetCountry, countrylogic.NewGetNodeCountryLogic(serverCtx))
	// Send email task
	mux.Handle(types.ForthwithSendEmail, emailLogic.NewSendEmailLogic(serverCtx))
	// Send sms task
	mux.Handle(types.ForthwithSendSms, smslogic.NewSendSmsLogic(serverCtx))
	// Defer close order task
	mux.Handle(types.DeferCloseOrder, orderLogic.NewDeferCloseOrderLogic(serverCtx))
	// Forthwith activate order task
	mux.Handle(types.ForthwithActivateOrder, orderLogic.NewActivateOrderLogic(serverCtx))

	// Forthwith traffic statistics
	mux.Handle(types.ForthwithTrafficStatistics, traffic.NewTrafficStatisticsLogic(serverCtx))

	// Schedule check subscription
	mux.Handle(types.SchedulerCheckSubscription, subscription.NewCheckSubscriptionLogic(serverCtx))

	// Schedule total server data
	mux.Handle(types.SchedulerTotalServerData, traffic.NewServerDataLogic(serverCtx))

	// Schedule reset traffic
	mux.Handle(types.SchedulerResetTraffic, traffic.NewResetTrafficLogic(serverCtx))
}
