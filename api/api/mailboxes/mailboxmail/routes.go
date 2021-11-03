package mailboxmail

import (
	"github.com/wspowell/spiderweb/endpoint"
	"github.com/wspowell/spiderweb/server/restful"
	"github.com/wspowell/spiderweb/server/route"
)

var (
	RouteExchange = route.Get("/mailboxes/{mailbox_guid}/mail", &exchangeMail{})
)

func Routes(server *restful.Server, config *endpoint.Config) {
	server.Handle(config, RouteExchange)
}
