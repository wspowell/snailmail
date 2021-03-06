package mailboxes

import (
	"net/http"

	"github.com/wspowell/context"
	"github.com/wspowell/errors"

	"github.com/wspowell/snailmail/resources/db"
)

type getMailboxResponse struct {
	mailboxModel
}

// FIXME: Must only allow mailbox owner to get mailbox details.
type getMailbox struct {
	MailboxAddress string              `spiderweb:"path=mailbox_address"`
	Datastore      db.Datastore        `spiderweb:"resource=datastore"`
	ResponseBody   *getMailboxResponse `spiderweb:"response,mime=application/json"`
}

func (self *getMailbox) Handle(ctx context.Context) (int, error) {
	foundMailbox, err := self.Datastore.GetMailbox(ctx, self.MailboxAddress)
	if err != nil {
		if errors.Is(err, db.ErrMailboxNotFound) {
			return http.StatusNotFound, errors.Propagate(icGetMailboxNotFound, err)
		} else if errors.Is(err, db.ErrInternalFailure) {
			return http.StatusInternalServerError, errors.Propagate(icGetMailboxDbError, err)
		} else {
			return http.StatusInternalServerError, errors.Convert(icGetMailboxUnknownDbError, err, errUncaughtDbError)
		}
	}

	self.ResponseBody.Address = foundMailbox.Address
	self.ResponseBody.Owner = string(foundMailbox.Owner)
	self.ResponseBody.Capacity = foundMailbox.Capacity
	self.ResponseBody.Location = geoLocation{
		Latitude:  float32(foundMailbox.Location.Lat),
		Longitude: float32(foundMailbox.Location.Lng),
	}

	return http.StatusOK, nil
}
