package main

import (
	"github.com/wspowell/spiderweb/server/lambda"

	"github.com/wspowell/snailmail/api"
	"github.com/wspowell/snailmail/api/mailboxes/mailboxmail"
)

func main() {
	lambda.New(api.Config(), mailboxmail.RouteExchange()).Start()
}
